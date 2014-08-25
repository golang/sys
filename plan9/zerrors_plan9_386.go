// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plan9

import "syscall"

// Constants
const (
	// Invented values to support what package os expects.
	O_CREAT    = 0x02000
	O_APPEND   = 0x00400
	O_NOCTTY   = 0x00000
	O_NONBLOCK = 0x00000
	O_SYNC     = 0x00000
	O_ASYNC    = 0x00000

	S_IFMT   = 0x1f000
	S_IFIFO  = 0x1000
	S_IFCHR  = 0x2000
	S_IFDIR  = 0x4000
	S_IFBLK  = 0x6000
	S_IFREG  = 0x8000
	S_IFLNK  = 0xa000
	S_IFSOCK = 0xc000
)

// Errors
var (
	EINVAL       = syscall.ErrorString("bad arg in system call")
	ENOTDIR      = syscall.ErrorString("not a directory")
	EISDIR       = syscall.ErrorString("file is a directory")
	ENOENT       = syscall.ErrorString("file does not exist")
	EEXIST       = syscall.ErrorString("file already exists")
	EMFILE       = syscall.ErrorString("no free file descriptors")
	EIO          = syscall.ErrorString("i/o error")
	ENAMETOOLONG = syscall.ErrorString("file name too long")
	EINTR        = syscall.ErrorString("interrupted")
	EPERM        = syscall.ErrorString("permission denied")
	EBUSY        = syscall.ErrorString("no free devices")
	ETIMEDOUT    = syscall.ErrorString("connection timed out")
	EPLAN9       = syscall.ErrorString("not supported by plan 9")

	// The following errors do not correspond to any
	// Plan 9 system messages. Invented to support
	// what package os and others expect.
	EACCES       = syscall.ErrorString("access permission denied")
	EAFNOSUPPORT = syscall.ErrorString("address family not supported by protocol")
)
