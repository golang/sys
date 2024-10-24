// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package unix

import "time"

// TimeToPtpClockTime returns t as PtpClockTime
func TimeToPtpClockTime(t time.Time) PtpClockTime {
	sec := t.Unix()
	nsec := uint32(t.Nanosecond())
	return PtpClockTime{Sec: sec, Nsec: nsec}
}

// Time returns PTPClockTime as time.Time
func (t *PtpClockTime) Time() time.Time {
	return time.Unix(t.Sec, int64(t.Nsec))
}

// Unix returns the time stored in t as seconds plus nanoseconds.
func (t *PtpClockTime) Unix() (sec int64, nsec int64) {
	return t.Sec, int64(t.Nsec)
}
