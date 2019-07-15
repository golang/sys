// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// merge processes the generated Go files to factorize
// constants, types and functions definitions.
// For all constants, types, and functions that are defined
// precisely identically for each GOARCH, move them into
// a single unified file named after the source file and GOARCH
// (e.g. zerrors_linux.go).
//
//TODO merge is run after ???; see README.md.
package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	// Load the generated source code (file names start with 'z').
	filter := func(fi os.FileInfo) bool {
		name := fi.Name()
		// Skip files not of the form: z<name>_<goos>_<goarch>.go
		return strings.HasPrefix(name, "z") && strings.Count(name, "_") == 2
	}
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ".", filter, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("invalid number of packages: got %d; want 1", len(pkgs))
	}
	var pkg *ast.Package
	for _, p := range pkgs {
		pkg = p
	}

	// Process the package.
	reg := newRegistry(pkg)
	reg.build()

	// Print out the new files and updated source code.
	if err := reg.print(pkg, fset); err != nil {
		log.Fatal(err)
	}
}

type (
	// registry keeps the files merge is interested in grouped by their name:
	// - <name>_<goos>_<goarch>.go are the interesting files
	// - <name>_<goos>.go is the name of the file that will contain the merged objects and points at a *gofile
	registry map[string]*gofile
	// kinds holds the objects (const, type, func) that will be factored out.
	// Assumption: all files have the objects defined in the same order and with the same layout.
	kinds struct {
		consts *ast.ValueSpec  // Constants
		types  []*ast.TypeSpec // Types
		funcs  []*ast.FuncDecl // Functions
	}
	// goarch represents an arch file.
	goarch struct {
		*ast.File
		kinds // Local objects for an arch
	}
	// gofile holds all the arch dependent files for a given interesting file.
	gofile struct {
		arch  map[string]*goarch
		kinds // Factorized objects for all arch
	}
)

//-----------------------------------------------------------------------------
// registry methods.

func newRegistry(pkg *ast.Package) registry {
	reg := make(map[string]*gofile)

	// Group files by name_os and arch.
	for fname, file := range pkg.Files {
		// Group by name_os and arch.
		i := strings.LastIndexByte(fname, '_')
		name := fname[:i]
		arch := fname[i+1:]

		if reg[name] == nil {
			reg[name] = &gofile{arch: make(map[string]*goarch)}
		}
		reg[name].arch[arch] = &goarch{File: file}
	}

	return registry(reg)
}

// String lists the merged files and their architectures. Used for debugging.
func (r registry) String() string {
	var b strings.Builder
	var fnames []string
	for file, gf := range r {
		b.WriteString(file)
		b.WriteByte('\n')
		for k := range gf.arch {
			fnames = append(fnames, k)
		}
		sort.Strings(fnames)
		for _, fname := range fnames {
			b.WriteString(fmt.Sprintf("  %s\n", fname))
		}
		fnames = fnames[:0]
	}
	return b.String()
}

// build populates the kinds of every file and arch.
func (r registry) build() {
	r.build_kinds()

	// Build the intersection of all objects for all arch.
	for _, gf := range r {
		k := &gf.kinds
		for _, ga := range gf.arch {
			k.inter(&ga.kinds)
		}
	}
	// Remove factorized objects for all arch.
	for _, gf := range r {
		k := &gf.kinds
		for _, ga := range gf.arch {
			ga.kinds.diff(k)
			// Update the input file ast.
			trimFile(ga.File, &ga.kinds)
		}
	}
}

// build_kinds populates the kinds of every arch: constants, types and functions.
func (r registry) build_kinds() {
	for _, gf := range r {
		for _, ga := range gf.arch {
			k := &ga.kinds
			for _, decl := range ga.File.Decls {
				switch d := decl.(type) {
				case *ast.GenDecl:
					switch d.Tok {
					case token.CONST:
						k.pushConst(d)

					case token.TYPE:
						k.pushType(d)
					}

				case *ast.FuncDecl:
					k.pushFunc(d)
				}
			}
		}
	}
}

// bufferedFile provides a buffered os.File.
func newbufferedFile(name string) (io.WriteCloser, error) {
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewWriter(f)
	return &bufferedFile{f, buf}, nil
}

type bufferedFile struct {
	f *os.File
	*bufio.Writer
}

func (bf *bufferedFile) Close() error {
	if err := bf.Writer.Flush(); err != nil {
		return err
	}
	return bf.f.Close()
}

