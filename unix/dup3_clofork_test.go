// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build dragonfly || freebsd

package unix_test

import (
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/sys/unix"
)

func TestDup3Flags(t *testing.T) {
	f, err := os.Create(filepath.Join(t.TempDir(), "TestDup3Flags"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	oldFd := int(f.Fd())

	tests := []struct {
		name   string
		flags  int
		fdflag int
	}{
		{"none", 0, 0},
		{"O_CLOEXEC", unix.O_CLOEXEC, unix.FD_CLOEXEC},
		{"O_CLOFORK", unix.O_CLOFORK, unix.FD_CLOFORK},
		{"O_CLOEXEC|O_CLOFORK", unix.O_CLOEXEC | unix.O_CLOFORK, unix.FD_CLOEXEC | unix.FD_CLOFORK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newFd, err := unix.Open(os.DevNull, unix.O_RDONLY, 0)
			if err != nil {
				t.Fatal(err)
			}
			defer unix.Close(newFd)

			err = unix.Dup3(oldFd, newFd, tt.flags)
			if err == unix.EINVAL && tt.flags&unix.O_CLOFORK != 0 {
				t.Skip("kernel does not support F_DUP3FD")
			}
			if err != nil {
				t.Fatalf("Dup3(%d, %d, %#x): %v", oldFd, newFd, tt.flags, err)
			}
			got, err := unix.FcntlInt(uintptr(newFd), unix.F_GETFD, 0)
			if err != nil {
				t.Fatal(err)
			}
			if want := tt.fdflag; got != want {
				t.Errorf("F_GETFD = %#x, want %#x", got, want)
			}
		})
	}

	if got, want := unix.Dup3(oldFd, oldFd, 0), unix.EINVAL; got != want {
		t.Errorf("Dup3 with equal fds = %v, want %v", got, want)
	}
	if got, want := unix.Dup3(oldFd, 1000, unix.O_APPEND), unix.EINVAL; got != want {
		t.Errorf("Dup3 with invalid flags = %v, want %v", got, want)
	}
}
