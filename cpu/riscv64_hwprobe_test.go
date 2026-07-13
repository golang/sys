// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && riscv64

package cpu

import (
	"testing"
	"unsafe"

	"golang.org/x/sys/unix"
)

var (
	_ int64  = riscvHWProbePairs{}.key
	_ int64  = unix.RISCVHWProbePairs{}.Key
	_ uint64 = riscvHWProbePairs{}.value
	_ uint64 = unix.RISCVHWProbePairs{}.Value
)

func TestRISCV64HWProbeCopiedConstants(t *testing.T) {
	copiedConsts := []struct {
		name string
		got  uint
		want uint
	}{
		{"riscv_HWPROBE_KEY_IMA_EXT_0", riscv_HWPROBE_KEY_IMA_EXT_0, unix.RISCV_HWPROBE_KEY_IMA_EXT_0},
		{"riscv_HWPROBE_IMA_C", riscv_HWPROBE_IMA_C, unix.RISCV_HWPROBE_IMA_C},
		{"riscv_HWPROBE_IMA_V", riscv_HWPROBE_IMA_V, unix.RISCV_HWPROBE_IMA_V},
		{"riscv_HWPROBE_EXT_ZBA", riscv_HWPROBE_EXT_ZBA, unix.RISCV_HWPROBE_EXT_ZBA},
		{"riscv_HWPROBE_EXT_ZBB", riscv_HWPROBE_EXT_ZBB, unix.RISCV_HWPROBE_EXT_ZBB},
		{"riscv_HWPROBE_EXT_ZBS", riscv_HWPROBE_EXT_ZBS, unix.RISCV_HWPROBE_EXT_ZBS},
		{"riscv_HWPROBE_EXT_ZBC", riscv_HWPROBE_EXT_ZBC, unix.RISCV_HWPROBE_EXT_ZBC},
		{"riscv_HWPROBE_EXT_ZVBB", riscv_HWPROBE_EXT_ZVBB, unix.RISCV_HWPROBE_EXT_ZVBB},
		{"riscv_HWPROBE_EXT_ZVBC", riscv_HWPROBE_EXT_ZVBC, unix.RISCV_HWPROBE_EXT_ZVBC},
		{"riscv_HWPROBE_EXT_ZVKB", riscv_HWPROBE_EXT_ZVKB, unix.RISCV_HWPROBE_EXT_ZVKB},
		{"riscv_HWPROBE_EXT_ZVKG", riscv_HWPROBE_EXT_ZVKG, unix.RISCV_HWPROBE_EXT_ZVKG},
		{"riscv_HWPROBE_EXT_ZVKNED", riscv_HWPROBE_EXT_ZVKNED, unix.RISCV_HWPROBE_EXT_ZVKNED},
		{"riscv_HWPROBE_EXT_ZVKNHB", riscv_HWPROBE_EXT_ZVKNHB, unix.RISCV_HWPROBE_EXT_ZVKNHB},
		{"riscv_HWPROBE_EXT_ZVKSED", riscv_HWPROBE_EXT_ZVKSED, unix.RISCV_HWPROBE_EXT_ZVKSED},
		{"riscv_HWPROBE_EXT_ZVKSH", riscv_HWPROBE_EXT_ZVKSH, unix.RISCV_HWPROBE_EXT_ZVKSH},
		{"riscv_HWPROBE_EXT_ZVKT", riscv_HWPROBE_EXT_ZVKT, unix.RISCV_HWPROBE_EXT_ZVKT},
		{"riscv_HWPROBE_KEY_CPUPERF_0", riscv_HWPROBE_KEY_CPUPERF_0, unix.RISCV_HWPROBE_KEY_CPUPERF_0},
		{"riscv_HWPROBE_MISALIGNED_FAST", riscv_HWPROBE_MISALIGNED_FAST, unix.RISCV_HWPROBE_MISALIGNED_FAST},
		{"riscv_HWPROBE_MISALIGNED_MASK", riscv_HWPROBE_MISALIGNED_MASK, unix.RISCV_HWPROBE_MISALIGNED_MASK},
	}

	for _, c := range copiedConsts {
		if c.got != c.want {
			t.Errorf("%s = %#x, want %#x", c.name, c.got, c.want)
		}
	}
}

func TestRISCV64HWProbeCopiedSyscallNumber(t *testing.T) {
	if sys_RISCV_HWPROBE != unix.SYS_RISCV_HWPROBE {
		t.Errorf("sys_RISCV_HWPROBE = %#x, want %#x", sys_RISCV_HWPROBE, unix.SYS_RISCV_HWPROBE)
	}
}

func TestRISCV64HWProbePairsLayout(t *testing.T) {
	if got, want := unsafe.Sizeof(riscvHWProbePairs{}), unsafe.Sizeof(unix.RISCVHWProbePairs{}); got != want {
		t.Errorf("riscvHWProbePairs size = %d, want RISCVHWProbePairs size = %d", got, want)
	}
	if got, want := unsafe.Offsetof(riscvHWProbePairs{}.key), unsafe.Offsetof(unix.RISCVHWProbePairs{}.Key); got != want {
		t.Errorf("riscvHWProbePairs.key offset = %d, want RISCVHWProbePairs.Key offset = %d", got, want)
	}
	if got, want := unsafe.Offsetof(riscvHWProbePairs{}.value), unsafe.Offsetof(unix.RISCVHWProbePairs{}.Value); got != want {
		t.Errorf("riscvHWProbePairs.value offset = %d, want RISCVHWProbePairs.Value offset = %d", got, want)
	}
}
