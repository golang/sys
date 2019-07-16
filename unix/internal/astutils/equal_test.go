// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package astutils

import (
	"fmt"
	"go/ast"
	"go/parser"
	"testing"
)

func TestIniCheck(t *testing.T) {
	type typ1 struct{}
	type typ2 struct{}
	for _, tc := range []struct {
		label string
		a, b  interface{}
		res   iCheck
	}{
		{"nil interface /nil interface", nil, nil, same},
		{"nil/nil", (*typ1)(nil), (*typ1)(nil), same},
		{"nil/nil interface", (*typ1)(nil), nil, different},
		{"nil/not nil", (*typ1)(nil), &typ1{}, different},
		{"not nil/not nil different", &typ1{}, &typ2{}, different},
		{"not nil/not nil same", &typ1{}, &typ1{}, unknown},
	} {
		t.Run(tc.label, func(t *testing.T) {
			if got, want := iniCheck(tc.a, tc.b), tc.res; got != want {
				t.Fatalf("got %v; want %v", got, want)
			}
		})
	}
}

func TestIdentEqual(t *testing.T) {
	for _, tc := range []struct {
		label string
		a, b  *ast.Ident
		ok    bool
	}{
		{"a a", ast.NewIdent("a"), ast.NewIdent("a"), true},
		{"a b", ast.NewIdent("a"), ast.NewIdent("b"), false},
	} {
		t.Run(tc.label, func(t *testing.T) {
			if got, want := IdentEqual(tc.a, tc.b), tc.ok; got != want {
				t.Fatalf("got %v; want %v", got, want)
			}
			if got, want := IdentEqual(tc.b, tc.a), tc.ok; got != want {
				t.Fatalf("got %v; want %v", got, want)
			}
		})
	}
}

func TestIdentMultiEqual(t *testing.T) {
	toSlice := func(s ...string) (res []*ast.Ident) {
		for _, s := range s {
			res = append(res, ast.NewIdent(s))
		}
		return
	}

	for _, tc := range []struct {
		label string
		a, b  []*ast.Ident
		ok    bool
	}{
		{"a a", toSlice("a"), toSlice("a"), true},
		{"a b", toSlice("a"), toSlice("b"), false},
		{"a,a a,a", toSlice("a", "a"), toSlice("a", "a"), true},
		{"a,b a,a", toSlice("a", "b"), toSlice("a", "a"), false},
	} {
		t.Run(tc.label, func(t *testing.T) {
			if got, want := IdentMultiEqual(tc.a, tc.b), tc.ok; got != want {
				t.Fatalf("got %v; want %v", got, want)
			}
			if got, want := IdentMultiEqual(tc.b, tc.a), tc.ok; got != want {
				t.Fatalf("got %v; want %v", got, want)
			}
		})
	}
}

func TestFieldListEqual(t *testing.T) {
	toList := func(s ...string) *ast.FieldList {
		res := &ast.FieldList{}
		for _, s := range s {
			field := &ast.Field{}
			field.Names = append(field.Names, ast.NewIdent(s))
			res.List = append(res.List, field)
		}
		return res
	}

	for _, tc := range []struct {
		label string
		a, b  *ast.FieldList
		ok    bool
	}{
		{"a a", toList("a"), toList("a"), true},
		{"a b", toList("a"), toList("b"), false},
	} {
		t.Run(tc.label, func(t *testing.T) {
			if got, want := FieldListEqual(tc.a, tc.b), tc.ok; got != want {
				t.Fatalf("got %v; want %v", got, want)
			}
			if got, want := FieldListEqual(tc.b, tc.a), tc.ok; got != want {
				t.Fatalf("got %v; want %v", got, want)
			}
		})
	}
}