func (r registry) print(pkg *ast.Package, fset *token.FileSet) error {
	handleClose := func(errp *error, c io.Closer) {
		if err := c.Close(); err != nil && *errp == nil {
			*errp = err
		}
	}
	doInput := func(name string, file *ast.File) (err error) {
		f, err := newbufferedFile(name)
		if err != nil {
			return err
		}
		defer handleClose(&err, f)
		return printer.Fprint(f, fset, file)
	}
	do := func(name string, gf *gofile) (err error) {
		f, err := newbufferedFile(name + ".go")
		if err != nil {
			return err
		}
		defer handleClose(&err, f)
		// Print header.
		if _, err := fmt.Fprintf(f, "package %s\n\n", pkg.Name); err != nil {
			return err
		}
		// Print factorized objects.
		if err := gf.kinds.print(f); err != nil {
			return err
		}
		// Update input files.
		for arch, ga := range gf.arch {
			name := fmt.Sprintf("%s_%s", name, arch)
			if err := doInput(name, ga.File); err != nil {
				return err
			}
		}
		return nil
	}
	for name, gf := range r {
		if err := do(name, gf); err != nil {
			return err
		}
	}
	return nil
}

//-----------------------------------------------------------------------------
// kinds methods.

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
	interValueSpec(k.consts, kk.consts)
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
	k.types = interTypeSpec(k.types, kk.types)
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
	k.funcs = interFuncDecl(k.funcs, kk.funcs)
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
kloop:
	for i := 0; i < len(k.consts.Names); {
		id := k.consts.Names[i]
		for ki, kid := range kk.consts.Names {
			if idEqual(id, kid) && exprEqual(k.consts.Values[i], kk.consts.Values[ki]) {
				// Constant is in k and kk: factorized, remove from k.
				delValueSpecAt(k.consts, i)
				continue kloop
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
kloop:
	for i := 0; i < len(k.types); {
		t := k.types[i]
		for _, kt := range kk.types {
			if specEqual(t, kt) {
				// Type is in k and kk: factorized, remove from k.
				k.types = delTypeSpecAt(k.types, i)
				continue kloop
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
kloop:
	for i := 0; i < len(k.funcs); {
		f := k.funcs[i]
		for _, kf := range kk.funcs {
			if declEqual(f, kf) {
				// Function is in k and kk: factorized, remove from k.
				k.funcs = delFuncDeclAt(k.funcs, i)
				continue kloop
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

// trimFile removes objects from f that are not in k.
func trimFile(f *ast.File, k *kinds) {
	for i := 0; i < len(f.Decls); {
		switch d := f.Decls[i].(type) {
		case *ast.GenDecl:
			switch d.Tok {
			case token.CONST:
				for i := 0; i < len(d.Specs); {
					v := d.Specs[i].(*ast.ValueSpec)
					interValueSpec(v, k.consts)
					if len(v.Names) > 0 {
						i++
						continue
					}
					// Remove the spec as it has become empty.
					d.Specs = delSpecAt(d.Specs, i)
				}
				if len(d.Specs) == 0 {
					// Remove the decl as it has become empty.
					f.Decls = delDeclAt(f.Decls, i)
					continue
				}

			case token.TYPE:
			topLoop:
				for i := 0; i < len(d.Specs); {
					t := d.Specs[i].(*ast.TypeSpec)
					for _, kt := range k.types {
						if specEqual(t, kt) {
							i++
							continue topLoop
						}
					}
					// Factorized spec, remove it.
					d.Specs = delSpecAt(d.Specs, i)
				}
				if len(d.Specs) == 0 {
					// Remove the decl as it has become empty.
					f.Decls = delDeclAt(f.Decls, i)
					continue
				}
			}

		case *ast.FuncDecl:
			//TODO func support
		}
		i++
	}
}

//-----------------------------------------------------------------------------
// go/ast related helpers.
//
// For all the *Equal functions below: positions and comments are ignored
// but pointers are followed.

func exprMultiEqual(a, b []ast.Expr) bool {
	if len(a) != len(b) {
		return false
	}
	for i, s := range a {
		if !exprEqual(s, b[i]) {
			return false
		}
	}
	return true
}

// exprEqual returns whether or not the expressions are deeply equal.
func exprEqual(a, b ast.Expr) bool {
	switch a := a.(type) {
	case nil:
		return b == nil

	case *ast.ArrayType:
		b, ok := b.(*ast.ArrayType)
		if !ok {
			return false
		}
		return exprEqual(a.Len, b.Len) && exprEqual(a.Elt, b.Elt)

	case *ast.BasicLit:
		b, ok := b.(*ast.BasicLit)
		if !ok {
			return false
		}
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil {
			return false
		}
		return a.Kind == b.Kind && a.Value == b.Value

	case *ast.BinaryExpr:
		b, ok := b.(*ast.BinaryExpr)
		if !ok {
			return false
		}
		return exprEqual(a.X, b.X) && exprEqual(a.Y, b.Y)

	case *ast.CallExpr:
		b, ok := b.(*ast.CallExpr)
		if !ok {
			return false
		}
		if !exprEqual(a.Fun, b.Fun) {
			return false
		}
		return exprMultiEqual(a.Args, b.Args)

	case *ast.ChanType:
		b, ok := b.(*ast.ChanType)
		if !ok {
			return false
		}
		return a.Dir == b.Dir && exprEqual(a.Value, b.Value)

	case *ast.CompositeLit:
		b, ok := b.(*ast.CompositeLit)
		if !ok {
			return false
		}
		if !exprEqual(a.Type, b.Type) {
			return false
		}
		return exprMultiEqual(a.Elts, b.Elts)

	case *ast.Ellipsis:
		b, ok := b.(*ast.Ellipsis)
		if !ok {
			return false
		}
		return exprEqual(a.Elt, b.Elt)

	case *ast.FuncLit:
		b, ok := b.(*ast.FuncLit)
		if !ok {
			return false
		}
		if !fieldListEqual(a.Type.Params, b.Type.Params) || !fieldListEqual(a.Type.Results, b.Type.Results) {
			return false
		}
		return stmtMultiEqual(a.Body.List, b.Body.List)

	case *ast.Ident:
		b, ok := b.(*ast.Ident)
		if !ok {
			return false
		}
		return idEqual(a, b)

	case *ast.IndexExpr:
		b, ok := b.(*ast.IndexExpr)
		if !ok {
			return false
		}
		return exprEqual(a.X, b.X) && exprEqual(a.Index, b.Index)

	case *ast.InterfaceType:
		b, ok := b.(*ast.InterfaceType)
		if !ok {
			return false
		}
		return fieldListEqual(a.Methods, b.Methods)

	case *ast.KeyValueExpr:
		b, ok := b.(*ast.KeyValueExpr)
		if !ok {
			return false
		}
		return exprEqual(a.Key, b.Key) && exprEqual(a.Value, b.Value)

	case *ast.MapType:
		b, ok := b.(*ast.MapType)
		if !ok {
			return false
		}
		return exprEqual(a.Key, b.Key) && exprEqual(a.Value, b.Value)

	case *ast.SelectorExpr:
		b, ok := b.(*ast.SelectorExpr)
		if !ok {
			return false
		}
		return exprEqual(a.Sel, b.Sel) && exprEqual(a.X, b.X)

	case *ast.SliceExpr:
		b, ok := b.(*ast.SliceExpr)
		if !ok {
			return false
		}
		return a.Slice3 == b.Slice3 &&
			exprEqual(a.X, b.X) && exprEqual(a.Max, b.Max) &&
			exprEqual(a.High, b.High) && exprEqual(a.Low, b.Low)

	case *ast.StarExpr:
		b, ok := b.(*ast.StarExpr)
		if !ok {
			return false
		}
		return exprEqual(a.X, b.X)

	case *ast.StructType:
		b, ok := b.(*ast.StructType)
		if !ok {
			return false
		}
		return fieldListEqual(a.Fields, b.Fields)

	case *ast.UnaryExpr:
		b, ok := b.(*ast.UnaryExpr)
		if !ok {
			return false
		}
		return exprEqual(a.X, b.X)
	}
	panic(fmt.Sprintf("unsupported expr %T", a))
}

// fieldListEqual returns whether or not a and b are deeply equal.
func fieldListEqual(a, b *ast.FieldList) bool {
	if len(a.List) != len(b.List) {
		return false
	}
	for i, fa := range a.List {
		fb := b.List[i]
		if !exprEqual(fa.Type, fb.Type) {
			return false
		}
		if !exprEqual(fa.Tag, fb.Tag) {
			return false
		}
		if len(fa.Names) != len(fb.Names) {
			return false
		}
		for i, id := range fa.Names {
			if !exprEqual(id, fb.Names[i]) {
				return false
			}
		}
	}
	return true
}

func stmtMultiEqual(a, b []ast.Stmt) bool {
	if len(a) != len(b) {
		return false
	}
	for i, s := range a {
		if !stmtEqual(s, b[i]) {
			return false
		}
	}
	return true
}

// stmtEqual returns whether or not a and b are deeply equal.
func stmtEqual(a, b ast.Stmt) bool {
	switch a := a.(type) {
	case nil:
		return b == nil

	case *ast.BadStmt:
		_, ok := b.(*ast.BadStmt)
		return ok

	case *ast.AssignStmt:
		b, ok := b.(*ast.AssignStmt)
		if !ok {
			return false
		}
		return exprMultiEqual(a.Lhs, b.Lhs) && exprMultiEqual(a.Rhs, b.Rhs)

	case *ast.BlockStmt:
		b, ok := b.(*ast.BlockStmt)
		if !ok {
			return false
		}
		return stmtMultiEqual(a.List, b.List)

	case *ast.BranchStmt:
		b, ok := b.(*ast.BranchStmt)
		if !ok {
			return false
		}
		return exprEqual(a.Label, b.Label)

	case *ast.DeclStmt:
		b, ok := b.(*ast.DeclStmt)
		if !ok {
			return false
		}
		return declEqual(a.Decl, b.Decl)

	case *ast.DeferStmt:
		b, ok := b.(*ast.DeferStmt)
		if !ok {
			return false
		}
		return exprEqual(a.Call, b.Call)

	case *ast.ExprStmt:
		b, ok := b.(*ast.ExprStmt)
		if !ok {
			return false
		}
		return exprEqual(a.X, b.X)

	case *ast.ForStmt:
		b, ok := b.(*ast.ForStmt)
		if !ok {
			return false
		}
		return stmtEqual(a.Init, b.Init) && exprEqual(a.Cond, b.Cond) &&
			stmtEqual(a.Post, b.Post) && stmtEqual(a.Body, b.Body)

	case *ast.GoStmt:
		b, ok := b.(*ast.GoStmt)
		if !ok {
			return false
		}
		return exprEqual(a.Call, b.Call)

	case *ast.IfStmt:
		b, ok := b.(*ast.IfStmt)
		if !ok {
			return false
		}
		return stmtEqual(a.Init, b.Init) &&
			exprEqual(a.Cond, b.Cond) && stmtEqual(a.Body, b.Body) &&
			stmtEqual(a.Else, b.Else)

	case *ast.IncDecStmt:
		b, ok := b.(*ast.IncDecStmt)
		if !ok {
			return false
		}
		return exprEqual(a.X, b.X)

	case *ast.LabeledStmt:
		b, ok := b.(*ast.LabeledStmt)
		if !ok {
			return false
		}
		return stmtEqual(a.Stmt, b.Stmt)

	case *ast.RangeStmt:
		b, ok := b.(*ast.RangeStmt)
		if !ok {
			return false
		}
		return exprEqual(a.X, b.X) &&
			exprEqual(a.Key, b.Key) && exprEqual(a.Value, b.Value) &&
			stmtEqual(a.Body, b.Body)

	case *ast.ReturnStmt:
		b, ok := b.(*ast.ReturnStmt)
		if !ok {
			return false
		}
		return exprMultiEqual(a.Results, b.Results)

	case *ast.SelectStmt:
		b, ok := b.(*ast.SelectStmt)
		if !ok {
			return false
		}
		return stmtEqual(a.Body, b.Body)

	case *ast.SendStmt:
		b, ok := b.(*ast.SendStmt)
		if !ok {
			return false
		}
		return exprEqual(a.Chan, b.Chan) && exprEqual(a.Value, b.Value)

	case *ast.SwitchStmt:
		b, ok := b.(*ast.SwitchStmt)
		if !ok {
			return false
		}
		return stmtEqual(a.Init, b.Init) && exprEqual(a.Tag, b.Tag)

	case *ast.TypeSwitchStmt:
		b, ok := b.(*ast.TypeSwitchStmt)
		if !ok {
			return false
		}
		return stmtEqual(a.Init, b.Init) && stmtEqual(a.Assign, b.Assign)

	default:
		panic(fmt.Sprintf("unsupported statement %T", a))
	}
	return false
}

// declEqual returns whether or not a and b are deeply equal.
func declEqual(a, b ast.Decl) bool {
	switch a := a.(type) {
	case nil:
		return b == nil

	case *ast.BadDecl:
		_, ok := b.(*ast.BadDecl)
		return ok

	case *ast.FuncDecl:
		b, ok := b.(*ast.FuncDecl)
		if !ok {
			return false
		}
		if a.Body == nil && b.Body == nil {
			return true
		}
		if a.Body == nil || b.Body == nil {
			return false
		}
		return stmtEqual(a.Body, b.Body)

	case *ast.GenDecl:
		b, ok := b.(*ast.GenDecl)
		if !ok {
			return false
		}
		return specMultiEqual(a.Specs, b.Specs)

	default:
		panic(fmt.Sprintf("unsupported declaration %T: %v", a, a))
	}
	return false
}

func idMultiEqual(a, b []*ast.Ident) bool {
	if len(a) != len(b) {
		return false
	}
	for i, as := range a {
		if !exprEqual(as, b[i]) {
			return false
		}
	}
	return true
}

func idEqual(a, b *ast.Ident) bool {
	return a.Name == b.Name
}

func specMultiEqual(a, b []ast.Spec) bool {
	if len(a) != len(b) {
		return false
	}
	for i, as := range a {
		if !specEqual(as, b[i]) {
			return false
		}
	}
	return true
}

// specEqual returns whether or not a and b are deeply equal.
func specEqual(a, b ast.Spec) bool {
	switch a := a.(type) {
	case nil:
		return b == nil

	case *ast.ImportSpec:
		if b, ok := b.(*ast.ImportSpec); ok {
			return exprEqual(a.Name, b.Name) && exprEqual(a.Path, b.Path)
		}

	case *ast.TypeSpec:
		if b, ok := b.(*ast.TypeSpec); ok {
			return idEqual(a.Name, b.Name) && exprEqual(a.Type, b.Type)
		}

	case *ast.ValueSpec:
		if b, ok := b.(*ast.ValueSpec); ok {
			return exprEqual(a.Type, b.Type) && idMultiEqual(a.Names, b.Names) && exprMultiEqual(a.Values, b.Values)
		}

	default:
		panic(fmt.Sprintf("unsupported spec %T: %v", a, a))
	}
	return false
}

// intersection of a and b. a is mutated.
func interValueSpec(a, b *ast.ValueSpec) {
	if a == nil || b == nil {
		return
	}
	if !exprEqual(a.Type, b.Type) {
		return
	}
topLoop:
	for i := 0; i < len(a.Names); {
		id := a.Names[i]
		for ki, kid := range b.Names {
			if idEqual(id, kid) && exprEqual(a.Values[i], b.Values[ki]) {
				// Same constant!
				i++
				continue topLoop
			}
		}
		// Constant is in a but not in b: remove from a.
		delValueSpecAt(a, i)
	}
}

// Returns the intersection of a and b.
func interTypeSpec(a, b []*ast.TypeSpec) []*ast.TypeSpec {
	if len(a) == 0 || len(b) == 0 {
		return a
	}
topLoop:
	for i := 0; i < len(a); {
		s := a[i]
		for _, ks := range b {
			if specEqual(s, ks) {
				// Same type!
				i++
				continue topLoop
			}
		}
		// Type is in a but not in b: remove from a.
		a = delTypeSpecAt(a, i)
	}
	return a
}

// Returns the intersection of a and b.
func interFuncDecl(a, b []*ast.FuncDecl) []*ast.FuncDecl {
	if len(a) == 0 || len(b) == 0 {
		return a
	}
topLoop:
	for i := 0; i < len(a); {
		s := a[i]
		for _, ks := range b {
			if declEqual(s, ks) {
				// Same function!
				i++
				continue topLoop
			}
		}
		// Function is in a but not in b: remove from a.
		a = delFuncDeclAt(a, i)
	}
	return a
}

// delValueSpecAt removes the value spec at index i.
func delValueSpecAt(v *ast.ValueSpec, i int) {
	if i+1 < len(v.Names) {
		copy(v.Names[i:], v.Names[i+1:])
		copy(v.Values[i:], v.Values[i+1:])
	}
	v.Names[len(v.Names)-1] = nil
	v.Values[len(v.Values)-1] = nil
	v.Names = v.Names[:len(v.Names)-1]
	v.Values = v.Values[:len(v.Values)-1]
}

// delTypeSpecAt removes the type spec at index i.
func delTypeSpecAt(s []*ast.TypeSpec, i int) []*ast.TypeSpec {
	if i+1 < len(s) {
		copy(s[i:], s[i+1:])
	}
	s[len(s)-1] = nil
	return s[:len(s)-1]
}

// delFuncDeclAt removes the func decl at index i.
func delFuncDeclAt(s []*ast.FuncDecl, i int) []*ast.FuncDecl {
	if i+1 < len(s) {
		copy(s[i:], s[i+1:])
	}
	s[len(s)-1] = nil
	return s[:len(s)-1]
}

// delDecl removes the decl at index i and returns the modified slice.
func delDeclAt(s []ast.Decl, i int) []ast.Decl {
	if i+1 < len(s) {
		copy(s[i:], s[i+1:])
	}
	s[len(s)-1] = nil
	return s[:len(s)-1]
}

// delSpec removes the spec at index i and returns the modified slice.
func delSpecAt(s []ast.Spec, i int) []ast.Spec {
	if i+1 < len(s) {
		copy(s[i:], s[i+1:])
	}
	s[len(s)-1] = nil
	return s[:len(s)-1]
}
