// Copyright 2024 The Go Authors. All right reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package unix_test

import (
	"testing"
	"time"

	"golang.org/x/sys/unix"
)

func TestTimeToPtpClockTime(t *testing.T) {
	testcases := []struct {
		time  time.Time
	}{
		{time.Unix(0, 0)},
		{time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)},
		{time.Date(2262, time.December, 31, 23, 0, 0, 0, time.UTC)},
		{time.Unix(0x7FFFFFFF, 0)},
		{time.Unix(0x80000000, 0)},
		{time.Unix(0x7FFFFFFF, 1000000000)},
		{time.Unix(0x7FFFFFFF, 999999999)},
		{time.Unix(-0x80000000, 0)},
		{time.Unix(-0x80000001, 0)},
		{time.Date(2038, time.January, 19, 3, 14, 7, 0, time.UTC)},
		{time.Date(2038, time.January, 19, 3, 14, 8, 0, time.UTC)},
		{time.Date(1901, time.December, 13, 20, 45, 52, 0, time.UTC)},
		{time.Date(1901, time.December, 13, 20, 45, 51, 0, time.UTC)},
	}

	for _, tc := range testcases {
		ts := unix.TimeToPtpClockTime(tc.time)
		tstime := time.Unix(int64(ts.Sec), int64(ts.Nsec))
		if !tstime.Equal(tc.time) {
			t.Errorf("TimeToPtpClockTime(%v) is the time %v", tc.time, tstime)
		}
	}
}