func TestDeclEqual(t *testing.T) {
	for _, tc := range []struct {
		label string
		a, b  ast.Decl
		ok    bool
	}{
		{"bad nil", &ast.BadDecl{}, nil, false},
		{"bad bad", &ast.BadDecl{}, &ast.BadDecl{}, true},
		{"func 1 nil", f1, nil, false},
		{"func 1 1", f1, f1, true},
		{"func 1 2", f1, f2, false},
		{"gen 1 nil", g1, nil, false},
		{"gen 1 1", g1, g1, true},
		{"gen 1 2", g1, g2, false},
	} {
		t.Run(tc.label, func(t *testing.T) {
			if got, want := DeclEqual(tc.a, tc.b), tc.ok; got != want {
				t.Fatalf("got %v; want %v", got, want)
			}
			if got, want := DeclEqual(tc.b, tc.a), tc.ok; got != want {
				t.Fatalf("got %v; want %v", got, want)
			}
		})
	}
}

func TestExprEqual(t *testing.T) {
	for _, tc := range []struct {
		a, b string
		ok   bool
	}{
		{"1 == 1", "1 == 1", true},
		{"a == a", "a == a", true},
		{"a == a", "a == b", false},
		{"len([]int{2}) > 1", "len([]int{2}) > 1", true},
		{"make(chan int)", "make(chan int)", true},
		{"make(<-chan int)", "make(<-chan int)", true},
		{"make(<-chan int)", "make(chan<- int)", false},
		{"[...]int{1} == [...]int{1}", "[...]int{1} == [...]int{1}", true},
		{"[]int{1} == []int{1}", "[]int{1} == []int{1}", true},
		{"func(){}()", "func(){}()", true},
		{"func(a int){}(1)", "func(a int){}(1)", true},
		{"func(a int){}(1)", "func(b int){}(1)", false},
		{"func(a int){}(1)", "func(){}(1)", false},
		{"func(){}(int)(1)", "func(){}(int)(1)", true},
		{"func(){}(int)(1)", "func()(a int){}(1)", false},
		{"func(){}()", "func(){ _ = 1 }()", false},
		{"make(map[int]int)", "make(map[int]int)", true},
		{"make(map[int]int)", "make(map[int]uint)", false},
		{"make(map[int]int)", "make(map[uint]int)", false},
		{"map[int]int{1: 2}", "map[int]int{1: 2}", true},
		{"map[int]int{1: 2}", "map[int]int{1: 3}", false},
		{"map[int]int{1: 2}", "map[int]int{3: 2}", false},
		{`"a"[0]`, `"a"[0]`, true},
		{`"a"[0]`, `"a"[1]`, false},
		{`"a"[0]`, `"b"[0]`, false},
		{"struct{string}{``} == 1", "struct{string}{``} == 1", true},
		{"struct{string}{``} == 1", "struct{int}{0} == 1", false},
		{"struct{string}{``} == 1", "struct{string;int}{``,0} == 1", false},
		{"io.Reader == io.Reader", "io.Reader == io.Reader", true},
		{"io.Reader == io.Reader", "io.Reader == io.Writer", false},
		{"[]int{} == 0", "[]int{} == 0", true},
	} {
		label := fmt.Sprintf("%q / %q", tc.a, tc.b)
		t.Run(label, func(t *testing.T) {
			a, err := parser.ParseExpr(tc.a)
			if err != nil {
				t.Fatal(err)
			}
			b, err := parser.ParseExpr(tc.b)
			if err != nil {
				t.Fatal(err)
			}
			if got, want := ExprEqual(a, b), tc.ok; got != want {
				t.Fatalf("got %v; want %v", got, want)
			}
			if got, want := ExprEqual(b, a), tc.ok; got != want {
				t.Fatalf("got %v; want %v", got, want)
			}
		})
	}
}

func TestExprMultiEqual(t *testing.T) {

}

func TestSpecEqual(t *testing.T) {

}

func TestSpecMultiEqual(t *testing.T) {

}

func TestStmtEqual(t *testing.T) {

}

func TestStmtMultiEqual(t *testing.T) {

}
