package parser

import (
	"github.com/gojisvm/gojis/internal/parser/ast"
	"github.com/gojisvm/gojis/internal/parser/token"
)

func parseLabelIdentifier(i *isolate, p param) *ast.LabelIdentifier {
	if !p.is(pYield) && i.acceptType(token.Yield) {
		return &ast.LabelIdentifier{
			Yield: true,
		}
	}
	if !p.is(pAwait) && i.acceptType(token.Await) {
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
