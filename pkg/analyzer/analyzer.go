package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/ast/inspector"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name:     "lenchecklint",
	Doc:      "Checks if the length of text-based fields are using the built-in len() function",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

var defaultBlacklistedArgNames = map[string]bool{
	"text":    true,
	"message": true,
	"msg":     true,
	"body":    true,
}

func run(pass *analysis.Pass) (interface{}, error) {
	astInspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	checkArgName := func(ident *ast.Ident) {
		if _, ok := defaultBlacklistedArgNames[strings.ToLower(ident.Name)]; ok {
			pass.Reportf(
				ident.Pos(),
				"length check of argument %s using builtin len function, consider using utf8.RuneCountInString instead",
				ident.Name,
			)
		}
	}

	astInspector.Preorder(nodeFilter, func(node ast.Node) {
		funcCall := node.(*ast.CallExpr)

		funcIdent, ok := funcCall.Fun.(*ast.Ident)
		if !ok {
			return
		}

		if funcIdent.Name != "len" {
			return
		}

		switch funcCall.Args[0].(type) {
		case *ast.Ident:
			argIdent, ok := funcCall.Args[0].(*ast.Ident)
			if ok {
				checkArgName(argIdent)
			}
		case *ast.SelectorExpr:
			selExpr, ok := funcCall.Args[0].(*ast.SelectorExpr)
			if ok {
				checkArgName(selExpr.Sel)
			}
		}
	})

	return nil, nil
}
