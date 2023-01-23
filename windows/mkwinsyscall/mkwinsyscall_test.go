// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"testing"
)

func TestDLLFilenameEscaping(t *testing.T) {
	tests := []struct {
		name     string
		filename string
	}{
		{"no escaping necessary", "kernel32"},
		{"escape period", "windows.networking"},
		{"escape dash", "api-ms-win-wsl-api-l1-1-0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Write a made-up syscall into a temp file for testing.
			const prefix = "package windows\n//sys Example() = "
			const suffix = ".Example"
			name := filepath.Join(t.TempDir(), "syscall.go")
			if err := os.WriteFile(name, []byte(prefix+tt.filename+suffix), 0666); err != nil {
				t.Fatal(err)
			}

			// Ensure parsing, generating, and formatting run without errors.
			// This is good enough to show that escaping is working.
			src, err := ParseFiles([]string{name})
			if err != nil {
				t.Fatal(err)
			}
			var buf bytes.Buffer
			if err := src.Generate(&buf); err != nil {
				t.Fatal(err)
			}
			if _, err := format.Source(buf.Bytes()); err != nil {
				t.Log(buf.String())
				t.Fatal(err)
			}
		})
	}
}
