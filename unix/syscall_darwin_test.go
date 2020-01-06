// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix_test

import (
	"testing"

	"golang.org/x/sys/unix"
)

// stringsFromByteSlice converts a sequence of attributes to a []string.
// On Darwin, each entry is a NULL-terminated string.
func stringsFromByteSlice(buf []byte) []string {
	var result []string
	off := 0
	for i, b := range buf {
		if b == 0 {
			result = append(result, string(buf[off:i]))
			off = i + 1
		}
	}
	return result
}

func TestSysctlClockinfo(t *testing.T) {
	ci, err := unix.SysctlClockinfo("kern.clockrate")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("tick = %v, tickadj = %v, hz = %v, profhz = %v, stathz = %v",
		ci.Tick, ci.Tickadj, ci.Hz, ci.Profhz, ci.Stathz)
}
