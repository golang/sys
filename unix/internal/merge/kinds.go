// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package merge

import (
	"go/ast"
	"go/printer"
	"go/token"
	"io"
)

type (
	// kinds holds the objects (const, type, func) that will be factored out.
	// Assumption: all files have the objects defined in the same order and with the same layout.
	kinds struct {
		consts []ast.Spec // Constants
		types  []ast.Spec // Types
		funcs  []ast.Decl // Functions
	}
)

// Add the new const to the flatten list of all const definitions.
func (k *kinds) pushConst(decl *ast.GenDecl) {
	k.consts = specUnion(k.consts, decl.Specs)
}

// Add the new type to the flatten list of all type definitions.
func (k *kinds) pushType(decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		s := spec.(*ast.TypeSpec)
		ns := &ast.TypeSpec{
			Name: s.Name,
			Type: s.Type,
		}
		k.types = append(k.types, ns)
	}
}

// Add the new type to the flatten list of all type definitions.
func (k *kinds) pushFunc(decl *ast.FuncDecl) {
	k.funcs = append(k.funcs, decl)
}

// Intersection of all objects from kk into k.
func (k *kinds) inter(kk *kinds) {
	k.interConst(kk)
	k.interType(kk)
	k.interFunc(kk)
}

// Intersection of constants in k and kk.
func (k *kinds) interConst(kk *kinds) {
	if len(kk.consts) == 0 {
		return
	}
	if len(k.consts) == 0 {
		// Clone the first kinds.
		k.consts = append(k.consts, kk.consts...)
		return
	}
	k.consts = specInter(k.consts, kk.consts)
}

// Intersection of types in k and kk.
func (k *kinds) interType(kk *kinds) {
	if len(kk.types) == 0 {
		return
	}
	if len(k.types) == 0 {
		// Clone the first kinds.
		k.types = append([]ast.Spec{}, kk.types...)
		return
	}
	k.types = specInter(k.types, kk.types)
}

// Intersection of functions in k and kk.
func (k *kinds) interFunc(kk *kinds) {
	if len(kk.funcs) == 0 {
		return
	}
	if len(k.funcs) == 0 {
		// Clone the first kinds.
		k.funcs = append([]ast.Decl{}, kk.funcs...)
		return
	}
	k.funcs = declInter(k.funcs, kk.funcs)
}

// Difference of all objects from kk into k.
func (k *kinds) diff(kk *kinds) {
	k.diffConst(kk)
	k.diffType(kk)
	k.diffFunc(kk)
}

// Difference of k.consts and kk.consts.
func (k *kinds) diffConst(kk *kinds) {
	if len(k.consts) == 0 || len(kk.consts) == 0 {
		return
	}
	k.consts = specDiff(k.consts, kk.consts)
}

// Difference of k.types and kk.types.
func (k *kinds) diffType(kk *kinds) {
	if len(k.types) == 0 || len(kk.types) == 0 {
		return
	}
	k.types = specDiff(k.types, kk.types)
}

// Difference of k.funcs and kk.funcs.
func (k *kinds) diffFunc(kk *kinds) {
	if len(k.funcs) == 0 || len(kk.funcs) == 0 {
		return
	}
	k.funcs = declDiff(k.funcs, kk.funcs)
}

func (k *kinds) print(w io.Writer) error {
	if err := k.printConst(w); err != nil {
		return err
	}
	// Add a separator at the end of constant definitions to avoid invalid source code.
	if _, err := w.Write([]byte("\n")); err != nil {
		return err
	}
	if err := k.printType(w); err != nil {
		return err
	}
	if err := k.printFunc(w); err != nil {
		return err
	}
	return nil
}

func (k *kinds) printConst(w io.Writer) error {
	if len(k.consts) == 0 {
		// No constant.
		return nil
	}
	node := &ast.GenDecl{
		Lparen: 1, // Make sure there is a parenthesis
		Tok:    token.CONST,
		Specs:  k.consts,
	}
	return printer.Fprint(w, token.NewFileSet(), node)
}

func (k *kinds) printType(w io.Writer) error {
	if len(k.types) == 0 {
		// No type.
		return nil
	}
	node := &ast.GenDecl{
		Lparen: 1, // Make sure there is a parenthesis
		Tok:    token.TYPE,
		Specs:  make([]ast.Spec, len(k.types)),
	}
	for i, ts := range k.types {
		node.Specs[i] = ts
	}
	return printer.Fprint(w, token.NewFileSet(), node)
}

func (k *kinds) printFunc(w io.Writer) error {
	fset := token.NewFileSet()
	for _, f := range k.funcs {
		if err := printer.Fprint(w, fset, f); err != nil {
			return err
		}
		if _, err := w.Write([]byte("\n")); err != nil {
			return err
		}
	}
	return nil
}
