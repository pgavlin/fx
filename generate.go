//go:build ignore

package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/types"
	"iter"
	"log"
	"maps"
	"os"
	"slices"
	"strings"
	"text/template"

	"github.com/pgavlin/fx/v2"
	"golang.org/x/tools/go/packages"
)

type list[E any] interface {
	At(i int) E
	Len() int
}

func iterList[L list[E], E any](l L) iter.Seq[E] {
	return func(yield func(E) bool) {
		for i := 0; i < l.Len(); i++ {
			if !yield(l.At(i)) {
				return
			}
		}
	}
}

func indexList[L list[E], E comparable](l L, e E) int {
	for i := 0; i < l.Len(); i++ {
		if l.At(i) == e {
			return i
		}
	}
	return -1
}

var slicesTemplate = template.Must(template.New("adapters.go").Parse(`package slices

import (
	"iter"
	"slices"

	"github.com/pgavlin/fx/v2"
)

{{ range .Functions }}
{{ .Doc }}
func {{ .Name }}[{{ .TypeParams }}](s S{{ .Params }}){{ .Results }} {
	{{ .Return }}fx.{{ .Callee }}(slices.Values(s){{.Args}})
}
{{ end }}
`))

var mapsTemplate = template.Must(template.New("adapters.go").Parse(`package maps

import (
	"iter"
	"maps"

	"github.com/pgavlin/fx/v2"
)

{{ range .Functions }}
{{ .Doc }}
func {{ .Name }}[{{ .TypeParams }}](m M{{ .Params }}){{ .Results }} {
	{{ .Return }}fx.{{ .Callee }}(maps.All(m){{.Args}})
}
{{ end }}
`))

var setsTemplate = template.Must(template.New("adapters.go").Parse(`package sets

import (
	"iter"
	"maps"

	"github.com/pgavlin/fx/v2"
)

{{ range .Functions }}
{{ .Doc }}
func {{ .Name }}[{{ .TypeParams }}](s S{{ .Params }}){{ .Results }} {
	{{ .Return }}fx.{{ .Callee }}(maps.Keys(s){{.Args}})
}
{{ end }}
`))

func prependComma(s string) string {
	if s == "" {
		return ""
	}
	return ", " + s
}

func parenthesize(s []string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return " " + s[0]
	default:
		return "(" + strings.Join(s, ", ") + ")"
	}
}

type functionData struct {
	Doc        string
	Name       string
	Callee     string
	TypeParams string
	Params     string
	Results    string
	Return     string
	Args       string
}

type collector struct {
	seq  types.Type
	seq2 types.Type
	docs map[types.Object]string
}

func (c *collector) matchFirstParam(fn *types.Func, want types.Type) (*types.Named, bool) {
	sig := fn.Signature()

	firstParam, ok := fx.First(iterList(sig.Params()))
	if !ok {
		return nil, false
	}
	firstParamNamed, ok := firstParam.Type().(*types.Named)
	if !ok {
		return nil, false
	}
	if !types.Identical(firstParamNamed.Origin(), want) {
		return nil, false
	}

	return firstParamNamed, true
}

func (c *collector) finishFunctionData(fn *types.Func, name string, newTypeParams []string) functionData {
	sig := fn.Signature()

	newParams := slices.Collect(fx.Skip(iterList(sig.Params()), 1))

	newParamStrings := slices.Collect(fx.Map(slices.Values(newParams), func(v *types.Var) string { return v.Name() + " " + v.Type().String() }))
	newResults := slices.Collect(fx.Map(iterList(sig.Results()), func(v *types.Var) string { return v.Type().String() }))

	returns := ""
	if len(newResults) != 0 {
		returns = "return "
	}

	newArgs := slices.Collect(fx.Map(slices.Values(newParams), func(p *types.Var) string { return p.Name() }))

	return functionData{
		Doc:        strings.Replace(c.docs[fn], fn.Name(), name, 1),
		Name:       name,
		Callee:     fn.Name(),
		TypeParams: strings.Join(newTypeParams, ", "),
		Params:     prependComma(strings.Join(newParamStrings, ", ")),
		Results:    parenthesize(newResults),
		Return:     returns,
		Args:       prependComma(strings.Join(newArgs, ", ")),
	}
}

func (c *collector) needsSetOrSliceMethod(fn *types.Func) (typeParamIndex int, ok bool) {
	// All slice-adapted methods are of the form `func Name[..., T any, ...](it iter.Seq[T], ...) ...`
	// The projected form is `func Name[..., S ~[]T, T any, ...](s S, ...) ...`

	firstParamNamed, ok := c.matchFirstParam(fn, c.seq)
	if !ok {
		return 0, false
	}
	sig := fn.Signature()

	// The type arg for the Seq param should be a type parameter.
	typeParamArg, ok := firstParamNamed.TypeArgs().At(0).(*types.TypeParam)
	if !ok {
		return 0, false
	}
	typeParamIndex = indexList(sig.TypeParams(), typeParamArg)
	return typeParamIndex, typeParamIndex != -1
}

func (c *collector) needsSliceMethod(fn *types.Func) (functionData, bool) {
	typeParamIndex, ok := c.needsSetOrSliceMethod(fn)
	if !ok {
		return functionData{}, false
	}
	sig := fn.Signature()

	// Insert a new type parameter for the slice type preceding the original type parameter.
	newTypeParams := make([]string, 0, sig.TypeParams().Len())
	for i := 0; i < sig.TypeParams().Len(); i++ {
		p := sig.TypeParams().At(i)

		if i == typeParamIndex {
			newTypeParams = append(newTypeParams, "S ~[]"+p.Obj().Name())
		}

		newTypeParams = append(newTypeParams, p.Obj().Name()+" "+p.Constraint().String())
	}

	return c.finishFunctionData(fn, fn.Name(), newTypeParams), true
}

