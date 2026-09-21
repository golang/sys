// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"syscall"
	"testing"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

//go:generate go run . -systemdll=false -output zhook_windows_test.go hook_windows_test.go

//sys	regOpenKeyExWithGarbageUpperBits(key syscall.Handle, subkey *uint16, options uint32, desiredAccess uint32, result *syscall.Handle) (regerrno error) = hook.RegOpenKeyExWWithGarbageUpperBits
//sys	isValidCodePageWithGarbageUpperBits(codePage uint32) (ok bool) = hook.IsValidCodePageWithGarbageUpperBits
//sys	getHandleInformationWithGarbageUpperBits(handle syscall.Handle, flags *uint32) (err error) = hook.GetHandleInformationWithGarbageUpperBits

// hookC is built into hook.dll. Its functions call real 32-bit APIs and, like
// a non-conforming hook, leave garbage in the upper half of the return register.
const hookC = `#include <stdint.h>
#include <windows.h>

#define GARBAGE 0xDEADBEEF00000000ULL

__declspec(dllexport) uintptr_t RegOpenKeyExWWithGarbageUpperBits(HKEY key, LPCWSTR subkey, DWORD options, REGSAM desiredAccess, PHKEY result) {
	return GARBAGE | (uint32_t)RegOpenKeyExW(key, subkey, options, desiredAccess, result);
}

__declspec(dllexport) uintptr_t IsValidCodePageWithGarbageUpperBits(UINT codePage) {
	return GARBAGE | (uint32_t)IsValidCodePage(codePage);
}

__declspec(dllexport) uintptr_t GetHandleInformationWithGarbageUpperBits(HANDLE handle, LPDWORD flags) {
	return GARBAGE | (uint32_t)GetHandleInformation(handle, flags);
}
`

func TestGeneratedCodeIgnoresUpperReturnBits(t *testing.T) {
	if runtime.GOARCH != "amd64" && runtime.GOARCH != "arm64" {
		t.Skip("the return register has no upper 32 bits on this architecture")
	}
	if _, err := exec.LookPath("gcc"); err != nil {
		t.Skip("skipping test: gcc is missing")
	}

	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "hook.c"), []byte(hookC), 0666); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("gcc", "-shared", "-s", "-Werror", "-o", "hook.dll", "hook.c", "-ladvapi32")
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build dll: %v - %v", err, string(out))
	}
	if err := windows.SetDllDirectory(dir); err != nil {
		t.Fatal(err)
	}
	// FreeLibrary below invalidates the cached lazy DLL and procs, so recreate them.
	modhook = syscall.NewLazyDLL("hook.dll")
	procRegOpenKeyExWWithGarbageUpperBits = modhook.NewProc("RegOpenKeyExWWithGarbageUpperBits")
	procIsValidCodePageWithGarbageUpperBits = modhook.NewProc("IsValidCodePageWithGarbageUpperBits")
	procGetHandleInformationWithGarbageUpperBits = modhook.NewProc("GetHandleInformationWithGarbageUpperBits")
	if err := modhook.Load(); err != nil {
		t.Fatal(err)
	}
	// Windows cannot delete a loaded DLL, so free it before t.TempDir's cleanup.
	defer func() {
		if err := syscall.FreeLibrary(syscall.Handle(modhook.Handle())); err != nil {
			t.Error(err)
		}
	}()

	for _, tt := range []struct {
		subkey string
		want   error
	}{
		{`SOFTWARE\Microsoft\Windows NT\CurrentVersion`, nil},
		{`SOFTWARE\golang.org-x-sys-mkwinsyscall-missing`, syscall.ERROR_FILE_NOT_FOUND},
	} {
		subkey, err := windows.UTF16PtrFromString(tt.subkey)
		if err != nil {
			t.Fatal(err)
		}
		var h syscall.Handle
		err = regOpenKeyExWithGarbageUpperBits(syscall.HKEY_LOCAL_MACHINE, subkey, 0, syscall.KEY_QUERY_VALUE, &h)
		if err == nil {
			k := registry.Key(h)
			if _, err := k.Stat(); err != nil {
				t.Error(err)
			}
			if err := k.Close(); err != nil {
				t.Error(err)
			}
		}
		if err != tt.want {
			t.Errorf("regOpenKeyExWithGarbageUpperBits(%q): got %#v, want %#v", tt.subkey, err, tt.want)
		}
	}
	for _, tt := range []struct {
		codePage uint32
		want     bool
	}{
		{65001, true}, // UTF-8
		{12345, false},
	} {
		if got := isValidCodePageWithGarbageUpperBits(tt.codePage); got != tt.want {
			t.Errorf("isValidCodePageWithGarbageUpperBits(%d): got %v, want %v", tt.codePage, got, tt.want)
		}
	}
	f, err := os.Open(filepath.Join(dir, "hook.c"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	for _, tt := range []struct {
		handle syscall.Handle
		want   error
	}{
		{syscall.Handle(f.Fd()), nil},
		{0, windows.ERROR_INVALID_HANDLE},
	} {
		var flags uint32
		if err := getHandleInformationWithGarbageUpperBits(tt.handle, &flags); err != tt.want {
			t.Errorf("getHandleInformationWithGarbageUpperBits(%#x): got %#v, want %#v", tt.handle, err, tt.want)
		}
	}
}
