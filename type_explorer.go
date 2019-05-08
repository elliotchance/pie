package main

import (
	"fmt"
	"go/ast"
)

type TypeExplorer struct {
	TypeName string
	Methods  []string
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
	lookingFor := fmt.Sprintf("Equals(%s)", explorer.TypeName)
	for _, method := range explorer.Methods {
		if method == lookingFor {
			return true
		}
	}

	return false
}
