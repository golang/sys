// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (darwin && amd64) || linux
// +build darwin,amd64 linux

package unix

import (
	"unsafe"

	"golang.org/x/sys/internal/unsafeheader"
)

// SysvShmAttach attaches the Sysv shared memory segment associated with the
// shared memory identifier id.
func SysvShmAttach(id int, addr uintptr, flag int) ([]byte, error) {
	addr, errno := shmat(id, addr, flag)
	if errno != nil {
		return nil, errno
	}

	// Retrieve the size of the shared memory to enable slice creation
	var info SysvShmDesc

	_, err := SysvShmCtl(id, IPC_STAT, &info)
	if err != nil {
		// release the shared memory if we can't find the size

		// ignoring error from shmdt as there's nothing sensible to return here
		shmdt(addr)
		return nil, err
	}

	// Use unsafe to convert addr into a []byte.
	var b []byte
	hdr := (*unsafeheader.Slice)(unsafe.Pointer(&b))
	hdr.Data = unsafe.Pointer(addr)
	hdr.Cap = int(info.Segsz)
	hdr.Len = int(info.Segsz)
	return b, nil
}

// SysvShmCtl performs control operations on the shared memory segment
// specified by id.
func SysvShmCtl(id, cmd int, desc *SysvShmDesc) (result int, err error) {
	return shmctl(id, cmd, desc)
}

// SysvShmDetach unmaps the shared memory slice returned from SysvShmAttach.
//
// It is not safe to use the slice after calling this function.
func SysvShmDetach(data []byte) error {
	if data == nil {
		return EINVAL
	}

	return shmdt(uintptr(unsafe.Pointer(&data[0])))
}

// SysvShmGet returns the Sysv shared memory identifier associated with key.
func SysvShmGet(key, size, flag int) (id int, err error) {
	return shmget(key, size, flag)
}