func (c *collector) needsSetMethod(fn *types.Func) (functionData, bool) {
	typeParamIndex, ok := c.needsSetOrSliceMethod(fn)
	if !ok {
		return functionData{}, false
	}
	sig := fn.Signature()

	// Insert a new type parameter for the slice type preceding the original type parameter.
	newTypeParams := make([]string, 0, sig.TypeParams().Len())
	for i := 0; i < sig.TypeParams().Len(); i++ {
		p := sig.TypeParams().At(i)

		if i == typeParamIndex {
			newTypeParams = append(newTypeParams, fmt.Sprintf("S ~map[%v]struct{}", p.Obj().Name()))
			newTypeParams = append(newTypeParams, p.Obj().Name()+" comparable")
		} else {
			newTypeParams = append(newTypeParams, p.Obj().Name()+" "+p.Constraint().String())
		}
	}

	return c.finishFunctionData(fn, fn.Name(), newTypeParams), true
}

func (c *collector) needsMapMethod(fn *types.Func) (functionData, bool) {
	// All map-adapted methods are of the form `func Name[..., T any, U any, ...](it iter.Seq2[T, U], ...) ...`
	// The projected form is `func Name[..., M ~map[T]U, T comparable, U any, ...](m M, ...) ...`

	firstParamNamed, ok := c.matchFirstParam(fn, c.seq2)
	if !ok {
		return functionData{}, false
	}
	sig := fn.Signature()

	// The type args for the Seq2 param should be type parameters.
	keyTypeParamArg, ok := firstParamNamed.TypeArgs().At(0).(*types.TypeParam)
	if !ok {
		return functionData{}, false
	}
	keyTypeParamIndex := indexList(sig.TypeParams(), keyTypeParamArg)
	if keyTypeParamIndex == -1 {
		return functionData{}, false
	}

	valueTypeParamArg, ok := firstParamNamed.TypeArgs().At(1).(*types.TypeParam)
	if !ok {
		return functionData{}, false
	}
	valueTypeParamIndex := indexList(sig.TypeParams(), valueTypeParamArg)
	if valueTypeParamIndex == -1 {
		return functionData{}, false
	}

	// Insert a new type parameter for the map type preceding the key type parameter and replace the constraint on the
	// key type parameter with comparable.
	newTypeParams := make([]string, 0, sig.TypeParams().Len())
	for i := 0; i < sig.TypeParams().Len(); i++ {
		p := sig.TypeParams().At(i)
		if i == keyTypeParamIndex {
			newTypeParams = append(newTypeParams, fmt.Sprintf("M ~map[%v]%v", keyTypeParamArg.Obj().Name(), valueTypeParamArg.Obj().Name()))
			newTypeParams = append(newTypeParams, p.Obj().Name()+" comparable")
		} else {
			newTypeParams = append(newTypeParams, p.Obj().Name()+" "+p.Constraint().String())
		}
	}

	name := strings.Replace(fn.Name(), "2", "", 1)
	return c.finishFunctionData(fn, name, newTypeParams), true
}

func main() {
	// Load the root fx package
	pkgs, err := packages.Load(&packages.Config{Mode: packages.LoadSyntax}, ".")
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("expected a single package, not %v", len(pkgs))
	}

	pkg := pkgs[0]

	// Build a mapping from object -> docstring
	allNodes := fx.ConcatMany(fx.Map(slices.Values(pkg.Syntax), func(f *ast.File) iter.Seq[ast.Node] { return ast.Preorder(f) }))
	allFuncDecls := fx.OfType[*ast.FuncDecl](allNodes)
	docs := maps.Collect(fx.MapUnpack(allFuncDecls, func(n *ast.FuncDecl) (types.Object, string) {
		return pkg.TypesInfo.Defs[n.Name], "// " + strings.ReplaceAll(strings.TrimSpace(n.Doc.Text()), "\n", "\n// ")
	}))

	c := collector{
		seq:  pkg.Imports["iter"].Types.Scope().Lookup("Seq").(*types.TypeName).Type(),
		seq2: pkg.Imports["iter"].Types.Scope().Lookup("Seq2").(*types.TypeName).Type(),
		docs: docs,
	}

	// Find all of the functions in the package
	objects := fx.Map(slices.Values(pkg.Types.Scope().Names()), func(n string) types.Object { return pkg.Types.Scope().Lookup(n) })
	functions := fx.OfType[*types.Func](objects)

	// Collect functions to adapt
	var slicesFunctions []functionData
	var mapsFunctions []functionData
	var setsFunctions []functionData
	for fn := range functions {
		if data, ok := c.needsSliceMethod(fn); ok {
			slicesFunctions = append(slicesFunctions, data)
		}
		if data, ok := c.needsMapMethod(fn); ok {
			mapsFunctions = append(mapsFunctions, data)
		}
		if data, ok := c.needsSetMethod(fn); ok {
			setsFunctions = append(setsFunctions, data)
		}
	}

	// Write adapters
	writeTemplate := func(path string, tmpl *template.Template, data []functionData) error {
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()

		return tmpl.Execute(f, map[string]any{"Functions": data})
	}

	errSlices := writeTemplate("slices/adapters.go", slicesTemplate, slicesFunctions)
	errMaps := writeTemplate("maps/adapters.go", mapsTemplate, mapsFunctions)
	errSets := writeTemplate("sets/adapters.go", setsTemplate, setsFunctions)
	if err := errors.Join(errSlices, errMaps, errSets); err != nil {
		log.Fatal(err)
	}
}
