// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux
// +build linux

package unix

import (
	"testing"
	"unsafe"
)

// An ifreqUnion is shorthand for a byte array matching the
// architecture-dependent size of an ifreq's union field.
type ifreqUnion = [len(ifreq{}.Ifru)]byte

func TestNewIfreq(t *testing.T) {
	// Interface name too long.
	if _, err := NewIfreq("abcdefghijklmnop"); err != EINVAL {
		t.Fatalf("expected error EINVAL, but got: %v", err)
	}
}

func TestIfreqSize(t *testing.T) {
	// Ensure ifreq (generated) and Ifreq/ifreqData (hand-written to create a
	// safe wrapper and store a pointer field) are identical in size.
	want := unsafe.Sizeof(ifreq{})
	if got := unsafe.Sizeof(Ifreq{}); want != got {
		t.Fatalf("unexpected Ifreq size: got: %d, want: %d", got, want)
	}

	if got := unsafe.Sizeof(ifreqData{}); want != got {
		t.Fatalf("unexpected IfreqData size: got: %d, want: %d", got, want)
	}
}

func TestIfreqName(t *testing.T) {
	// Invalid ifreq (no NULL terminator), so expect empty string.
	var name [IFNAMSIZ]byte
	for i := range name {
		name[i] = 0xff
	}

	bad := &Ifreq{raw: ifreq{Ifrn: name}}
	if got := bad.Name(); got != "" {
		t.Fatalf("expected empty ifreq name, but got: %q", got)
	}

	// Valid ifreq, expect the hard-coded testIfreq name.
	ifr := testIfreq(t)
	if want, got := ifreqName, ifr.Name(); want != got {
		t.Fatalf("unexpected ifreq name: got: %q, want: %q", got, want)
	}
}

func TestIfreqWithData(t *testing.T) {
	ifr := testIfreq(t)

	// Store pointer data in the ifreq so we can retrieve it and cast back later
	// for comparison.
	want := [5]byte{'h', 'e', 'l', 'l', 'o'}
	ifrd := ifr.withData(unsafe.Pointer(&want[0]))

	// Ensure the memory of the original Ifreq was not modified by SetData.
	if ifr.raw.Ifru != (ifreqUnion{}) {
		t.Fatalf("ifreq was unexpectedly modified: % #x", ifr.raw.Ifru)
	}

	got := *(*[5]byte)(ifrd.data)
	if want != got {
		t.Fatalf("unexpected ifreq data bytes:\n got: % #x\nwant: % #x", got, want)
	}
}

func TestIfreqUint16(t *testing.T) {
	ifr := testIfreq(t)
	const in = 0x0102
	ifr.SetUint16(in)

	// The layout of the bytes depends on the machine's endianness.
	var want ifreqUnion
	if isBigEndian {
		want[0] = 0x01
		want[1] = 0x02
	} else {
		want[0] = 0x02
		want[1] = 0x01
	}

	if got := ifr.raw.Ifru; want != got {
		t.Fatalf("unexpected ifreq uint16 bytes:\n got: % #x\nwant: % #x", got, want)
	}

	if got := ifr.Uint16(); in != got {
		t.Fatalf("unexpected ifreq uint16: got: %d, want: %d", got, in)
	}
}

func TestIfreqUint32(t *testing.T) {
	ifr := testIfreq(t)
	const in = 0x01020304
	ifr.SetUint32(in)

	// The layout of the bytes depends on the machine's endianness.
	var want ifreqUnion
	if isBigEndian {
		want[0] = 0x01
		want[1] = 0x02
		want[2] = 0x03
		want[3] = 0x04
	} else {
		want[0] = 0x04
		want[1] = 0x03
		want[2] = 0x02
		want[3] = 0x01
	}

	if got := ifr.raw.Ifru; want != got {
		t.Fatalf("unexpected ifreq uint32 bytes:\n got: % #x\nwant: % #x", got, want)
	}

	if got := ifr.Uint32(); in != got {
		t.Fatalf("unexpected ifreq uint32: got: %d, want: %d", got, in)
	}
}

// ifreqName is a hard-coded name for testIfreq.
const ifreqName = "eth0"

// testIfreq returns an Ifreq with a populated interface name.
func testIfreq(t *testing.T) *Ifreq {
	t.Helper()

	ifr, err := NewIfreq(ifreqName)
	if err != nil {
		t.Fatalf("failed to create ifreq: %v", err)
	}

	return ifr
}
