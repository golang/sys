// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package astutils

import (
	"fmt"
	"go/ast"
	"reflect"
)

type iCheck int

const (
	different iCheck = iota
	same
	unknown
)

// iniCheck returns one of the iCheck value.
//  - different: a and b are known to be different
//  - same: a and b are known to be the same (nil interface or nil pointer)
//  - unknown: cannot compare a and b, but both are a non nil pointer.
//
// If typed, a and b must be pointers.
func iniCheck(a, b interface{}) iCheck {
	if a == nil && b == nil {
		return same
	}
	if a == nil || b == nil {
		return different
	}
	// reflect based check for same type and if both are typed and/or nil.
	na := reflect.TypeOf(a).String()
	nb := reflect.TypeOf(b).String()
	if na != nb {
		return different
	}
	nila := !reflect.ValueOf(a).Elem().IsValid()
	nilb := !reflect.ValueOf(b).Elem().IsValid()
	if nila && nilb {
		return same
	}
	if nila || nilb {
		return different
	}
	// At this point, a and b are non nil pointers of the same type.
	return unknown
}

// ExprMultiEqual returns whether or not the a and b are deeply equal.
func ExprMultiEqual(a, b []ast.Expr) bool {
	if len(a) != len(b) {
		return false
	}
	for i, s := range a {
		if !ExprEqual(s, b[i]) {
			return false
		}
	}
	return true
}

// ExprEqual returns whether or not the a and b are deeply equal.
func ExprEqual(a, b ast.Expr) bool {
	switch iniCheck(a, b) {
	case different:
		return false
	case same:
		return true
	}

	switch a := a.(type) {
	case *ast.ArrayType:
		b := b.(*ast.ArrayType)
		return ExprEqual(a.Len, b.Len) && ExprEqual(a.Elt, b.Elt)

	case *ast.BasicLit:
		b := b.(*ast.BasicLit)
		return a.Kind == b.Kind && a.Value == b.Value

	case *ast.BinaryExpr:
		b := b.(*ast.BinaryExpr)
		return ExprEqual(a.X, b.X) && ExprEqual(a.Y, b.Y)

	case *ast.CallExpr:
		b := b.(*ast.CallExpr)
		return ExprEqual(a.Fun, b.Fun) && ExprMultiEqual(a.Args, b.Args)

	case *ast.ChanType:
		b := b.(*ast.ChanType)
		return a.Dir == b.Dir && ExprEqual(a.Value, b.Value)

	case *ast.CompositeLit:
		b := b.(*ast.CompositeLit)
		return ExprEqual(a.Type, b.Type) && ExprMultiEqual(a.Elts, b.Elts)

	case *ast.Ellipsis:
		b := b.(*ast.Ellipsis)
		return ExprEqual(a.Elt, b.Elt)

	case *ast.FuncLit:
		b := b.(*ast.FuncLit)
		return FieldListEqual(a.Type.Params, b.Type.Params) &&
			FieldListEqual(a.Type.Results, b.Type.Results) &&
			StmtMultiEqual(a.Body.List, b.Body.List)

	case *ast.Ident:
		b := b.(*ast.Ident)
		return IdentEqual(a, b)

	case *ast.IndexExpr:
		b := b.(*ast.IndexExpr)
		return ExprEqual(a.X, b.X) && ExprEqual(a.Index, b.Index)

	case *ast.InterfaceType:
		b := b.(*ast.InterfaceType)
		return FieldListEqual(a.Methods, b.Methods)

	case *ast.KeyValueExpr:
		b := b.(*ast.KeyValueExpr)
		return ExprEqual(a.Key, b.Key) && ExprEqual(a.Value, b.Value)

	case *ast.MapType:
		b := b.(*ast.MapType)
		return ExprEqual(a.Key, b.Key) && ExprEqual(a.Value, b.Value)

	case *ast.SelectorExpr:
		b := b.(*ast.SelectorExpr)
		return ExprEqual(a.Sel, b.Sel) && ExprEqual(a.X, b.X)

	case *ast.SliceExpr:
		b := b.(*ast.SliceExpr)
		return a.Slice3 == b.Slice3 &&
			ExprEqual(a.X, b.X) && ExprEqual(a.Max, b.Max) &&
			ExprEqual(a.High, b.High) && ExprEqual(a.Low, b.Low)

	case *ast.StarExpr:
		b := b.(*ast.StarExpr)
		return ExprEqual(a.X, b.X)

	case *ast.StructType:
		b := b.(*ast.StructType)
		return FieldListEqual(a.Fields, b.Fields)

	case *ast.UnaryExpr:
		b := b.(*ast.UnaryExpr)
		return ExprEqual(a.X, b.X)
	}
	panic(fmt.Sprintf("unsupported expr %T", a))
}

// FieldListEqual returns whether or not a and b are deeply equal.
func FieldListEqual(a, b *ast.FieldList) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a.List) != len(b.List) {
		return false
	}
	for i, fa := range a.List {
		fb := b.List[i]
		if !ExprEqual(fa.Type, fb.Type) {
			return false
		}
		if !ExprEqual(fa.Tag, fb.Tag) {
			return false
		}
		if !IdentMultiEqual(fa.Names, fb.Names) {
			return false
		}
	}
	return true
}

// StmtMultiEqual returns whether or not a and b are deeply equal.
func StmtMultiEqual(a, b []ast.Stmt) bool {
	if len(a) != len(b) {
		return false
	}
	for i, s := range a {
		if !StmtEqual(s, b[i]) {
			return false
		}
	}
	return true
}

