// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

func doinit() {
	c := readCaps()
	ARM64.HasFP = c.has(commpageHasVFP)
	ARM64.HasASIMD = c.has(commpageHasNeon)
	ARM64.HasCRC32 = c.has(commpageHasARMv8CRC32)
	ARM64.HasATOMICS = c.has(commpageHasARMv81Atomics)
	ARM64.HasFPHP = c.has(commpageHasNeonFP16)
	ARM64.HasASIMDHP = c.has(commpageHasNeonHPFP)
	ARM64.HasSHA3 = c.has(commpageHasARMv82SHA3)
	ARM64.HasSHA512 = c.has(commpageHasARMv82SHA512)
	ARM64.HasASIMDFHM = c.has(commpageHasARMv82FHM)
	ARM64.HasSVE = c.has(commpageHasEvent)

	// As of xnu-7195.101.1, there aren't any commpage values for
	// the following features.
	//
	// Assume the following features are available on Apple M1.
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

// Constants taken from
// https://github.com/apple-oss-distributions/xnu/blob/1031c584a5e37aff177559b9f69dbd3c8c3fd30a/osfmk/arm/cpu_capabilities.h
const (
	commpageHasNeonFP16      caps = 0x00000008         // ARM v8.2 NEON FP16
	commpageHasNeon          caps = 0x00000100         // Advanced SIMD
	commpageHasNeonHPFP      caps = 0x00000200         // Advanced SIMD half-precision
	commpageHasVFP           caps = 0x00000400         // VFP
	commpageHasEvent         caps = 0x00001000         // WFE/SVE and period event wakeup
	commpageHasFMA           caps = 0x00002000         // Fused multiply add
	commpageHasARMv82FHM     caps = 0x00004000         // Optional ARMv8.2 FMLAL/FMLSL instructions (required in ARMv8.4)
	commpageHasARMv8Crypto   caps = 0x01000000         // Optional ARMv8 Crypto extensions
	commpageHasARMv81Atomics caps = 0x02000000         // ARMv8.1 Atomic instructions
	commpageHasARMv8CRC32    caps = 0x04000000         // Optional ARMv8 crc32 instructions (required in ARMv8.1)
	commpageHasARMv82SHA512  caps = 0x80000000         // Optional ARMv8.2 SHA512 instructions
	commpageHasARMv82SHA3    caps = 0x0000000100000000 // Optional ARMv8.2 SHA3 instructions
)

// caps is the set of commpage capabilities.
type caps uint64

// has reports whether the capability is enabled.
func (c caps) has(x caps) bool {
	return c&x == x
}

// readCaps loads the current capabilities from commmpage.
func readCaps() (res caps)
