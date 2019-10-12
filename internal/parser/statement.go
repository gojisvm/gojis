package parser

import (
	"github.com/gojisvm/gojis/internal/parser/ast"
	"github.com/gojisvm/gojis/internal/parser/token"
)

func parseStatementList(i *isolate, p param) *ast.StatementList {
	var items []*ast.StatementListItem
	for {
		item := parseStatementListItem(i, p)
		if item == nil {
			break
		}
		items = append(items, item)
	}
	return &ast.StatementList{
		StatementListItems: items,
	}
}

func parseStatementListItem(i *isolate, p param) *ast.StatementListItem {
	return &ast.StatementListItem{
		Statement:   parseStatement(i, p.only(pYield|pAwait|pReturn)),
		Declaration: parseDeclaration(i, p.only(pYield|pAwait)),
	}
}

func parseStatement(i *isolate, p param) *ast.Statement {
	stmt := &ast.Statement{
		BlockStatement:      parseBlockStatement(i, p.only(pYield|pAwait|pReturn)),
		VariableStatement:   parseVariableStatement(i, p.only(pYield|pAwait)),
		EmptyStatement:      parseEmptyStatement(i, 0),
		ExpressionStatement: parseExpressionStatement(i, p.only(pYield|pAwait)),
		IfStatement:         parseIfStatement(i, p.only(pYield|pAwait|pReturn)),
		BreakableStatement:  parseBreakableStatement(i, p.only(pYield|pAwait|pReturn)),
		ContinueStatement:   parseContinueStatement(i, p.only(pYield|pAwait)),
		BreakStatement:      parseBreakStatement(i, p.only(pYield|pAwait)),
		WithStatement:       parseWithStatement(i, p.only(pYield|pAwait|pReturn)),
		LabelledStatement:   parseLabelledStatement(i, p.only(pYield|pAwait|pReturn)),
		ThrowStatement:      parseThrowStatement(i, p.only(pYield|pAwait)),
		TryStatement:        parseTryStatement(i, p.only(pYield|pAwait|pReturn)),
		DebuggerStatement:   parseDebuggerStatement(i, 0),
	}
	if p.is(pReturn) {
		stmt.ReturnStatement = parseReturnStatement(i, p.only(pYield|pAwait))
	}
	return stmt
}

func parseBlockStatement(i *isolate, p param) *ast.BlockStatement {
	return &ast.BlockStatement{
		Block: parseBlock(i, p.only(pYield|pAwait|pReturn)),
	}
}

func parseVariableStatement(i *isolate, p param) *ast.VariableStatement {
	chck := i.checkpoint()

	if !i.acceptTypes(token.Var) {
		i.restore(chck)
		return nil
	}

	list := parseVariableDeclarationList(i, p.only(pYield|pAwait).add(pIn))
	if list == nil {
		i.restore(chck)
		return nil
	}

	if !i.acceptTypes(token.SemiColon) {
		i.restore(chck)
		return nil
	}

	return &ast.VariableStatement{
		VariableDeclarationList: list,
	}
}

func parseEmptyStatement(i *isolate, p param) *ast.EmptyStatement {
	if _, ok := i.accept(token.SemiColon); !ok {
		return nil
	}
	return &ast.EmptyStatement{}
}

func parseExpressionStatement(i *isolate, p param) *ast.ExpressionStatement {

}

func parseIfStatement(i *isolate, p param) *ast.IfStatement {

}

func parseBreakableStatement(i *isolate, p param) *ast.BreakableStatement {

}

func parseContinueStatement(i *isolate, p param) *ast.ContinueStatement {
	chck := i.checkpoint()

	if !i.acceptTypes(token.Continue) {
		i.restore(chck)
		return nil
	}

	var labelIdent *ast.LabelIdentifier

	i.drain(token.Whitespace)
	if i.acceptTypes(token.LineTerminator) {
		i.drain(token.Whitespace, token.LineTerminator)
	} else {
		labelIdent = parseLabelIdentifier(i, p.only(pYield|pAwait))
	}

	if !i.acceptTypes(token.SemiColon) {
		i.restore(chck)
		return nil
	}

	return &ast.ContinueStatement{
		LabelIdentifier: labelIdent,
	}
}

func parseBreakStatement(i *isolate, p param) *ast.BreakStatement {
	chck := i.checkpoint()

	if !i.acceptTypes(token.Break) {
		i.restore(chck)
		return nil
	}

	var labelIdent *ast.LabelIdentifier

	i.drain(token.Whitespace)
	if i.acceptTypes(token.LineTerminator) {
		i.drain(token.Whitespace, token.LineTerminator)
	} else {
		labelIdent = parseLabelIdentifier(i, p.only(pYield|pAwait))
	}

	if !i.acceptTypes(token.SemiColon) {
		i.restore(chck)
		return nil
	}

	return &ast.BreakStatement{
		LabelIdentifier: labelIdent,
	}
}

func parseReturnStatement(i *isolate, p param) *ast.ReturnStatement {
	chck := i.checkpoint()

	if !i.acceptTypes(token.Return) {
		i.restore(chck)
		return nil
	}

	var expr *ast.Expression

	i.drain(token.Whitespace)
	if i.acceptTypes(token.LineTerminator) {
		i.drain(token.Whitespace, token.LineTerminator)
	} else {
		expr = parseExpression(i, p.only(pYield|pAwait).add(pIn))
	}

	if !i.acceptTypes(token.SemiColon) {
		i.restore(chck)
		return nil
	}

	return &ast.ReturnStatement{
		Expression: expr,
	}
}

func parseWithStatement(i *isolate, p param) *ast.WithStatement {

}

func parseLabelledStatement(i *isolate, p param) *ast.LabelledStatement {

}

func parseThrowStatement(i *isolate, p param) *ast.ThrowStatement {

}

func parseTryStatement(i *isolate, p param) *ast.TryStatement {

}

func parseDebuggerStatement(i *isolate, p param) *ast.DebuggerStatement {

}
