// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build 386 || amd64 || amd64p32

package cpu

import "runtime"

const cacheLineSize = 64

func initOptions() {
	options = []option{
		{Name: "adx", Feature: &X86.HasADX},
		{Name: "aes", Feature: &X86.HasAES},
		{Name: "avx", Feature: &X86.HasAVX},
		{Name: "avx2", Feature: &X86.HasAVX2},
		{Name: "avx512", Feature: &X86.HasAVX512},
		{Name: "avx512f", Feature: &X86.HasAVX512F},
		{Name: "avx512cd", Feature: &X86.HasAVX512CD},
		{Name: "avx512er", Feature: &X86.HasAVX512ER},
		{Name: "avx512pf", Feature: &X86.HasAVX512PF},
		{Name: "avx512vl", Feature: &X86.HasAVX512VL},
		{Name: "avx512bw", Feature: &X86.HasAVX512BW},
		{Name: "avx512dq", Feature: &X86.HasAVX512DQ},
		{Name: "avx512ifma", Feature: &X86.HasAVX512IFMA},
		{Name: "avx512vbmi", Feature: &X86.HasAVX512VBMI},
		{Name: "avx512vnniw", Feature: &X86.HasAVX5124VNNIW},
		{Name: "avx5124fmaps", Feature: &X86.HasAVX5124FMAPS},
		{Name: "avx512vpopcntdq", Feature: &X86.HasAVX512VPOPCNTDQ},
		{Name: "avx512vpclmulqdq", Feature: &X86.HasAVX512VPCLMULQDQ},
		{Name: "avx512vnni", Feature: &X86.HasAVX512VNNI},
		{Name: "avx512gfni", Feature: &X86.HasAVX512GFNI},
		{Name: "avx512vaes", Feature: &X86.HasAVX512VAES},
		{Name: "avx512vbmi2", Feature: &X86.HasAVX512VBMI2},
		{Name: "avx512bitalg", Feature: &X86.HasAVX512BITALG},
		{Name: "avx512bf16", Feature: &X86.HasAVX512BF16},
		{Name: "amxtile", Feature: &X86.HasAMXTile},
		{Name: "amxint8", Feature: &X86.HasAMXInt8},
		{Name: "amxbf16", Feature: &X86.HasAMXBF16},
		{Name: "bmi1", Feature: &X86.HasBMI1},
		{Name: "bmi2", Feature: &X86.HasBMI2},
		{Name: "cx16", Feature: &X86.HasCX16},
		{Name: "erms", Feature: &X86.HasERMS},
		{Name: "fma", Feature: &X86.HasFMA},
		{Name: "osxsave", Feature: &X86.HasOSXSAVE},
		{Name: "pclmulqdq", Feature: &X86.HasPCLMULQDQ},
		{Name: "popcnt", Feature: &X86.HasPOPCNT},
		{Name: "rdrand", Feature: &X86.HasRDRAND},
		{Name: "rdseed", Feature: &X86.HasRDSEED},
		{Name: "sse3", Feature: &X86.HasSSE3},
		{Name: "sse41", Feature: &X86.HasSSE41},
		{Name: "sse42", Feature: &X86.HasSSE42},
		{Name: "ssse3", Feature: &X86.HasSSSE3},
		{Name: "avxifma", Feature: &X86.HasAVXIFMA},
		{Name: "avxvnni", Feature: &X86.HasAVXVNNI},
		{Name: "avxvnniint8", Feature: &X86.HasAVXVNNIInt8},

		// These capabilities should always be enabled on amd64:
		{Name: "sse2", Feature: &X86.HasSSE2, Required: runtime.GOARCH == "amd64"},
	}
}