// StmtEqual returns whether or not a and b are deeply equal.
func StmtEqual(a, b ast.Stmt) bool {
	switch iniCheck(a, b) {
	case different:
		return false
	case same:
		return true
	}

	switch a := a.(type) {
	case *ast.BadStmt:
		return true

	case *ast.AssignStmt:
		b := b.(*ast.AssignStmt)
		return ExprMultiEqual(a.Lhs, b.Lhs) && ExprMultiEqual(a.Rhs, b.Rhs)

	case *ast.BlockStmt:
		b := b.(*ast.BlockStmt)
		return StmtMultiEqual(a.List, b.List)

	case *ast.BranchStmt:
		b := b.(*ast.BranchStmt)
		return ExprEqual(a.Label, b.Label)

	case *ast.DeclStmt:
		b := b.(*ast.DeclStmt)
		return DeclEqual(a.Decl, b.Decl)

	case *ast.DeferStmt:
		b := b.(*ast.DeferStmt)
		return ExprEqual(a.Call, b.Call)

	case *ast.ExprStmt:
		b := b.(*ast.ExprStmt)
		return ExprEqual(a.X, b.X)

	case *ast.ForStmt:
		b := b.(*ast.ForStmt)
		return StmtEqual(a.Init, b.Init) && ExprEqual(a.Cond, b.Cond) &&
			StmtEqual(a.Post, b.Post) && StmtEqual(a.Body, b.Body)

	case *ast.GoStmt:
		b := b.(*ast.GoStmt)
		return ExprEqual(a.Call, b.Call)

	case *ast.IfStmt:
		b := b.(*ast.IfStmt)
		return StmtEqual(a.Init, b.Init) &&
			ExprEqual(a.Cond, b.Cond) && StmtEqual(a.Body, b.Body) &&
			StmtEqual(a.Else, b.Else)

	case *ast.IncDecStmt:
		b := b.(*ast.IncDecStmt)
		return ExprEqual(a.X, b.X)

	case *ast.LabeledStmt:
		b := b.(*ast.LabeledStmt)
		return StmtEqual(a.Stmt, b.Stmt)

	case *ast.RangeStmt:
		b := b.(*ast.RangeStmt)
		return ExprEqual(a.X, b.X) &&
			ExprEqual(a.Key, b.Key) && ExprEqual(a.Value, b.Value) &&
			StmtEqual(a.Body, b.Body)

	case *ast.ReturnStmt:
		b := b.(*ast.ReturnStmt)
		return ExprMultiEqual(a.Results, b.Results)

	case *ast.SelectStmt:
		b := b.(*ast.SelectStmt)
		return StmtEqual(a.Body, b.Body)

	case *ast.SendStmt:
		b := b.(*ast.SendStmt)
		return ExprEqual(a.Chan, b.Chan) && ExprEqual(a.Value, b.Value)

	case *ast.SwitchStmt:
		b := b.(*ast.SwitchStmt)
		return StmtEqual(a.Init, b.Init) && ExprEqual(a.Tag, b.Tag)

	case *ast.TypeSwitchStmt:
		b := b.(*ast.TypeSwitchStmt)
		return StmtEqual(a.Init, b.Init) && StmtEqual(a.Assign, b.Assign)
	}
	panic(fmt.Sprintf("unsupported statement %T", a))
}

// DeclEqual returns whether or not a and b are deeply equal.
func DeclEqual(a, b ast.Decl) bool {
	switch iniCheck(a, b) {
	case different:
		return false
	case same:
		return true
	}

	switch a := a.(type) {
	case *ast.BadDecl:
		return true

	case *ast.FuncDecl:
		b := b.(*ast.FuncDecl)
		if !IdentEqual(a.Name, b.Name) || !FieldListEqual(a.Recv, b.Recv) {
			return false
		}
		switch iniCheck(a.Type, b.Type) {
		case different:
			return false
		case unknown:
			// a.Type and b.Type are not nil.
			if !FieldListEqual(a.Type.Params, b.Type.Params) || !FieldListEqual(a.Type.Results, b.Type.Results) {
				return false
			}
		}
		return StmtEqual(a.Body, b.Body)

	case *ast.GenDecl:
		b := b.(*ast.GenDecl)
		return SpecMultiEqual(a.Specs, b.Specs)
	}
	panic(fmt.Sprintf("unsupported declaration %T", a))
}

// IdentMultiEqual returns whether or not a and b are deeply equal.
func IdentMultiEqual(a, b []*ast.Ident) bool {
	if len(a) != len(b) {
		return false
	}
	for i, as := range a {
		if !ExprEqual(as, b[i]) {
			return false
		}
	}
	return true
}

// IdentEqual returns whether or not a and b are equal.
func IdentEqual(a, b *ast.Ident) bool {
	return a.Name == b.Name
}

// SpecMultiEqual returns whether or not a and b are deeply equal.
func SpecMultiEqual(a, b []ast.Spec) bool {
	if len(a) != len(b) {
		return false
	}
	for i, as := range a {
		if !SpecEqual(as, b[i]) {
			return false
		}
	}
	return true
}

// SpecEqual returns whether or not a and b are deeply equal.
func SpecEqual(a, b ast.Spec) bool {
	switch iniCheck(a, b) {
	case different:
		return false
	case same:
		return true
	}

	switch a := a.(type) {
	case *ast.ImportSpec:
		b := b.(*ast.ImportSpec)
		return ExprEqual(a.Name, b.Name) && ExprEqual(a.Path, b.Path)

	case *ast.TypeSpec:
		b := b.(*ast.TypeSpec)
		return IdentEqual(a.Name, b.Name) && ExprEqual(a.Type, b.Type)

	case *ast.ValueSpec:
		b := b.(*ast.ValueSpec)
		return ExprEqual(a.Type, b.Type) && IdentMultiEqual(a.Names, b.Names) && ExprMultiEqual(a.Values, b.Values)
	}
	panic(fmt.Sprintf("unsupported spec %T", a))
}
