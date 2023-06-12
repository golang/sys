// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux
// +build linux

package unix_test

import (
	"golang.org/x/sys/unix"
	"testing"
)

func TestMremap(t *testing.T) {
	b, err := unix.Mmap(-1, 0, unix.Getpagesize()*2, unix.PROT_NONE, unix.MAP_ANON|unix.MAP_PRIVATE)
	if err != nil {
		t.Fatalf("Mmap: %v", err)
	}
	if err := unix.Mprotect(b, unix.PROT_READ|unix.PROT_WRITE); err != nil {
		t.Fatalf("Mprotect: %v", err)
	}

	b[0] = 42
	if err := unix.Msync(b, unix.MS_SYNC); err != nil {
		t.Fatalf("Msync: %v", err)
	}

	bNew, err := unix.Mmap(-1, int64(unix.Getpagesize()*4), unix.Getpagesize(), unix.PROT_NONE, unix.MAP_ANON|unix.MAP_PRIVATE)
	if err != nil {
		t.Fatalf("Mmap: %v", err)
	}

	bRemapped, err := unix.Mremap(b, bNew, 0)
	if err != nil {
		t.Fatalf("Mremap: %v", err)
	}
	if &bRemapped[0] != &bNew[0] {
		t.Fatal("bNew and bRemapped start at different pointers")
	}
	if bRemapped[0] != 42 {
		t.Fatal("first element wasn't mapped")
	}
	if len(bNew) != len(bRemapped) {
		t.Fatal("bNew len doesn't equal bRemapped len")
	}
	if cap(bNew) != cap(bRemapped) {
		t.Fatal("bNew cap doesn't equal bRemapped len")
	}
}

func TestMremap2(t *testing.T) {
	b, err := unix.Mmap(-1, 0, unix.Getpagesize()*2, unix.PROT_NONE, unix.MAP_ANON|unix.MAP_PRIVATE)
	if err != nil {
		t.Fatalf("Mmap: %v", err)
	}
	if err := unix.Mprotect(b, unix.PROT_READ|unix.PROT_WRITE); err != nil {
		t.Fatalf("Mprotect: %v", err)
	}

	b[0] = 42
	if err := unix.Msync(b, unix.MS_SYNC); err != nil {
		t.Fatalf("Msync: %v", err)
	}

	bNew, err := unix.Mremap2(b, unix.Getpagesize(), unix.MREMAP_MAYMOVE)
	if err != nil {
		t.Fatalf("Mremap2: %v", err)
	}

	if bNew[0] != 42 {
		t.Fatal("first element value was changed")
	}
	if len(bNew) != unix.Getpagesize() {
		t.Fatal("new memory len not equal to specified len")
	}
	if cap(bNew) != unix.Getpagesize() {
		t.Fatal("new memory cap not equal to specified len")
	}
}
