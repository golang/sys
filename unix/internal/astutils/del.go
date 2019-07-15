// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package astutils

import (
	"go/ast"
)

// DelValueSpecAt removes the value spec at index i.
func DelValueSpecAt(v *ast.ValueSpec, i int) {
	if i+1 < len(v.Names) {
		copy(v.Names[i:], v.Names[i+1:])
	}
	v.Names[len(v.Names)-1] = nil
	v.Names = v.Names[:len(v.Names)-1]
	if i >= len(v.Values) {
		return
	}
	if i+1 < len(v.Values) {
		copy(v.Values[i:], v.Values[i+1:])
	}
	v.Values[len(v.Values)-1] = nil
	v.Values = v.Values[:len(v.Values)-1]
}

// DelTypeSpecAt removes the type spec at index i.
func DelTypeSpecAt(s []*ast.TypeSpec, i int) []*ast.TypeSpec {
	if i+1 < len(s) {
		copy(s[i:], s[i+1:])
	}
	s[len(s)-1] = nil
	return s[:len(s)-1]
}

// DelFuncDeclAt removes the func decl at index i.
func DelFuncDeclAt(s []*ast.FuncDecl, i int) []*ast.FuncDecl {
	if i+1 < len(s) {
		copy(s[i:], s[i+1:])
	}
	s[len(s)-1] = nil
	return s[:len(s)-1]
}

// DelDeclAt removes the decl at index i and returns the modified slice.
func DelDeclAt(s []ast.Decl, i int) []ast.Decl {
	if i+1 < len(s) {
		copy(s[i:], s[i+1:])
	}
	s[len(s)-1] = nil
	return s[:len(s)-1]
}

// DelSpecAt removes the spec at index i and returns the modified slice.
func DelSpecAt(s []ast.Spec, i int) []ast.Spec {
	if i+1 < len(s) {
		copy(s[i:], s[i+1:])
	}
	s[len(s)-1] = nil
	return s[:len(s)-1]
}
