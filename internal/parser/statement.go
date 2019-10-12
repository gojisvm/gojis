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
	if statement := parseStatement(i, p.only(pYield|pAwait|pReturn)); statement != nil {
		return &ast.StatementListItem{
			Statement: statement,
		}
	}

	if declaration := parseDeclaration(i, p.only(pYield|pAwait)); declaration != nil {
		return &ast.StatementListItem{
			Declaration: declaration,
		}
	}

	return nil
}

func parseStatement(i *isolate, p param) *ast.Statement {
	if blockStatement := parseBlockStatement(i, p.only(pYield|pAwait|pReturn)); blockStatement != nil {
		return &ast.Statement{
			BlockStatement: blockStatement,
		}
	}

	if variableStatement := parseVariableStatement(i, p.only(pYield|pAwait)); variableStatement != nil {
		return &ast.Statement{
			VariableStatement: variableStatement,
		}
	}

	if emptyStatement := parseEmptyStatement(i, 0); emptyStatement != nil {
		return &ast.Statement{
			EmptyStatement: emptyStatement,
		}
	}

	if expressionStatement := parseExpressionStatement(i, p.only(pYield|pAwait)); expressionStatement != nil {
		return &ast.Statement{
			ExpressionStatement: expressionStatement,
		}
	}

	if ifStatement := parseIfStatement(i, p.only(pYield|pAwait|pReturn)); ifStatement != nil {
		return &ast.Statement{
			IfStatement: ifStatement,
		}
	}

	if breakableStatement := parseBreakableStatement(i, p.only(pYield|pAwait|pReturn)); breakableStatement != nil {
		return &ast.Statement{
			BreakableStatement: breakableStatement,
		}
	}

	if continueStatement := parseContinueStatement(i, p.only(pYield|pAwait)); continueStatement != nil {
		return &ast.Statement{
			ContinueStatement: continueStatement,
		}
	}

	if breakStatement := parseBreakStatement(i, p.only(pYield|pAwait)); breakStatement != nil {
		return &ast.Statement{
			BreakStatement: breakStatement,
		}
	}

	if p.is(pReturn) {
		if returnStatement := parseReturnStatement(i, p.only(pYield|pAwait)); returnStatement != nil {
			return &ast.Statement{
				ReturnStatement: returnStatement,
			}
		}
	}

	if withStatement := parseWithStatement(i, p.only(pYield|pAwait|pReturn)); withStatement != nil {
		return &ast.Statement{
			WithStatement: withStatement,
		}
	}

	if labelledStatement := parseLabelledStatement(i, p.only(pYield|pAwait|pReturn)); labelledStatement != nil {
		return &ast.Statement{
			LabelledStatement: labelledStatement,
		}
	}

	if throwStatement := parseThrowStatement(i, p.only(pYield|pAwait)); throwStatement != nil {
		return &ast.Statement{
			ThrowStatement: throwStatement,
		}
	}

	if tryStatement := parseTryStatement(i, p.only(pYield|pAwait|pReturn)); tryStatement != nil {
		return &ast.Statement{
			TryStatement: tryStatement,
		}
	}

	if debuggerStatement := parseDebuggerStatement(i, 0); debuggerStatement != nil {
		return &ast.Statement{
			DebuggerStatement: debuggerStatement,
		}
	}

	return nil
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
	chck := i.checkpoint()

	// neagtive lookaheads
	if i.acceptTypes(token.BraceOpen, token.Function, token.Class) {
		i.restore(chck)
		return nil
	}

	// more negative lookaheads
	if i.acceptTypes(token.Async) {
		i.drain(token.Whitespace)
		if i.acceptTypes(token.Function) {
			i.restore(chck)
			return nil
		}
	}
	if i.acceptTypes(token.Let) {
		i.drain(token.Whitespace, token.LineTerminator)
		if i.acceptTypes(token.BracketOpen) {
			i.restore(chck)
			return nil
		}
	}

	// actual parsing

	// restore checkpoint to make sure lookaheads didn't move position marker
	i.restore(chck)

	expr := parseExpression(i, p.only(pYield|pAwait).add(pIn))

	if !i.acceptTypes(token.SemiColon) {
		i.restore(chck)
		return nil
	}

	return &ast.ExpressionStatement{
		Expression: expr,
	}
}

