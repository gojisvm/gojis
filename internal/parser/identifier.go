package parser

import (
	"github.com/gojisvm/gojis/internal/parser/ast"
	"github.com/gojisvm/gojis/internal/parser/token"
)

func parseLabelIdentifier(i *isolate, p param) *ast.LabelIdentifier {
	if !p.is(pYield) && i.acceptOneOfTypes(token.Yield) {
		return &ast.LabelIdentifier{
			Yield: true,
		}
	}
	if !p.is(pAwait) && i.acceptOneOfTypes(token.Await) {
		return &ast.LabelIdentifier{
			Await: true,
		}
	}
	return &ast.LabelIdentifier{
		Identifier: parseIdentifier(i, 0),
	}
}

func parseIdentifier(i *isolate, p param) *ast.Identifier {
	// TODO: IdentifierName but not ReservedWord
	if t, ok := i.accept(token.IdentifierName); ok {
		return &ast.Identifier{
			IdentifierName: t.Value,
		}
	}
	return nil
}

func parseBindingIdentifier(i *isolate, p param) *ast.BindingIdentifier {
	if ident := parseIdentifier(i, 0); ident != nil {
		return &ast.BindingIdentifier{
			Identifier: ident,
		}
	}

	if i.acceptOneOfTypes(token.Yield) {
		return &ast.BindingIdentifier{
			Yield: true,
		}
	}

	if i.acceptOneOfTypes(token.Await) {
		return &ast.BindingIdentifier{
			Await: true,
		}
	}

	return nil
}
