package astutils

import (
	"fmt"
	"go/ast"
	"reflect"
	"testing"
)

var (
	e1 = &ast.BasicLit{Value: "1"}
	e2 = &ast.BasicLit{Value: "2"}
	e3 = &ast.BasicLit{Value: "3"}
	i1 = ast.NewIdent("i1")
	i2 = ast.NewIdent("i2")
	i3 = ast.NewIdent("i3")
	v1 = &ast.ValueSpec{Names: []*ast.Ident{i1}, Values: []ast.Expr{e1}}
	v2 = &ast.ValueSpec{Names: []*ast.Ident{i1, i2}, Values: []ast.Expr{e1, e2}}
	v3 = &ast.ValueSpec{Names: []*ast.Ident{i1, i2, i3}, Values: []ast.Expr{e1, e2, e3}}
	g1 = &ast.GenDecl{Specs: []ast.Spec{v1}}
	g2 = &ast.GenDecl{Specs: []ast.Spec{v2}}
	g3 = &ast.GenDecl{Specs: []ast.Spec{v3}}
	t1 = &ast.TypeSpec{Name: ast.NewIdent("t1")}
	t2 = &ast.TypeSpec{Name: ast.NewIdent("t2")}
	t3 = &ast.TypeSpec{Name: ast.NewIdent("t3")}
	f1 = &ast.FuncDecl{Name: ast.NewIdent("f1")}
	f2 = &ast.FuncDecl{Name: ast.NewIdent("f2")}
	f3 = &ast.FuncDecl{Name: ast.NewIdent("f3")}
)

func dupValueSpec(v *ast.ValueSpec) *ast.ValueSpec {
	return &ast.ValueSpec{
		Names:  append([]*ast.Ident{}, v.Names...),
		Values: append([]ast.Expr{}, v.Values...),
	}
}

// checkDel checks that the slice b contains a and that a is one item less than b.
func checkDel(a, b interface{}) error {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)
	// a has one item less than b.
	if an, bn := av.Len(), bv.Len()-1; an != bn {
		return fmt.Errorf("invalid length: got %d; want %d", an, bn)
	}
	// b contains a.
	var num int
	for i := 0; i < av.Len(); i++ {
		ai := av.Index(i).Interface()
		for i := 0; i < bv.Len(); i++ {
			if bv.Index(i).Interface() == ai {
				num++
				break
			}
		}
	}
	if num != av.Len() {
		return fmt.Errorf("b does not contain a")
	}
	return nil
}

func TestDelDeclAt(t *testing.T) {
	toSlice := func(v ...ast.Decl) []ast.Decl { return v }

	for _, tc := range []struct {
		label string
		s     []ast.Decl
		i     int
	}{
		{"one item", toSlice(g1), 0},
		{"two items at 0", toSlice(g1, g2), 0},
		{"two items at 1", toSlice(g1, g2), 1},
		{"three items at 0", toSlice(g1, g2, g3), 0},
		{"three items at 1", toSlice(g1, g2, g3), 1},
		{"three items at 2", toSlice(g1, g2, g3), 2},
	} {
		t.Run(tc.label, func(t *testing.T) {
			s := DelDeclAt(tc.s, tc.i)
			if err := checkDel(s, tc.s); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestDelSpecAt(t *testing.T) {
	toSlice := func(v ...ast.Spec) []ast.Spec { return v }

	for _, tc := range []struct {
		label string
		s     []ast.Spec
		i     int
	}{
		{"one item", toSlice(v1), 0},
		{"two items at 0", toSlice(v1, v2), 0},
		{"two items at 1", toSlice(v1, v2), 1},
		{"three items at 0", toSlice(v1, v2, v3), 0},
		{"three items at 1", toSlice(v1, v2, v3), 1},
		{"three items at 2", toSlice(v1, v2, v3), 2},
	} {
		t.Run(tc.label, func(t *testing.T) {
			s := DelSpecAt(tc.s, tc.i)
			if err := checkDel(s, tc.s); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestTypeSpecAt(t *testing.T) {
	toSlice := func(v ...*ast.TypeSpec) []*ast.TypeSpec { return v }

	for _, tc := range []struct {
		label string
		s     []*ast.TypeSpec
		i     int
	}{
		{"one item", toSlice(t1), 0},
		{"two items at 0", toSlice(t1, t2), 0},
		{"two items at 1", toSlice(t1, t2), 1},
		{"three items at 0", toSlice(t1, t2, t3), 0},
		{"three items at 1", toSlice(t1, t2, t3), 1},
		{"three items at 2", toSlice(t1, t2, t3), 2},
	} {
		t.Run(tc.label, func(t *testing.T) {
			s := DelTypeSpecAt(tc.s, tc.i)
			if err := checkDel(s, tc.s); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestFuncDeclAt(t *testing.T) {
	toSlice := func(v ...*ast.FuncDecl) []*ast.FuncDecl { return v }

	for _, tc := range []struct {
		label string
		s     []*ast.FuncDecl
		i     int
	}{
		{"one item", toSlice(f1), 0},
		{"two items at 0", toSlice(f1, f2), 0},
		{"two items at 1", toSlice(f1, f2), 1},
		{"three items at 0", toSlice(f1, f2, f3), 0},
		{"three items at 1", toSlice(f1, f2, f3), 1},
		{"three items at 2", toSlice(f1, f2, f3), 2},
	} {
		t.Run(tc.label, func(t *testing.T) {
			s := DelFuncDeclAt(tc.s, tc.i)
			if err := checkDel(s, tc.s); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestValueSpecAt(t *testing.T) {
	for _, tc := range []struct {
		label string
		s     *ast.ValueSpec
		i     int
	}{
		{"one item", dupValueSpec(v1), 0},
		{"two items at 0", dupValueSpec(v2), 0},
		{"two items at 1", dupValueSpec(v2), 1},
		{"three items at 0", dupValueSpec(v3), 0},
		{"three items at 1", dupValueSpec(v3), 1},
		{"three items at 2", dupValueSpec(v3), 2},
	} {
		t.Run(tc.label, func(t *testing.T) {
			names := tc.s.Names
			values := tc.s.Values
			DelValueSpecAt(tc.s, tc.i)
			if err := checkDel(tc.s.Names, names); err != nil {
				t.Fatal(err)
			}
			if err := checkDel(tc.s.Values, values); err != nil {
				t.Fatal(err)
			}
		})
	}
}
