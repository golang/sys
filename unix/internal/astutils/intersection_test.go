// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package astutils

import (
	"go/ast"
	"reflect"
	"testing"
)

type set map[interface{}]struct{}

func toSet(a interface{}) set {
	s := set{}
	av := reflect.ValueOf(a)
	for i := 0; i < av.Len(); i++ {
		s[av.Index(i).Interface()] = struct{}{}
	}
	return s
}

func (s set) inter(ss set) set {
	for k := range s {
		if _, ok := ss[k]; !ok {
			delete(s, k)
		}
	}
	return s
}

func (s set) equal(ss set) bool {
	if len(s) != len(ss) {
		return false
	}
	for k := range s {
		if _, ok := ss[k]; !ok {
			return false
		}
	}
	return true
}

// checkInter checks that r is the intersection of a and b.
func checkInter(a, b interface{}, r interface{}) bool {
	return toSet(a).inter(toSet(b)).equal(toSet(r))
}

func TestInterFuncDecl(t *testing.T) {
	toSlice := func(v ...*ast.FuncDecl) []*ast.FuncDecl { return v }

	// Special case: empty second slice.
	a := toSlice(f1)
	b := InterFuncDecl(a, nil)
	if !toSet(a).equal(toSet(b)) {
		t.Fatal("invalid intersection for empty second slice")
	}

	for _, tc := range []struct {
		label string
		a, b  []*ast.FuncDecl
	}{
		{"1 inter 1", toSlice(f1), toSlice(f1)},
		{"1 inter 1 2", toSlice(f1), toSlice(f1, f2)},
		{"2 inter 1 2", toSlice(f2), toSlice(f1, f2)},
		{"2 inter 1 2 3", toSlice(f2), toSlice(f1, f2, f3)},
		{"1 2 inter 1 2 3", toSlice(f1, f2), toSlice(f1, f2, f3)},
		{"1 3 inter 1 2 3", toSlice(f1, f3), toSlice(f1, f2, f3)},
		{"2 3 inter 1 2 3", toSlice(f2, f3), toSlice(f1, f2, f3)},
	} {
		t.Run(tc.label, func(t *testing.T) {
			s := InterFuncDecl(tc.a, tc.b)
			if !checkInter(tc.a, tc.b, s) {
				t.Fatal("invalid intersection")
			}
		})
	}
}

func TestInterTypeSpec(t *testing.T) {
	toSlice := func(v ...*ast.TypeSpec) []*ast.TypeSpec { return v }

	// Special case: empty second slice.
	a := toSlice(t1)
	b := InterTypeSpec(a, nil)
	if !toSet(a).equal(toSet(b)) {
		t.Fatal("invalid intersection for empty second slice")
	}

	for _, tc := range []struct {
		label string
		a, b  []*ast.TypeSpec
	}{
		{"1 inter 1", toSlice(t1), toSlice(t1)},
		{"1 inter 1 2", toSlice(t1), toSlice(t1, t2)},
		{"2 inter 1 2", toSlice(t2), toSlice(t1, t2)},
		{"2 inter 1 2 3", toSlice(t2), toSlice(t1, t2, t3)},
		{"1 2 inter 1 2 3", toSlice(t1, t2), toSlice(t1, t2, t3)},
		{"1 3 inter 1 2 3", toSlice(t1, t3), toSlice(t1, t2, t3)},
		{"2 3 inter 1 2 3", toSlice(t2, t3), toSlice(t1, t2, t3)},
	} {
		t.Run(tc.label, func(t *testing.T) {
			s := InterTypeSpec(tc.a, tc.b)
			if !checkInter(tc.a, tc.b, s) {
				t.Fatal("invalid intersection")
			}
		})
	}
}

func TestInterValueSpec(t *testing.T) {
	// Special case: empty second slice.
	n, v := v1.Names, v1.Values
	InterValueSpec(v1, nil)
	if !toSet(v1.Names).equal(toSet(n)) || !toSet(v1.Values).equal(toSet(v)) {
		t.Fatal("invalid intersection for empty second slice")
	}

	for _, tc := range []struct {
		label string
		a, b  *ast.ValueSpec
	}{
		{"1 inter 1", dupValueSpec(v1), dupValueSpec(v1)},
		{"1 inter 2", dupValueSpec(v1), dupValueSpec(v2)},
		{"1 inter 3", dupValueSpec(v1), dupValueSpec(v3)},
		{"2 inter 1", dupValueSpec(v2), dupValueSpec(v1)},
		{"2 inter 2", dupValueSpec(v2), dupValueSpec(v2)},
		{"2 inter 3", dupValueSpec(v2), dupValueSpec(v3)},
		{"3 inter 3", dupValueSpec(v3), dupValueSpec(v3)},
	} {
		t.Run(tc.label, func(t *testing.T) {
			names, values := tc.a.Names, tc.a.Values
			defer func() {
				// Restore slices are they are modified by the tested function.
				tc.a.Names = names
				tc.a.Values = values
			}()
			InterValueSpec(tc.a, tc.b)
			if !checkInter(names, tc.b.Names, tc.a.Names) {
				t.Fatal("invalid intersection")
			}
			if !checkInter(values, tc.b.Values, tc.a.Values) {
				t.Fatal("invalid intersection")
			}
		})
	}
}
