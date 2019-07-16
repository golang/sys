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

//-----------------------------------------------------------------------------
// ast.Spec

// specCache is used to cache the results of printer.Fprint as it is pretty slow.
var specCache = map[ast.Spec]string{}

func specFromCache(spec ast.Spec) string {
	if s, ok := specCache[spec]; ok {
		return s
	}
	s := specToString(spec)
	specCache[spec] = s
	return s
}

func specToString(spec ast.Spec) string {
	var buf strings.Builder
	// The node comes from successfully parsed code, so it should be safe to ignore the error.
	_ = printer.Fprint(&buf, token.NewFileSet(), spec)
	return buf.String()
}

func specEqual(a, b ast.Spec) bool {
	return specFromCache(a) == specFromCache(b)
}

func specIn(item ast.Spec, s []ast.Spec) bool {
	for _, v := range s {
		if specEqual(v, item) {
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

// declCache is used to cache the results of printer.Fprint as it is pretty slow.
var declCache = map[ast.Decl]string{}

func declFromCache(decl ast.Decl) string {
	if s, ok := declCache[decl]; ok {
		return s
	}
	s := declToString(decl)
	declCache[decl] = s
	return s
}

func declToString(decl ast.Decl) string {
	var buf strings.Builder
	// The node comes from successfully parsed code, so it should be safe to ignore the error.
	_ = printer.Fprint(&buf, token.NewFileSet(), decl)
	return buf.String()
}

func declEqual(a, b ast.Decl) bool {
	return declFromCache(a) == declFromCache(b)
}

func declIn(item ast.Decl, s []ast.Decl) bool {
	for _, v := range s {
		if declEqual(v, item) {
			return true
		}
	}
	return false
}

func declInter(a, b []ast.Decl) []ast.Decl {
	var s []ast.Decl
	for _, v := range a {
		for _, w := range b {
			if declEqual(v, w) {
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
