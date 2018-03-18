// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import (
	"log"
	"os"
	"syscall"
)

func ExampleExec() {
	err := syscall.Exec("/bin/ls", []string{"ls", "-al"}, os.Environ())
	log.Fatal(err)
}
