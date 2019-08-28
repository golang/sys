// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package consolidate

import (
	"go/ast"
	"go/printer"
	"go/token"
	"strings"
)

var emptyFileSet = token.NewFileSet()

//-----------------------------------------------------------------------------
// ast.Expr

// Cache for the results of printer.Fprint.
var exprCache = map[ast.Expr]string{}

func exprFromCache(e ast.Expr) string {
	if s, ok := exprCache[e]; ok {
		return s
	}
	s := exprToString(e)
	exprCache[e] = s
	return s
}

func exprToString(e ast.Expr) string {
	if e == nil {
		// Common case.
		return ""
	}
	var buf strings.Builder
	// The node comes from successfully parsed code, so it should be safe to ignore the error.
	_ = printer.Fprint(&buf, emptyFileSet, e)
	return buf.String()
}

func exprEqual(a, b ast.Expr) bool {
	return exprFromCache(a) == exprFromCache(b)
}

//-----------------------------------------------------------------------------
// ast.ValueSpec

func identIn(item *ast.Ident, s []*ast.Ident) bool {
	for _, v := range s {
		if v.Name == item.Name {
			return true
		}
	}
	return false
}

func identAt(item *ast.Ident, s []*ast.Ident) int {
	for i, v := range s {
		if v.Name == item.Name {
			return i
		}
	}
	return -1
}

// Remove values in a that are not in b: must have the same name and value.
// a is mutated.
func valInter(a, b *ast.ValueSpec) {
	for i := 0; i < len(a.Names); {
		j := identAt(a.Names[i], b.Names)
		if j >= 0 && exprEqual(a.Values[i], b.Values[j]) {
			// Same name and value.
			i++
			continue
		}
		// Value in a not in b or they have different values: remove.
		valDelAt(a, i)
	}
}

// Remove values in a that are in b (same name only). a is mutated.
func valDiff(a, b *ast.ValueSpec) {
	for i := 0; i < len(a.Names); {
		if !identIn(a.Names[i], b.Names) {
			i++
			continue
		}
		valDelAt(a, i)
	}
}

// Remove the item at index i (identifier and any corresponding value if it exists).
func valDelAt(val *ast.ValueSpec, i int) {
	val.Names = identDelAt(val.Names, i)
	if i < len(val.Values) {
		val.Values = exprDelAt(val.Values, i)
	}
}

func identDelAt(s []*ast.Ident, i int) []*ast.Ident {
	if i+1 < len(s) {
		copy(s[i:], s[i+1:])
	}
	s[len(s)-1] = nil
	return s[:len(s)-1]
}

func exprDelAt(s []ast.Expr, i int) []ast.Expr {
	if i+1 < len(s) {
		copy(s[i:], s[i+1:])
	}
	s[len(s)-1] = nil
	return s[:len(s)-1]
}

//-----------------------------------------------------------------------------
// ast.TypeSpec

// Cache for the results of printer.Fprint.
var typeCache = map[*ast.TypeSpec]string{}

func typeFromCache(spec *ast.TypeSpec) string {
	if s, ok := typeCache[spec]; ok {
		return s
	}
	s := typeToString(spec)
	typeCache[spec] = s
	return s
}

func typeToString(spec *ast.TypeSpec) string {
	var buf strings.Builder
	// The node comes from successfully parsed code, so it should be safe to ignore the error.
	_ = printer.Fprint(&buf, emptyFileSet, spec)
	return buf.String()
}

func typeEqual(a, b *ast.TypeSpec) bool {
	return typeFromCache(a) == typeFromCache(b)
}

func typeIn(item *ast.TypeSpec, s []*ast.TypeSpec) bool {
	for _, v := range s {
		if typeEqual(v, item) {
			return true
		}
	}
	return false
}

func typeInter(a, b []*ast.TypeSpec) []*ast.TypeSpec {
	var s []*ast.TypeSpec
	for _, v := range a {
		if typeIn(v, b) {
			s = append(s, v)
		}
	}
	return s
}

func typeDiff(a, b []*ast.TypeSpec) []*ast.TypeSpec {
	var s []*ast.TypeSpec
	for _, v := range a {
		if !typeIn(v, b) {
			s = append(s, v)
		}
	}
	return s
}

func typeDelAt(s []ast.Spec, i int) []ast.Spec {
	if i+1 < len(s) {
		copy(s[i:], s[i+1:])
	}
	s[len(s)-1] = nil
	return s[:len(s)-1]
}

//-----------------------------------------------------------------------------
// ast.FuncDecl

// Cache for the results of printer.Fprint.
var funcCache = map[*ast.FuncDecl]string{}

func funcFromCache(decl *ast.FuncDecl) string {
	if s, ok := funcCache[decl]; ok {
		return s
	}
	s := funcToString(decl)
	funcCache[decl] = s
	return s
}

func funcToString(decl *ast.FuncDecl) string {
	var buf strings.Builder
	// The node comes from successfully parsed code, so it should be safe to ignore the error.
	_ = printer.Fprint(&buf, emptyFileSet, decl)
	return buf.String()
}

func funcEqual(a, b *ast.FuncDecl) bool {
	return funcFromCache(a) == funcFromCache(b)
}

func funcIn(item *ast.FuncDecl, s []*ast.FuncDecl) bool {
	for _, v := range s {
		if funcEqual(v, item) {
			return true
		}
	}
	return false
}

func funcInter(a, b []*ast.FuncDecl) []*ast.FuncDecl {
	var s []*ast.FuncDecl
	for _, v := range a {
		for _, w := range b {
			if funcEqual(v, w) {
				s = append(s, v)
			}
		}
	}
	return s
}

func funcDiff(a, b []*ast.FuncDecl) []*ast.FuncDecl {
	var s []*ast.FuncDecl
	for _, v := range a {
		if !funcIn(v, b) {
			s = append(s, v)
		}
	}
	return s
}

//-----------------------------------------------------------------------------

func declDelAt(s []ast.Decl, i int) []ast.Decl {
	if i+1 < len(s) {
		copy(s[i:], s[i+1:])
	}
	s[len(s)-1] = nil
	return s[:len(s)-1]
}

//-----------------------------------------------------------------------------

type visitor func(ast.Node) bool

func (v visitor) Visit(node ast.Node) ast.Visitor {
	if v(node) {
		return nil
	}
	return v
}
