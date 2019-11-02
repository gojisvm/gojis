package parser

import (
	"github.com/gojisvm/gojis/internal/parser/ast"
	"github.com/gojisvm/gojis/internal/parser/token"
)

func parseStatementList(i *isolate, p param) *ast.StatementList {
	chck := i.checkpoint()

	var items []*ast.StatementListItem
	for {

		item := parseStatementListItem(i, p)
		if item == nil {
			if len(items) == 0 {
				// if not even a single one has been parsed
				i.restore(chck)
				return nil
			}
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

func parseBlock(i *isolate, p param) *ast.Block {
	return &ast.Block{
		StatementList: parseStatementList(i, p.only(pYield|pAwait|pReturn)),
	}
}

func parseVariableStatement(i *isolate, p param) *ast.VariableStatement {
	chck := i.checkpoint()

	if !i.acceptOneOfTypes(token.Var) {
		i.restore(chck)
		return nil
	}

	list := parseVariableDeclarationList(i, p.only(pYield|pAwait).add(pIn))
	if list == nil {
		i.restore(chck)
		return nil
	}

	if !i.acceptOneOfTypes(token.SemiColon) {
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
	if i.acceptOneOfTypes(token.BraceOpen, token.Function, token.Class) {
		i.restore(chck)
		return nil
	}

	// more negative lookaheads
	if i.acceptOneOfTypes(token.Async) {
		if i.acceptOneOfTypes(token.Function) {
			i.restore(chck)
			return nil
		}
	}
	if i.acceptOneOfTypes(token.Let) {
		if i.acceptOneOfTypes(token.BracketOpen) {
			i.restore(chck)
			return nil
		}
	}

	// actual parsing

	// restore checkpoint to make sure lookaheads didn't move position marker
	i.restore(chck)

	expr := parseExpression(i, p.only(pYield|pAwait).add(pIn))

	if !i.acceptOneOfTypes(token.SemiColon) {
		i.restore(chck)
		return nil
	}

	return &ast.ExpressionStatement{
		Expression: expr,
	}
}

func parseIfStatement(i *isolate, p param) *ast.IfStatement {
	chck := i.checkpoint()

	if !i.acceptOneOfTypes(token.If) {
		i.restore(chck)
		return nil
	}

	if !i.acceptOneOfTypes(token.ParOpen) {
		i.restore(chck)
		return nil
	}

	expr := parseExpression(i, p.only(pYield|pAwait).add(pIn))

	if !i.acceptOneOfTypes(token.ParClose) {
		i.restore(chck)
		return nil
	}

	stmt := parseStatement(i, p.only(pYield|pAwait|pReturn))

	chckBeforeWhitespace := i.checkpoint()

	// optional else part
	if i.acceptOneOfTypes(token.Else) {
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
			SwitchStatement: switchStmt,
		}
	}

	return nil
}

func parseContinueStatement(i *isolate, p param) *ast.ContinueStatement {
	chck := i.checkpoint()

	if !i.acceptOneOfTypes(token.Continue) {
		i.restore(chck)
		return nil
	}

	var labelIdent *ast.LabelIdentifier

	if i.acceptOneOfTypes(token.LineTerminator) {
	} else {
		labelIdent = parseLabelIdentifier(i, p.only(pYield|pAwait))
	}

	if !i.acceptOneOfTypes(token.SemiColon) {
		i.restore(chck)
		return nil
	}

	return &ast.ContinueStatement{
		LabelIdentifier: labelIdent,
	}
}

func parseBreakStatement(i *isolate, p param) *ast.BreakStatement {
	chck := i.checkpoint()

	if !i.acceptOneOfTypes(token.Break) {
		i.restore(chck)
		return nil
	}

	var labelIdent *ast.LabelIdentifier

	if i.acceptOneOfTypes(token.LineTerminator) {
	} else {
		labelIdent = parseLabelIdentifier(i, p.only(pYield|pAwait))
	}

	if !i.acceptOneOfTypes(token.SemiColon) {
		i.restore(chck)
		return nil
	}

	return &ast.BreakStatement{
		LabelIdentifier: labelIdent,
	}
}

func parseReturnStatement(i *isolate, p param) *ast.ReturnStatement {
	chck := i.checkpoint()

	if !i.acceptOneOfTypes(token.Return) {
		i.restore(chck)
		return nil
	}

	var expr *ast.Expression

	if i.acceptOneOfTypes(token.LineTerminator) {
	} else {
		expr = parseExpression(i, p.only(pYield|pAwait).add(pIn))
	}

	if !i.acceptOneOfTypes(token.SemiColon) {
		i.restore(chck)
		return nil
	}

	return &ast.ReturnStatement{
		Expression: expr,
	}
}

func parseWithStatement(i *isolate, p param) *ast.WithStatement {
	chck := i.checkpoint()

	if !i.acceptOneOfTypes(token.With) {
		i.restore(chck)
		return nil
	}

	if !i.acceptOneOfTypes(token.ParOpen) {
		i.restore(chck)
		return nil
	}

	expr := parseExpression(i, p.only(pYield|pAwait).add(pIn))
	if expr == nil {
		i.restore(chck)
		return nil
	}

	if !i.acceptOneOfTypes(token.ParClose) {
		i.restore(chck)
		return nil
	}

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

	if !i.acceptOneOfTypes(token.Colon) {
		i.restore(chck)
		return nil
	}

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

	if !i.acceptOneOfTypes(token.Throw) {
		i.restore(chck)
		return nil
	}

	expr := parseExpression(i, p.only(pYield|pAwait).add(pIn))
	if expr == nil {
		i.restore(chck)
		return nil
	}

	if !i.acceptOneOfTypes(token.SemiColon) {
		i.restore(chck)
		return nil
	}

	return &ast.ThrowStatement{
		Expression: expr,
	}
}

func parseTryStatement(i *isolate, p param) *ast.TryStatement {
	chck := i.checkpoint()

	if !i.acceptOneOfTypes(token.Try) {
		i.restore(chck)
		return nil
	}

	blk := parseBlock(i, p.only(pYield|pAwait|pReturn))
	if blk == nil {
		i.restore(chck)
		return nil
	}

	catch := parseCatch(i, p.only(pYield|pAwait|pReturn))

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

	if !i.acceptOneOfTypes(token.Debugger) {
		i.restore(chck)
		return nil
	}

	if !i.acceptOneOfTypes(token.SemiColon) {
		i.restore(chck)
		return nil
	}

	return &ast.DebuggerStatement{}
}

func parseLabelledItem(i *isolate, p param) *ast.LabelledItem {
	chck := i.checkpoint()

	if stmt := parseStatement(i, p.only(pYield|pAwait|pReturn)); stmt != nil {
		return &ast.LabelledItem{
			Statement: stmt,
		}
	} else if functionDeclaration := parseFunctionDeclaration(i, p.only(pYield|pAwait)); functionDeclaration != nil {
		return &ast.LabelledItem{
			FunctionDeclaration: functionDeclaration,
		}
	}

	i.restore(chck)
	return nil
}

func parseSwitchStatement(i *isolate, p param) *ast.SwitchStatement {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Switch) &&
		i.acceptOneOfTypes(token.ParOpen) {
		if expr := parseExpression(i, p.only(pYield|pAwait).add(pIn)); expr != nil {
			if i.acceptOneOfTypes(token.ParClose) {
				if caseBlock := parseCaseBlock(i, p.only(pYield|pAwait|pReturn)); caseBlock != nil {
					return &ast.SwitchStatement{
						Expression: expr,
						CaseBlock:  caseBlock,
					}
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseCaseBlock(i *isolate, p param) *ast.CaseBlock {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.BraceOpen) {
		cc1 := parseCaseClauses(i, p.only(pYield|pAwait|pReturn))   // cc1 is optional
		def := parseDefaultClause(i, p.only(pYield|pAwait|pReturn)) // def is effectively optional
		cc2 := parseCaseClauses(i, p.only(pYield|pAwait|pReturn))   // cc2 is optional
		if i.acceptOneOfTypes(token.BraceClose) {
			return &ast.CaseBlock{
				CaseClauses:   append(cc1, cc2...),
				DefaultClause: def,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseCaseClauses(i *isolate, p param) (clauses []*ast.CaseClause) {
	first := parseCaseClause(i, p.only(pYield|pAwait|pReturn))
	if first == nil {
		return nil
	}

	for {
		next := parseCaseClause(i, p.only(pYield|pAwait|pReturn))
		if next == nil {
			break
		}
		clauses = append(clauses, next)
	}

	return
}

func parseCaseClause(i *isolate, p param) *ast.CaseClause {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Case) {
		if expr := parseExpression(i, p.only(pYield|pAwait).add(pIn)); expr != nil {
			if i.acceptOneOfTypes(token.Colon) {
				stmtList := parseStatementList(i, p.only(pYield|pAwait|pReturn)) // statement list is optional
				return &ast.CaseClause{
					Expression:    expr,
					StatementList: stmtList,
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseDefaultClause(i *isolate, p param) *ast.DefaultClause {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Default) &&
		i.acceptOneOfTypes(token.Colon) {
		stmtList := parseStatementList(i, p.only(pYield|pAwait|pReturn)) // statement list is optional
		return &ast.DefaultClause{
			StatementList: stmtList,
		}
	}

	i.restore(chck)
	return nil
}

func parseCatch(i *isolate, p param) *ast.Catch {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Catch) {
		// the part in parenthesis (including the parenthesis) will become optional in ES10 (?)
		if i.acceptOneOfTypes(token.ParOpen) {
			if catchParameter := parseCatchParameter(i, p.only(pYield|pAwait)); catchParameter != nil {
				if i.acceptOneOfTypes(token.ParClose) {
					if block := parseBlock(i, p.only(pYield|pAwait|pReturn)); block != nil {
						return &ast.Catch{
							CatchParameter: catchParameter,
							Block:          block,
						}
					}
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseCatchParameter(i *isolate, p param) *ast.CatchParameter {
	if bindingIdent := parseBindingIdentifier(i, p.only(pYield|pAwait)); bindingIdent != nil {
		return &ast.CatchParameter{
			BindingIdentifier: bindingIdent,
		}
	} else if bindingPattern := parseBindingPattern(i, p.only(pYield|pAwait)); bindingPattern != nil {
		return &ast.CatchParameter{
			BindingPattern: bindingPattern,
		}
	}
	return nil
}

func parseFinally(i *isolate, p param) *ast.Finally {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Finally) {
		if block := parseBlock(i, p.only(pYield|pAwait|pReturn)); block != nil {
			return &ast.Finally{
				Block: block,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseIterationStatement(i *isolate, p param) *ast.IterationStatement {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Do) {
		if stmt := parseStatement(i, p.only(pYield|pAwait|pReturn)); stmt != nil {
			if i.acceptOneOfTypes(token.While) {
				if i.acceptOneOfTypes(token.ParOpen) {
					if expr := parseExpression(i, p.only(pYield|pAwait).add(pIn)); expr != nil {
						if i.acceptOneOfTypes(token.ParClose) {
							if i.acceptOneOfTypes(token.SemiColon) {
								return &ast.IterationStatement{
									Do:          true,
									While:       true,
									Statement:   stmt,
									Expression1: expr,
								}
							}
						}
					}
				}
			}
		}
	} else if i.acceptOneOfTypes(token.While) {
		if i.acceptOneOfTypes(token.ParOpen) {
			if expr := parseExpression(i, p.only(pYield|pAwait).add(pIn)); expr != nil {
				if i.acceptOneOfTypes(token.ParClose) {
					if stmt := parseStatement(i, p.only(pYield|pAwait|pReturn)); stmt != nil {
						return &ast.IterationStatement{
							While:       true,
							Expression1: expr,
							Statement:   stmt,
						}
					}
				}
			}
		}
	} else if i.acceptOneOfTypes(token.For) {
		return _parseForIterationStatement(i, p)
	}

	i.restore(chck)
	return nil
}

func _parseForIterationStatement(i *isolate, p param) *ast.IterationStatement {
	// token.For was already consumed
	panic("TODO")
}
