// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

import "runtime"

func doinit() {
	ARM64.HasASIMD = true
	ARM64.HasFP = true

	if runtime.GOOS != "ios" {
		// M-series SoCs are at least armv8.4-a
		ARM64.HasCRC32 = true   // armv8.1
		ARM64.HasATOMICS = true // armv8.2
		ARM64.HasJSCVT = true   // armv8.3, if HasFP

		// Go already assumes these to be available
		// because they were on the M1.
		ARM64.HasAES = true
		ARM64.HasPMULL = true
		ARM64.HasSHA1 = true
		ARM64.HasSHA2 = true
	}
}
