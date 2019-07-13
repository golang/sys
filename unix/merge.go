// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

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
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	// Load the source code excluding test files.
	filter := func(fi os.FileInfo) bool {
		return !strings.HasSuffix(fi.Name(), "_test.go")
	}
	pkgs, err := parser.ParseDir(token.NewFileSet(), ".", filter, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	//// Should only have loaded one package as tests are excluded.
	//delete(pkgs, "main") // Exclude main packages.
	//if len(pkgs) != 1 {
	//	log.Fatalf("unexpected number of packages found: got %d; want 1", len(pkgs))
	//}
	//var reg registry
	//for _, p := range pkgs {
	//	reg = newRegistry(p)
	//	break
	//}
	const pkgName = "unix"
	pkg, ok := pkgs[pkgName]
	if !ok {
		log.Fatalf("package name %q not found!", pkgName)
	}
	reg := newRegistry(pkg)
	reg.build()
}

type (
	// registry keeps the files merge is interested in grouped by their name:
	// - <name>_<goos>_<goarch>.go are the interesting files
	// - <name>_<goos>.go is the name of the file that will contain the merged objects and points at a *gofile
	registry map[string]*gofile
	// kinds holds the objects (const, type, func) that will be factored out.
	// Assumption: all files have the objects defined in the same order and with the same layout.
	kinds struct {
		consts [][]*ast.ValueSpec // Constants
	}
	// goarch represents an arch file.
	goarch struct {
		*ast.File
		kinds
	}
	// gofile holds all the arch dependent files for a given interesting file.
	gofile struct {
		arch map[string]*goarch
		kinds
	}
)

func (k *kinds) pushConst(decl *ast.GenDecl) {
	k.consts = make([][]*ast.ValueSpec, 0, len(decl.Specs))
	for _, spec := range decl.Specs {
		v := spec.(*ast.ValueSpec)
		s := make([]*ast.ValueSpec, 0, len(v.Names))
		for _, id := range v.Names {
			if isExported(id) {
				s = append(s, v)
			}
		}
		if len(s) > 0 {
			k.consts = append(k.consts, s)
		}
	}
}

func (k *kinds) merge(kk *kinds) {
	if len(kk.consts) > 0 {
		if k.consts == nil {
			// Clone the first kinds.
			k.consts = make([][]*ast.ValueSpec, len(kk.consts))
			for i, s := range kk.consts {
				k.consts[i] = append([]*ast.ValueSpec{}, s...)
			}
			return
		}
		// Intersection of k.consts and kk.consts.
		for i, kspecs := range k.consts {
			_ = kspecs
			_ = kkspecs
		}
	}
}

func newRegistry(pkg *ast.Package) registry {
	reg := make(map[string]*gofile)

	// Group files by name_os and arch.
	for fname, file := range pkg.Files {
		// Skip files not of the form: <name>_<goos>_<goarch>.go
		if strings.Count(fname, "_") != 2 {
			continue
		}
		// Group by name_os and arch.
		i := strings.LastIndexByte(fname, '_')
		name := fname[:i] + ".go" // add the file extension now
		arch := fname[i+1:]

		if reg[name] == nil {
			reg[name] = &gofile{arch: make(map[string]*goarch)}
		}
		reg[name].arch[arch] = &goarch{File: file}
	}

	return registry(reg)
}

// String lists the merged files and their architectures. Used for debugging.
func (r registry) String() string {
	var b strings.Builder
	var fnames []string
	for file, gf := range r {
		b.WriteString(file)
		b.WriteByte('\n')
		for k := range gf.arch {
			fnames = append(fnames, k)
		}
		sort.Strings(fnames)
		for _, fname := range fnames {
			b.WriteString(fmt.Sprintf("  %s\n", fname))
		}
		fnames = fnames[:0]
	}
	return b.String()
}

// build populates the kinds of every file and arch.
func (r registry) build() {
	r.build_arch()
}

// build_arch populates the kinds of every arch.
func (r registry) build_arch() {
	// Extract all constants for all arch.
	for _, gf := range r {
		for _, ga := range gf.arch {
			k := &ga.kinds
			for _, decl := range ga.File.Decls {
				switch d := decl.(type) {
				case *ast.GenDecl:
					if d.Tok != token.CONST {
						continue
					}
					// Constant found.
					k.pushConst(d)
				}
			}
		}
	}
	// Build the intersection of all constants for all arch.
	for fn, gf := range r {
		k := &gf.kinds
		for an, ga := range gf.arch {
			fmt.Println(fn, an)
			k.merge(&ga.kinds)
		}
	}
}

func isExported(id *ast.Ident) bool {
	name := []rune(id.Name)
	return unicode.IsUpper(name[0])
}
