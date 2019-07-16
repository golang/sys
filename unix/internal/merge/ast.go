// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package merge

import (
	"go/ast"
	"go/printer"
	"go/token"
	"strings"
)

// astCache is used to cache the results of printer.Fprint as it is pretty slow.
var astCache = map[ast.Node]string{}

func nodeFromCache(node ast.Node) string {
	if s, ok := astCache[node]; ok {
		return s
	}
	s := nodeToString(node)
	astCache[node] = s
	return s
}

func nodeToString(node interface{}) string {
	var buf strings.Builder
	// The node comes from successfully parsed code, so it should be safe to ignore the error.
	_ = printer.Fprint(&buf, token.NewFileSet(), node)
	return buf.String()
}

func astEqual(a, b ast.Node) bool {
	return nodeFromCache(a) == nodeFromCache(b)
}

//-----------------------------------------------------------------------------
// ast.Spec

func specIn(item ast.Spec, s []ast.Spec) bool {
	for _, v := range s {
		if astEqual(v, item) {
			return true
		}
	}
	return false
}

func specInter(a, b []ast.Spec) []ast.Spec {
	var s []ast.Spec
	for _, v := range a {
		if specIn(v, b) {
			s = append(s, v)
		}
	}
	return s
}

func specUnion(a, b []ast.Spec) []ast.Spec {
	if len(a) < len(b) {
		a, b = b, a
	}
	s := append([]ast.Spec{}, b...)
	for _, v := range a {
		if !specIn(v, b) {
			s = append(s, v)
		}
	}
	return s
}

func specDiff(a, b []ast.Spec) []ast.Spec {
	var s []ast.Spec
	for _, v := range a {
		if !specIn(v, b) {
			s = append(s, v)
		}
	}
	return s
}

//-----------------------------------------------------------------------------
// ast.Decl

func declIn(item ast.Decl, s []ast.Decl) bool {
	for _, v := range s {
		if astEqual(v, item) {
			return true
		}
	}
	return false
}

func declInter(a, b []ast.Decl) []ast.Decl {
	var s []ast.Decl
	for _, v := range a {
		for _, w := range b {
			if astEqual(v, w) {
				s = append(s, v)
			}
		}
	}
	return s
}

func declDiff(a, b []ast.Decl) []ast.Decl {
	var s []ast.Decl
	for _, v := range a {
		if !declIn(v, b) {
			s = append(s, v)
		}
	}
	return s
}

func delDeclAt(s []ast.Decl, i int) []ast.Decl {
	if i+1 < len(s) {
		copy(s[i:], s[i+1:])
	}
	s[len(s)-1] = nil
	return s[:len(s)-1]
}
