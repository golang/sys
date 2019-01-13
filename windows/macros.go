// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package windows

// HIBYTE macro retrieves the high-order byte from the given 16-bit value.
//
// https://msdn.microsoft.com/library/windows/desktop/ms632656
func HIBYTE(wValue uint16) byte {
	return byte(wValue >> 8 & 0xff)
}

// HIWORD macro retrieves the high-order word from the specified 32-bit value.
//
// https://msdn.microsoft.com/library/windows/desktop/ms632657
func HIWORD(dwValue uint32) uint16 {
	return uint16(dwValue >> 16 & 0xffff)
}

// LOBYTE macro retrieves the low-order byte from the specified 16-bit value.
//
// https://msdn.microsoft.com/library/windows/desktop/ms632658
func LOBYTE(wValue uint16) byte {
	return byte(wValue)
}

// LOWORD macro retrieves the low-order word from the specified 32-bit value.
//
// https://msdn.microsoft.com/library/windows/desktop/ms632659
func LOWORD(dwValue uint32) uint16 {
	return uint16(dwValue)
}

// MAKELONG macro creates a LONG value by concatenating the specified values.
//
// https://msdn.microsoft.com/library/windows/desktop/ms632660
func MAKELONG(wLow, wHigh uint16) uint32 {
	return uint32(wLow) | (uint32(wHigh) << 16)
}

// MAKEWORD macro creates a WORD value by concatenating the specified values.
//
// https://msdn.microsoft.com/library/windows/desktop/ms632663
func MAKEWORD(bLow, bHigh byte) uint16 {
	return uint16(bLow) | (uint16(bHigh) << 8)
}
