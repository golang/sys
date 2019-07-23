// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// consolidate processes the generated Go files to consolidate
// constants, types and functions definitions.
// For all constants, types, and functions that are defined
// precisely identically for each GOARCH, move them into
// a single unified file named after the source file and GOARCH
// (e.g. zerrors_linux.go).
//
// The z*_goos_goarch.go files must be generated prior to running this program.
//
package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"

	"golang.org/x/sys/unix/internal/consolidate"
)

func main() {
	var cpuProf, memProf string
	flag.StringVar(&cpuProf, "cpuprofile", "", "write cpu profile to file")
	flag.StringVar(&memProf, "memprofile", "", "write mem profile to file")
	var withTiming bool
	flag.BoolVar(&withTiming, "timing", false, "prints out time to execute")
	var withStats bool
	flag.BoolVar(&withStats, "stats", false, "prints out statistics before and after")
	var pkgPath string
	flag.StringVar(&pkgPath, "path", ".", "package path")
	flag.Parse()

	if memProf != "" {
		f, err := os.Create(memProf)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		defer func() {
			runtime.GC()
			if err := pprof.WriteHeapProfile(f); err != nil {
				log.Fatal(err)
			}
		}()
	}
	if cpuProf != "" {
		f, err := os.Create(cpuProf)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

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
	reg := consolidate.NewRegistry(pkg)

	if withStats {
		var stats consolidate.Stats
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

	// Format the new files.
	cmd := exec.Command("go", "fmt")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
