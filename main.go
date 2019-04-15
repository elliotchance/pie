//go:generate go run generate_template.go

package main

import (
	"fmt"
	"github.com/elliotchance/pie/functions"
	"github.com/elliotchance/pie/pie"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getIdentName(e ast.Expr) string {
	switch v := e.(type) {
	case *ast.Ident:
		return v.Name

	case *ast.StarExpr:
		return "*" + getIdentName(v.X)

	default:
		panic(fmt.Sprintf("cannot decode %T", e))
	}
}

func findType(pkgs map[string]*ast.Package, name string) (packageName, elementType string) {
	for pkgName, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				if genDecl, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range genDecl.Specs {
						if typeSpec, ok := spec.(*ast.TypeSpec); ok {
							if typeSpec.Name.String() == name {
								if t, ok := typeSpec.Type.(*ast.ArrayType); ok {
									return pkgName, getIdentName(t.Elt)
								} else {
									panic(fmt.Sprintf("type %s must be a slice", name))
								}
							}
						}
					}
				}
			}
		}
	}

	panic(fmt.Sprintf("type %s does not exist", name))
}

func getType(name string) int {
	switch name {
	case "int8", "uint8", "byte", "int16", "uint16", "int32", "rune", "uint32",
		"int64", "uint64", "int", "uint", "uintptr", "float32", "float64",
		"complex64", "complex128":
		return functions.ForNumbers

	case "string":
		return functions.ForStrings
	}

	return functions.ForStructs
}

func getImports(packageName, s string) (imports []string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", s, parser.ImportsOnly)
	if err != nil {
		panic(err)
	}

	for _, s := range f.Imports {
		importName := s.Path.Value

		if importName == `"github.com/elliotchance/pie/pie"` &&
			isSelfPackage(packageName) {
			continue
		}

		imports = append(imports, importName)
	}

	return
}

func getAllImports(packageName string, files []string) (imports []string) {
	mapImports := map[string]struct{}{}

	for _, file := range files {
		for _, imp := range getImports(packageName, file) {
			mapImports[imp] = struct{}{}
		}
	}

	for imp := range mapImports {
		imports = append(imports, imp)
	}

	sort.Strings(imports)

	return
}

// We have to generate imports slightly differently when we are building code
// that will go into its own packages vs an external package.
func isSelfPackage(packageName string) bool {
	return packageName == "pie"
}

func main() {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ".", nil, 0)
	check(err)

	for _, arg := range os.Args[1:] {
		sliceType, fns := getFunctionsFromArg(arg)
		packageName, elementType := findType(pkgs, sliceType)
		kind := getType(elementType)

		templates := []string{}
		for _, function := range functions.Functions {
			if len(fns) > 0 && !pie.Strings(fns).Contains(function.Name) {
				continue
			}

			if function.For&kind != 0 {
				templates = append(templates, pieTemplates[function.Name])
			}
		}

		// Aggregate imports.
		t := fmt.Sprintf("package %s\n\n", packageName)

		imports := getAllImports(packageName, templates)
		if len(imports) > 0 {
			t += fmt.Sprintf("import (")
			for _, imp := range imports {
				t += fmt.Sprintf("\n\t%s", imp)
			}
			t += "\n)\n\n"
		}

		for _, tmpl := range templates {
			i := strings.Index(tmpl, "//")
			t += tmpl[i:] + "\n"
		}

		t = strings.Replace(t, "StringSliceType", sliceType, -1)
		t = strings.Replace(t, "StringElementType", elementType, -1)
		t = strings.Replace(t, "SliceType", sliceType, -1)
		t = strings.Replace(t, "ElementType", elementType, -1)

		switch kind {
		case functions.ForNumbers:
			t = strings.Replace(t, "ElementZeroValue", "0", -1)

		case functions.ForStrings:
			t = strings.Replace(t, "ElementZeroValue", `""`, -1)

		case functions.ForStructs:
			zeroValue := fmt.Sprintf("%s{}", elementType)

			// If its a pointer we need to replace '*' -> '&' when
			// instantiating.
			if elementType[0] == '*' {
				zeroValue = "&" + zeroValue[1:]
			}

			t = strings.Replace(t, "ElementZeroValue", zeroValue, -1)
		}

		if isSelfPackage(packageName) {
			t = strings.Replace(t, "pie.Strings", "Strings", -1)
		}

		// The TrimRight is important to remove an extra new line that conflicts
		// with go fmt.
		t = strings.TrimRight(t, "\n") + "\n"

		// Filter out any functions we dont want.
		//t = filterFunctions(t, functions)

		err := ioutil.WriteFile(strings.ToLower(sliceType)+"_pie.go", []byte(t), 0755)
		check(err)
	}
}

func getFunctionsFromArg(arg string) (string, []string) {
	parts := strings.Split(arg, ".")

	return parts[0], parts[1:]
}
