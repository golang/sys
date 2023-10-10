// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin && arm64 && !ios
// +build darwin,arm64,!ios

package cpu

const (
	commpageHasNeonFP16      = 0x00000008         // ARM v8.2 NEON FP16 supported
	commpageHasNeon          = 0x00000100         // Advanced SIMD is supported
	commpageHasNeonHPFP      = 0x00000200         // Advanced SIMD half-precision
	commpageHasVfp           = 0x00000400         // VFP is supported
	commpageHasEvent         = 0x00001000         // WFE/SVE and period event wakeup
	commpageHasARMv82FHM     = 0x00004000         // Optional ARMv8.2 FMLAL/FMLSL instructions (required in ARMv8.4)
	commpageHasARMv8Crypto   = 0x01000000         // Optional ARMv8 Crypto extensions
	commpageHasARMv81Atomics = 0x02000000         // ARMv8.1 Atomic instructions supported
	commpageHasARMv8Crc32    = 0x04000000         // Optional ARMv8 crc32 instructions (required in ARMv8.1)
	commpageHasARMv82SHA512  = 0x80000000         // Optional ARMv8.2 SHA512 instructions
	commpageHasARMv82SHA3    = 0x0000000100000000 // Optional ARMv8.2 SHA3 instructions
)

func doinit() {
	ARM64.HasFP = darwinCheckFeatureEnabled(commpageHasVfp)
	ARM64.HasASIMD = darwinCheckFeatureEnabled(commpageHasNeon)
	ARM64.HasCRC32 = darwinCheckFeatureEnabled(commpageHasARMv8Crc32)
	ARM64.HasATOMICS = darwinCheckFeatureEnabled(commpageHasARMv81Atomics)
	ARM64.HasFPHP = darwinCheckFeatureEnabled(commpageHasNeonFP16)
	ARM64.HasASIMDHP = darwinCheckFeatureEnabled(commpageHasNeonHPFP)
	ARM64.HasSHA3 = darwinCheckFeatureEnabled(commpageHasARMv82SHA3)
	ARM64.HasSHA512 = darwinCheckFeatureEnabled(commpageHasARMv82SHA512)
	ARM64.HasASIMDFHM = darwinCheckFeatureEnabled(commpageHasARMv82FHM)
	ARM64.HasSVE = darwinCheckFeatureEnabled(commpageHasEvent)

	// There are no hw.optional sysctl values for the below features on Mac OS 11.0
	// to detect their supported state dynamically. Assume the CPU features that
	// Apple Silicon M1 supports to be available as a minimal set of features
	// to all Go programs running on darwin/arm64.
	ARM64.HasEVTSTRM = true
	ARM64.HasAES = true
	ARM64.HasPMULL = true
	ARM64.HasSHA1 = true
	ARM64.HasSHA2 = true
	ARM64.HasCPUID = true
	ARM64.HasASIMDRDM = true
	ARM64.HasJSCVT = true
	ARM64.HasLRCPC = true
	ARM64.HasDCPOP = true
	ARM64.HasSM3 = true
	ARM64.HasSM4 = true
	ARM64.HasASIMDDP = true
}

func darwinCheckFeatureEnabled(feature_vec uint64) bool
