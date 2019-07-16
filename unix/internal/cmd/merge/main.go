// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// merge processes the generated Go files to factorize
// constants, types and functions definitions.
// For all constants, types, and functions that are defined
// precisely identically for each GOARCH, move them into
// a single unified file named after the source file and GOARCH
// (e.g. zerrors_linux.go).
//
//TODO merge is run after ???; see README.md.
package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/sys/unix/internal/merge"
)

func main() {
	var withTiming bool
	flag.BoolVar(&withTiming, "timing", false, "prints out time to execute")
	var withStats bool
	flag.BoolVar(&withStats, "stats", false, "prints out statistics before and after")
	var pkgPath string
	flag.StringVar(&pkgPath, "path", "", "package path")
	flag.Parse()

	if withTiming {
		start := time.Now()
		defer func() {
			fmt.Println(time.Now().Sub(start))
		}()
	}

	// Load the generated source code (file names start with 'z').
	filter := func(fi os.FileInfo) bool {
		name := fi.Name()
		// Skip files not of the form: z<name>_<goos>_<goarch>.go
		return strings.HasPrefix(name, "z") && strings.Count(name, "_") == 2
	}
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, pkgPath, filter, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("invalid number of packages: got %d; want 1", len(pkgs))
	}
	var pkg *ast.Package
	for _, p := range pkgs {
		pkg = p
	}

	// Process the package.
	reg := merge.NewRegistry(pkg)

	if withStats {
		var stats merge.Stats
		reg.ReadStats(&stats)
		fmt.Println("BEFORE")
		fmt.Println(stats)
		defer func() {
			reg.ReadStats(&stats)
			fmt.Println("AFTER")
			fmt.Println(stats)
		}()
	}

	reg.Build()

	// Print out the new files and updated source code.
	if err := reg.Print(pkg, fset); err != nil {
		log.Fatal(err)
	}
}
