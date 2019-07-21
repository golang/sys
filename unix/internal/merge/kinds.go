// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package merge

import (
	"go/ast"
	"go/printer"
	"go/token"
	"io"
	"strconv"
)

type (
	// kinds holds the objects (const, type, func) that will be consolidated.
	kinds struct {
		consts []*ast.ValueSpec // Constants grouped by Type
		types  []*ast.TypeSpec  // Types
		funcs  []*ast.FuncDecl  // Functions
	}
)

// Add the new const to the flatten list of all const definitions.
// Only the name of the consts is considered as they are unique for a given GOOS/GOARCH.
func (k *kinds) pushConst(decl *ast.GenDecl) {
loop:
	for _, spec := range decl.Specs {
		val := spec.(*ast.ValueSpec)
		for _, v := range k.consts {
			if exprEqual(v.Type, val.Type) {
				// Existing Type.
				v.Names = append(v.Names, val.Names...)
				// In case the values are not defined for all names (not seen yet).
				if n := len(val.Names) - len(val.Values); n > 0 {
					v.Values = append(v.Values, make([]ast.Expr, n)...)
				}
				v.Values = append(v.Values, val.Values...)
				continue loop
			}
		}
		// New Type: clone the slices as they will be modified by constDiff().
		v := &ast.ValueSpec{
			Names:  append([]*ast.Ident{}, val.Names...),
			Values: append([]ast.Expr{}, val.Values...),
		}
		k.consts = append(k.consts, v)
	}
}

// Add the new type to the flatten list of all type definitions.
func (k *kinds) pushType(decl *ast.GenDecl) {
	for _, spec := range decl.Specs {
		k.types = append(k.types, spec.(*ast.TypeSpec))
	}
}

// Add the new func to the flatten list of all func definitions.
func (k *kinds) pushFunc(decl *ast.FuncDecl) {
	k.funcs = append(k.funcs, decl)
}

// Intersection of all objects from kk into k.
func (k *kinds) inter(kk *kinds) {
	k.interConst(kk)
	k.interType(kk)
	k.interFunc(kk)
}

// Delete the value at index i if it is empty.
func (k *kinds) constDelAt(i int) int {
	if len(k.consts[i].Names) > 0 {
		return i
	}
	if i+1 < len(k.consts) {
		copy(k.consts[i:], k.consts[i+1:])
	}
	k.consts[len(k.consts)-1] = nil
	k.consts = k.consts[:len(k.consts)-1]
	return i - 1
}

// Intersection of constants in k and kk.
func (k *kinds) interConst(kk *kinds) {
	if len(kk.consts) == 0 {
		return
	}
	if len(k.consts) == 0 {
		// Clone the first kinds.
		for _, v := range kk.consts {
			val := &ast.ValueSpec{
				Names:  append([]*ast.Ident{}, v.Names...),
				Values: append([]ast.Expr{}, v.Values...),
			}
			k.consts = append(k.consts, val)
		}
		return
	}
	// By Type, remove values in k.consts that are not in kk.consts.
	for i := 0; i < len(k.consts); i++ {
		val := k.consts[i]
		for _, v := range kk.consts {
			if exprEqual(v.Type, val.Type) {
				valInter(val, v)
				i = k.constDelAt(i)
				break
			}
		}
	}
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
	k.types = typeInter(k.types, kk.types)
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
	k.funcs = funcInter(k.funcs, kk.funcs)
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
	// By Type, remove values in k.consts that are in kk.consts.
	for i := 0; i < len(k.consts); i++ {
		val := k.consts[i]
		for _, v := range kk.consts {
			if exprEqual(v.Type, val.Type) {
				valDiff(k.consts[i], v)
				i = k.constDelAt(i)
				break
			}
		}
	}
}

// Difference of k.types and kk.types.
func (k *kinds) diffType(kk *kinds) {
	if len(k.types) == 0 || len(kk.types) == 0 {
		return
	}
	k.types = typeDiff(k.types, kk.types)
}

// Difference of k.funcs and kk.funcs.
func (k *kinds) diffFunc(kk *kinds) {
	if len(k.funcs) == 0 || len(kk.funcs) == 0 {
		return
	}
	k.funcs = funcDiff(k.funcs, kk.funcs)
}

// hasImport returns whether or not the import is found in any of the kinds' objects.
func (k *kinds) hasImport(name string) (found bool) {
	visit := visitor(func(node ast.Node) bool {
		if found {
			return true
		}
		s, ok := node.(*ast.SelectorExpr)
		if !ok {
			return false
		}
		id, ok := s.X.(*ast.Ident)
		if !ok {
			return false
		}
		if id.Obj != nil || id.Name != name {
			return false
		}
		found = true
		return true
	})
	name, _ = strconv.Unquote(name)
	for _, node := range k.consts {
		ast.Walk(visit, node)
		if found {
			return true
		}
	}
	for _, node := range k.types {
		ast.Walk(visit, node)
		if found {
			return true
		}
	}
	for _, node := range k.funcs {
		ast.Walk(visit, node)
		if found {
			return true
		}
	}
	return false
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
		Specs:  make([]ast.Spec, 0, len(k.consts)),
	}
	// Make sure that constants are on their own line.
	for _, v := range k.consts {
		for i := range v.Names {
			spec := &ast.ValueSpec{
				Names:  v.Names[i : i+1],
				Values: v.Values[i : i+1],
			}
			node.Specs = append(node.Specs, spec)
		}
	}
	return printer.Fprint(w, emptyFileSet, node)
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
	return printer.Fprint(w, emptyFileSet, node)
}

func (k *kinds) printFunc(w io.Writer) error {
	for _, f := range k.funcs {
		if err := printer.Fprint(w, emptyFileSet, f); err != nil {
			return err
		}
		if _, err := w.Write([]byte("\n")); err != nil {
			return err
		}
	}
	return nil
}
