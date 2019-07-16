// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package merge

import (
	"go/ast"
	"go/printer"
	"go/token"
	"io"

	"golang.org/x/sys/unix/internal/astutils"
)

type (
	// kinds holds the objects (const, type, func) that will be factored out.
	// Assumption: all files have the objects defined in the same order and with the same layout.
	kinds struct {
		consts *ast.ValueSpec  // Constants
		types  []*ast.TypeSpec // Types
		funcs  []*ast.FuncDecl // Functions
	}
)

// Add the new const to the flatten list of all const definitions.
func (k *kinds) pushConst(decl *ast.GenDecl) {
	if k.consts == nil {
		k.consts = &ast.ValueSpec{}
	}
	for _, s := range decl.Specs {
		v := s.(*ast.ValueSpec)
		k.consts.Names = append(k.consts.Names, v.Names...)
		k.consts.Values = append(k.consts.Values, v.Values...)
	}
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
	if kk.consts == nil {
		return
	}
	if k.consts == nil {
		// Clone the first kinds.
		k.consts = &ast.ValueSpec{
			Names:  append([]*ast.Ident{}, kk.consts.Names...),
			Values: append([]ast.Expr{}, kk.consts.Values...),
		}
		return
	}
	astutils.InterValueSpec(k.consts, kk.consts)
}

// Intersection of types in k and kk.
func (k *kinds) interType(kk *kinds) {
	if len(kk.types) == 0 {
		return
	}
	if len(k.types) == 0 {
		// Clone the first kinds.
		k.types = append([]*ast.TypeSpec{}, kk.types...)
		return
	}
	k.types = astutils.InterTypeSpec(k.types, kk.types)
}

// Intersection of functions in k and kk.
func (k *kinds) interFunc(kk *kinds) {
	if len(kk.funcs) == 0 {
		return
	}
	if len(k.funcs) == 0 {
		// Clone the first kinds.
		k.funcs = append([]*ast.FuncDecl{}, kk.funcs...)
		return
	}
	k.funcs = astutils.InterFuncDecl(k.funcs, kk.funcs)
}

// Difference of all objects from kk into k.
func (k *kinds) diff(kk *kinds) {
	k.diffConst(kk)
	k.diffType(kk)
	k.diffFunc(kk)
}

// Difference of k.consts and kk.consts.
func (k *kinds) diffConst(kk *kinds) {
	if k.consts == nil || kk.consts == nil {
		return
	}
loop:
	for i := 0; i < len(k.consts.Names); {
		id := k.consts.Names[i]
		for ki, kid := range kk.consts.Names {
			if astutils.IdentEqual(id, kid) && astutils.ExprEqual(k.consts.Values[i], kk.consts.Values[ki]) {
				// Constant is in k and kk: factorized, remove from k.
				astutils.DelValueSpecAt(k.consts, i)
				continue loop
			}
		}
		i++
	}
}

// Difference of k.types and kk.types.
func (k *kinds) diffType(kk *kinds) {
	if len(k.types) == 0 || len(kk.types) == 0 {
		return
	}
loop:
	for i := 0; i < len(k.types); {
		t := k.types[i]
		for _, kt := range kk.types {
			if astutils.SpecEqual(t, kt) {
				// Type is in k and kk: factorized, remove from k.
				k.types = astutils.DelTypeSpecAt(k.types, i)
				continue loop
			}
		}
		i++
	}
}

// Difference of k.funcs and kk.funcs.
func (k *kinds) diffFunc(kk *kinds) {
	if len(k.funcs) == 0 || len(kk.funcs) == 0 {
		return
	}
loop:
	for i := 0; i < len(k.funcs); {
		f := k.funcs[i]
		for _, kf := range kk.funcs {
			if astutils.DeclEqual(f, kf) {
				// Function is in k and kk: factorized, remove from k.
				k.funcs = astutils.DelFuncDeclAt(k.funcs, i)
				continue loop
			}
		}
		i++
	}
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
	if k.consts == nil || len(k.consts.Names) == 0 {
		// No constant.
		return nil
	}
	node := &ast.GenDecl{
		Lparen: 1, // Make sure there is a parenthesis
		Tok:    token.CONST,
		Specs:  make([]ast.Spec, len(k.consts.Names)),
	}
	// Convert single line constant definitions into multiple lines, one per definition.
	for i, id := range k.consts.Names {
		node.Specs[i] = &ast.ValueSpec{
			Type:   k.consts.Type,
			Names:  []*ast.Ident{id},
			Values: []ast.Expr{k.consts.Values[i]},
		}
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
