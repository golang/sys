// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package astutils

import "go/ast"

// InterValueSpec performs the intersection of a and b. a is mutated.
//
// a is unchanged if b and a do not have the same type.
func InterValueSpec(a, b *ast.ValueSpec) {
	if a == nil || b == nil {
		return
	}
	if !ExprEqual(a.Type, b.Type) {
		return
	}
loop:
	for i := 0; i < len(a.Names); {
		id := a.Names[i]
		for ki, kid := range b.Names {
			if IdentEqual(id, kid) && ExprEqual(a.Values[i], b.Values[ki]) {
				i++
				continue loop
			}
		}
		DelValueSpecAt(a, i)
	}
}

// Returns the intersection of a and b. a is mutated.
//
// a is unchanged if b is empty.
func InterTypeSpec(a, b []*ast.TypeSpec) []*ast.TypeSpec {
	if len(a) == 0 || len(b) == 0 {
		return a
	}
loop:
	for i := 0; i < len(a); {
		s := a[i]
		for _, ks := range b {
			if SpecEqual(s, ks) {
				i++
				continue loop
			}
		}
		a = DelTypeSpecAt(a, i)
	}
	return a
}

// Returns the intersection of a and b. a is mutated.
//
// a is unchanged if b is empty.
func InterFuncDecl(a, b []*ast.FuncDecl) []*ast.FuncDecl {
	if len(a) == 0 || len(b) == 0 {
		return a
	}
loop:
	for i := 0; i < len(a); {
		s := a[i]
		for _, ks := range b {
			if DeclEqual(s, ks) {
				i++
				continue loop
			}
		}
		a = DelFuncDeclAt(a, i)
	}
	return a
}