func archInit() {

	Initialized = true

	maxID, _, _, _ := cpuid(0, 0)

	if maxID < 1 {
		return
	}

	_, _, ecx1, edx1 := cpuid(1, 0)
	X86.HasSSE2 = isSet(edx1, 1<<26)

	X86.HasSSE3 = isSet(ecx1, 1<<0)
	X86.HasPCLMULQDQ = isSet(ecx1, 1<<1)
	X86.HasSSSE3 = isSet(ecx1, 1<<9)
	X86.HasFMA = isSet(ecx1, 1<<12)
	X86.HasCX16 = isSet(ecx1, 1<<13)
	X86.HasSSE41 = isSet(ecx1, 1<<19)
	X86.HasSSE42 = isSet(ecx1, 1<<20)
	X86.HasPOPCNT = isSet(ecx1, 1<<23)
	X86.HasAES = isSet(ecx1, 1<<25)
	X86.HasOSXSAVE = isSet(ecx1, 1<<27)
	X86.HasRDRAND = isSet(ecx1, 1<<30)

	var osSupportsAVX, osSupportsAVX512 bool
	// For XGETBV, OSXSAVE bit is required and sufficient.
	if X86.HasOSXSAVE {
		eax, _ := xgetbv()
		// Check if XMM and YMM registers have OS support.
		osSupportsAVX = isSet(eax, 1<<1) && isSet(eax, 1<<2)

		if runtime.GOOS == "darwin" {
			// Darwin requires special AVX512 checks, see cpu_darwin_x86.go
			osSupportsAVX512 = osSupportsAVX && darwinSupportsAVX512()
		} else {
			// Check if OPMASK and ZMM registers have OS support.
			osSupportsAVX512 = osSupportsAVX && isSet(eax, 1<<5) && isSet(eax, 1<<6) && isSet(eax, 1<<7)
		}
	}

	X86.HasAVX = isSet(ecx1, 1<<28) && osSupportsAVX

	if maxID < 7 {
		return
	}

	eax7, ebx7, ecx7, edx7 := cpuid(7, 0)
	X86.HasBMI1 = isSet(ebx7, 1<<3)
	X86.HasAVX2 = isSet(ebx7, 1<<5) && osSupportsAVX
	X86.HasBMI2 = isSet(ebx7, 1<<8)
	X86.HasERMS = isSet(ebx7, 1<<9)
	X86.HasRDSEED = isSet(ebx7, 1<<18)
	X86.HasADX = isSet(ebx7, 1<<19)

	X86.HasAVX512 = isSet(ebx7, 1<<16) && osSupportsAVX512 // Because avx-512 foundation is the core required extension
	if X86.HasAVX512 {
		X86.HasAVX512F = true
		X86.HasAVX512CD = isSet(ebx7, 1<<28)
		X86.HasAVX512ER = isSet(ebx7, 1<<27)
		X86.HasAVX512PF = isSet(ebx7, 1<<26)
		X86.HasAVX512VL = isSet(ebx7, 1<<31)
		X86.HasAVX512BW = isSet(ebx7, 1<<30)
		X86.HasAVX512DQ = isSet(ebx7, 1<<17)
		X86.HasAVX512IFMA = isSet(ebx7, 1<<21)
		X86.HasAVX512VBMI = isSet(ecx7, 1<<1)
		X86.HasAVX5124VNNIW = isSet(edx7, 1<<2)
		X86.HasAVX5124FMAPS = isSet(edx7, 1<<3)
		X86.HasAVX512VPOPCNTDQ = isSet(ecx7, 1<<14)
		X86.HasAVX512VPCLMULQDQ = isSet(ecx7, 1<<10)
		X86.HasAVX512VNNI = isSet(ecx7, 1<<11)
		X86.HasAVX512GFNI = isSet(ecx7, 1<<8)
		X86.HasAVX512VAES = isSet(ecx7, 1<<9)
		X86.HasAVX512VBMI2 = isSet(ecx7, 1<<6)
		X86.HasAVX512BITALG = isSet(ecx7, 1<<12)
	}

	X86.HasAMXTile = isSet(edx7, 1<<24)
	X86.HasAMXInt8 = isSet(edx7, 1<<25)
	X86.HasAMXBF16 = isSet(edx7, 1<<22)

	// These features depend on the second level of extended features.
	if eax7 >= 1 {
		eax71, _, _, edx71 := cpuid(7, 1)
		if X86.HasAVX512 {
			X86.HasAVX512BF16 = isSet(eax71, 1<<5)
		}
		if X86.HasAVX {
			X86.HasAVXIFMA = isSet(eax71, 1<<23)
			X86.HasAVXVNNI = isSet(eax71, 1<<4)
			X86.HasAVXVNNIInt8 = isSet(edx71, 1<<4)
		}
	}
}

func isSet(hwc uint32, value uint32) bool {
	return hwc&value != 0
}