func parseIfStatement(i *isolate, p param) *ast.IfStatement {
	chck := i.checkpoint()

	if !i.acceptTypes(token.If) {
		i.restore(chck)
		return nil
	}

	i.drain(token.Whitespace, token.LineTerminator)

	if !i.acceptTypes(token.ParOpen) {
		i.restore(chck)
		return nil
	}

	i.drain(token.Whitespace, token.LineTerminator)

	expr := parseExpression(i, p.only(pYield|pAwait).add(pIn))

	i.drain(token.Whitespace, token.LineTerminator)

	if !i.acceptTypes(token.ParClose) {
		i.restore(chck)
		return nil
	}

	i.drain(token.Whitespace, token.LineTerminator)

	stmt := parseStatement(i, p.only(pYield|pAwait|pReturn))

	chckBeforeWhitespace := i.checkpoint()
	i.drain(token.Whitespace, token.LineTerminator)

	// optional else part
	if i.acceptTypes(token.Else) {
		i.drain(token.Whitespace, token.LineTerminator)
		elseStmt := parseStatement(i, p.only(pYield|pAwait|pReturn))
		if elseStmt == nil {
			i.restore(chck)
			return nil
		}
		return &ast.IfStatement{
			Expression:    expr,
			Statement:     stmt,
			ElseStatement: elseStmt,
		}
	}

	i.restore(chckBeforeWhitespace)

	return &ast.IfStatement{
		Expression: expr,
		Statement:  stmt,
	}
}

func parseBreakableStatement(i *isolate, p param) *ast.BreakableStatement {
	if itStmt := parseIterationStatement(i, p.only(pYield|pAwait|pReturn)); itStmt != nil {
		return &ast.BreakableStatement{
			IterationStatement: itStmt,
		}
	}

	if switchStmt := parseSwitchStatement(i, p.only(pYield|pAwait|pReturn)); switchStmt != nil {
		return &ast.BreakableStatement{
			IterationStatement: switchStmt,
		}
	}

	return nil
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
	chck := i.checkpoint()

	if !i.acceptTypes(token.With) {
		i.restore(chck)
		return nil
	}

	i.drain(token.Whitespace, token.LineTerminator)

	if !i.acceptTypes(token.ParOpen) {
		i.restore(chck)
		return nil
	}

	i.drain(token.Whitespace, token.LineTerminator)

	expr := parseExpression(i, p.only(pYield|pAwait).add(pIn))
	if expr == nil {
		i.restore(chck)
		return nil
	}

	i.drain(token.Whitespace, token.LineTerminator)

	if !i.acceptTypes(token.ParClose) {
		i.restore(chck)
		return nil
	}

	i.drain(token.Whitespace, token.LineTerminator)

	stmt := parseStatement(i, p.only(pYield|pAwait|pReturn))
	if stmt == nil {
		i.restore(chck)
		return nil
	}

	return &ast.WithStatement{
		Expression: expr,
		Statement:  stmt,
	}
}

func parseLabelledStatement(i *isolate, p param) *ast.LabelledStatement {
	chck := i.checkpoint()

	labelIdent := parseLabelIdentifier(i, p.only(pYield|pAwait))
	if labelIdent == nil {
		i.restore(chck)
		return nil
	}

	i.drain(token.Whitespace, token.LineTerminator)

	if !i.acceptTypes(token.Colon) {
		i.restore(chck)
		return nil
	}

	i.drain(token.Whitespace, token.LineTerminator)

	labelledItem := parseLabelledItem(i, p.only(pYield|pAwait|pReturn))
	if labelledItem == nil {
		i.restore(chck)
		return nil
	}

	return &ast.LabelledStatement{
		LabelIdentifier: labelIdent,
		LabelledItem:    labelledItem,
	}
}

func parseThrowStatement(i *isolate, p param) *ast.ThrowStatement {
	chck := i.checkpoint()

	if !i.acceptTypes(token.Throw) {
		i.restore(chck)
		return nil
	}

	i.drain(token.Whitespace)

	expr := parseExpression(i, p.only(pYield|pAwait).add(pIn))
	if expr == nil {
		i.restore(chck)
		return nil
	}

	i.drain(token.Whitespace, token.LineTerminator)

	if !i.acceptTypes(token.SemiColon) {
		i.restore(chck)
		return nil
	}

	return &ast.ThrowStatement{
		Expression: expr,
	}
}

func parseTryStatement(i *isolate, p param) *ast.TryStatement {
	chck := i.checkpoint()

	if !i.acceptTypes(token.Try) {
		i.restore(chck)
		return nil
	}

	i.drain(token.Whitespace, token.LineTerminator)

	blk := parseBlock(i, p.only(pYield|pAwait|pReturn))
	if blk == nil {
		i.restore(chck)
		return nil
	}

	i.drain(token.Whitespace, token.LineTerminator)

	catch := parseCatch(i, p.only(pYield|pAwait|pReturn))

	i.drain(token.Whitespace, token.LineTerminator)

	finally := parseFinally(i, p.only(pYield|pAwait|pReturn))
	if finally == nil {
		i.restore(chck)
		return nil
	}

	return &ast.TryStatement{
		Block:   blk,
		Catch:   catch,
		Finally: finally,
	}
}

func parseDebuggerStatement(i *isolate, p param) *ast.DebuggerStatement {
	chck := i.checkpoint()

	if !i.acceptTypes(token.Debugger) {
		i.restore(chck)
		return nil
	}

	i.drain(token.Whitespace, token.LineTerminator)

	if !i.acceptTypes(token.SemiColon) {
		i.restore(chck)
		return nil
	}

	return &ast.DebuggerStatement{}
}
