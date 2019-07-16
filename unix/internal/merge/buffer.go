// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package merge

import (
	"bufio"
	"io"
	"os"
)

var _ io.WriteCloser = (*bufferedFile)(nil)

type bufferedFile struct {
	f *os.File
	*bufio.Writer
}

// bufferedFile provides a buffered os.File.
func newbufferedFile(name string) (io.WriteCloser, error) {
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewWriter(f)
	return &bufferedFile{f: f, Writer: buf}, nil
}

func (bf *bufferedFile) Close() error {
	if err := bf.Writer.Flush(); err != nil {
		return err
	}
	return bf.f.Close()
}
