// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// Constants taken from
// https://github.com/apple-oss-distributions/xnu/blob/1031c584a5e37aff177559b9f69dbd3c8c3fd30a/osfmk/arm/cpu_capabilities.h

#define _COMM_PAGE64_BASE_ADDRESS     0x0000000FFFFFC000
#define _COMM_PAGE_START_ADDRESS      _COMM_PAGE64_BASE_ADDRESS
#define _COMM_PAGE_CPU_CAPABILITIES64 (_COMM_PAGE_START_ADDRESS+0x010)
#define _COMM_PAGE_CPU_CAPABILITIES32 (_COMM_PAGE_START_ADDRESS+0x020)
#define _COMM_PAGE_VERSION            (_COMM_PAGE_START_ADDRESS+0x01E)

// func readCaps() (res caps)
TEXT ·readCaps(SB), NOSPLIT, $0-8
#define ptr R0
#define caps R1

	MOVD ZR, ptr
	MOVD ZR, caps

	// We can't check the 64-bit capabilities on iOS because they
	// might not exist. They were added in xnu-7195.50.7.100.1
	// (iOS 14 and macOS 11), and Go supports older iOS versions.
	// _COMM_PAGE_VERSION has stayed the same (3) across kernel
	// versions, so there doesn't appear to be a way to determine
	// whether the 64-bit capabilities exist.
	//
	// The story is different for macOS because Apple's M1 CPUs
	// only support Big Sur and newer.
#ifdef GOOS_ios
	MOVWU $_COMM_PAGE_CPU_CAPABILITIES32, ptr
	MOVWU (ptr), caps

#else
	MOVD $_COMM_PAGE_CPU_CAPABILITIES64, ptr
	MOVD (ptr), caps

#endif

done:
	MOVD caps, res+0(FP)
	RET

#undef ptr
#undef caps
