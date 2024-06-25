// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux || netbsd

package unix_test

import (
	"testing"
	"unsafe"

	"golang.org/x/sys/unix"
)

func TestMremap(t *testing.T) {
	b, err := unix.Mmap(-1, 0, unix.Getpagesize(), unix.PROT_NONE, unix.MAP_ANON|unix.MAP_PRIVATE)
	if err != nil {
		t.Fatalf("Mmap: %v", err)
	}
	if err := unix.Mprotect(b, unix.PROT_READ|unix.PROT_WRITE); err != nil {
		t.Fatalf("Mprotect: %v", err)
	}

	b[0] = 42

	bNew, err := unix.Mremap(b, unix.Getpagesize()*2, unix.MremapMaymove)
	if err != nil {
		t.Fatalf("Mremap2: %v", err)
	}
	bNew[unix.Getpagesize()+1] = 84 // checks

	if bNew[0] != 42 {
		t.Fatal("first element value was changed")
	}
	if len(bNew) != unix.Getpagesize()*2 {
		t.Fatal("new memory len not equal to specified len")
	}
	if cap(bNew) != unix.Getpagesize()*2 {
		t.Fatal("new memory cap not equal to specified len")
	}

	_, err = unix.Mremap(b, unix.Getpagesize(), unix.MremapFixed)
	if err != unix.EINVAL {
		t.Fatalf("remapping to a fixed address; got %v, want %v", err, unix.EINVAL)
	}
}

func TestMremapPtr(t *testing.T) {
	mmapProt := unix.PROT_NONE
	mmapPtrProt := unix.PROT_READ | unix.PROT_WRITE
	b, err := unix.Mmap(-1, 0, 2*unix.Getpagesize(), mmapProt, unix.MAP_ANON|unix.MAP_PRIVATE)
	if err != nil {
		t.Fatalf("Mmap: %v", err)
	}
	if _, err := unix.MmapPtr(-1, 0, unsafe.Pointer(&b[0]), uintptr(unix.Getpagesize()),
		mmapPtrProt, unix.MAP_ANON|unix.MAP_PRIVATE|unix.MAP_FIXED); err != nil {
		t.Fatalf("MmapPtr: %v", err)
	}

	b[0] = 42

	if _, err := unix.MremapPtr(
		unsafe.Pointer(&b[0]), uintptr(unix.Getpagesize()),
		unsafe.Pointer(&b[unix.Getpagesize()]), uintptr(unix.Getpagesize()),
		unix.MremapFixed|unix.MremapMaymove); err != nil {
		t.Fatalf("MremapPtr: %v", err)
	}
	if got := b[unix.Getpagesize()]; got != 42 {
		t.Errorf("got %d, want 42", got)
	}

	if err := unix.Munmap(b); err != nil {
		t.Fatalf("Munmap: %v", err)
	}
}
