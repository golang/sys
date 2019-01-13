// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package windows

import "testing"

func TestMacros(t *testing.T) {
	t.Run("HIBYTE", func(t *testing.T) {
		value, want := uint16(0x1234), byte(0x12)
		if got := HIBYTE(value); got != want {
			t.Errorf("got 0x%X, want 0x%X", got, want)
		}
	})
	t.Run("HIWORD", func(t *testing.T) {
		value, want := uint32(0x12345678), uint16(0x1234)
		if got := HIWORD(value); got != want {
			t.Errorf("got 0x%X, want 0x%X", got, want)
		}
	})
	t.Run("LOBYTE", func(t *testing.T) {
		value, want := uint16(0x1234), byte(0x34)
		if got := LOBYTE(value); got != want {
			t.Errorf("got 0x%X, want 0x%X", got, want)
		}
	})
	t.Run("LOWORD", func(t *testing.T) {
		value, want := uint32(0x12345678), uint16(0x5678)
		if got := LOWORD(value); got != want {
			t.Errorf("got 0x%X, want 0x%X", got, want)
		}
	})
	t.Run("MAKELONG", func(t *testing.T) {
		low, high, want := uint16(0x5678), uint16(0x1234), uint32(0x12345678)
		if got := MAKELONG(low, high); got != want {
			t.Errorf("got 0x%X, want 0x%X", got, want)
		}
	})
	t.Run("MAKEWORD", func(t *testing.T) {
		low, high, want := byte(0x34), byte(0x12), uint16(0x1234)
		if got := MAKEWORD(low, high); got != want {
			t.Errorf("got 0x%X, want 0x%X", got, want)
		}
	})
}
