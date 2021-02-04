package main

import (
	"fmt"
	"go/ast"
)

type TypeExplorer struct {
	TypeName    string
	Methods     []string
	IsInterface bool
}

func NewTypeExplorer(pkg *ast.Package, typeName string) *TypeExplorer {
	explorer := &TypeExplorer{
		TypeName: typeName,
	}

	ast.Walk(explorer, pkg)

	return explorer
}

func (explorer *TypeExplorer) Visit(node ast.Node) ast.Visitor {
	if n, ok := node.(*ast.FuncDecl); ok && n.Recv != nil {
		receiver := getIdentName(n.Recv.List[0].Type)
		if receiver == explorer.TypeName {
			method := fmt.Sprintf("%s(%v)", getIdentName(n.Name), getIdentName(n.Recv.List[0].Type))
			explorer.Methods = append(explorer.Methods, method)
		}
	}

	return explorer
}

func (explorer *TypeExplorer) HasEquals() bool {
	return explorer.HasMethod(fmt.Sprintf("Equals(%s)", explorer.TypeName))
}

func (explorer *TypeExplorer) HasString() bool {
	return explorer.HasMethod("String()")
}

func (explorer *TypeExplorer) HasMethod(lookingFor string) bool {
	for _, method := range explorer.Methods {
		if method == lookingFor {
			return true
		}
	}

	return false
}
