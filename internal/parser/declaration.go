package parser

import (
	"github.com/gojisvm/gojis/internal/parser/ast"
	"github.com/gojisvm/gojis/internal/parser/token"
)

func parseDeclaration(i *isolate, p param) *ast.Declaration {
	if hoistableDeclaration := parseHoistableDeclaration(i, p.only(pYield|pAwait)); hoistableDeclaration != nil {
		return &ast.Declaration{
			HoistableDeclaration: hoistableDeclaration,
		}
	}

	if classDeclaration := parseClassDeclaration(i, p.only(pYield|pAwait)); classDeclaration != nil {
		return &ast.Declaration{
			ClassDeclaration: classDeclaration,
		}
	}

	if lexicalDeclaration := parseLexicalDeclaration(i, p.only(pYield|pAwait).add(pIn)); lexicalDeclaration != nil {
		return &ast.Declaration{
			LexicalDeclaration: lexicalDeclaration,
		}
	}

	return nil
}

func parseHoistableDeclaration(i *isolate, p param) *ast.HoistableDeclaration {
	if functionDeclaration := parseFunctionDeclaration(i, p.only(pYield|pAwait|pReturn)); functionDeclaration != nil {
		return &ast.HoistableDeclaration{
			FunctionDeclaration: functionDeclaration,
		}
	}

	if generatorDeclaration := parseGeneratorDeclaration(i, p.only(pYield|pAwait|pReturn)); generatorDeclaration != nil {
		return &ast.HoistableDeclaration{
			GeneratorDeclaration: generatorDeclaration,
		}
	}

	if asyncFunctionDeclaration := parseAsyncFunctionDeclaration(i, p.only(pYield|pAwait|pReturn)); asyncFunctionDeclaration != nil {
		return &ast.HoistableDeclaration{
			AsyncFunctionDeclaration: asyncFunctionDeclaration,
		}
	}

	if asyncGeneratorDeclaration := parseAsyncGeneratorDeclaration(i, p.only(pYield|pAwait|pReturn)); asyncGeneratorDeclaration != nil {
		return &ast.HoistableDeclaration{
			AsyncGeneratorDeclaration: asyncGeneratorDeclaration,
		}
	}

	return nil
}

func parseVariableDeclarationList(i *isolate, p param) *ast.VariableDeclarationList {
	chck := i.checkpoint()

	var decls []*ast.VariableDeclaration

	first := parseVariableDeclaration(i, p.only(pYield|pAwait|pReturn)) // variable declaration list consists of at least one entry
	if first == nil {
		return nil
	}
	decls = append(decls, first)

	for { // parse until there are no more variable declarations parsable
		beforeComma := i.checkpoint()

		if !i.acceptOneOfTypes(token.Comma) {
			i.restore(chck)
			return nil
		}

		next := parseVariableDeclaration(i, p.only(pYield|pAwait|pReturn))
		if next == nil {
			i.restore(beforeComma) // comma was consumed, but no variable declaration, so reset to before comma
			break
		}
		decls = append(decls, next)
	}

	return &ast.VariableDeclarationList{
		VariableDeclarations: decls,
	}
}

func parseVariableDeclaration(i *isolate, p param) *ast.VariableDeclaration {
	chck := i.checkpoint()

	if bindingIdentifier := parseBindingIdentifier(i, p.only(pYield|pAwait)); bindingIdentifier != nil {
		return &ast.VariableDeclaration{
			BindingIdentifier: bindingIdentifier,
			Initializer:       parseInitializer(i, p.only(pIn|pYield|pAwait)), // initializer is optional, so we don't care if nil is returned
		}
	}

	// binding identifier is nil, binding pattern is expected

	bp := parseBindingPattern(i, p.only(pYield|pAwait))
	if bp == nil {
		return nil
	}

	in := parseInitializer(i, p.only(pIn|pYield|pAwait))
	if in == nil { // this initializer is not optional
		i.restore(chck) // parseBindingPattern consumed tokens, so this needs to restore the checkpoint
		return nil
	}

	return &ast.VariableDeclaration{
		BindingPattern: bp,
		Initializer:    in,
	}
}

func parseClassDeclaration(i *isolate, p param) *ast.ClassDeclaration {
	chck := i.checkpoint()

	if !i.acceptOneOfTypes(token.Class) {
		i.restore(chck)
		return nil
	}

	if p.is(pDefault) {
		if classTail := parseClassTail(i, p.only(pYield|pAwait)); classTail != nil {
			return &ast.ClassDeclaration{
				ClassTail: classTail,
			}
		}
	}

	if bindingIdentifier := parseBindingIdentifier(i, p.only(pYield|pAwait)); bindingIdentifier != nil {
		if classTail := parseClassTail(i, p.only(pYield|pAwait)); classTail != nil {
			return &ast.ClassDeclaration{
				BindingIdentifier: bindingIdentifier,
				ClassTail:         classTail,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseClassTail(i *isolate, p param) *ast.ClassTail {
	chck := i.checkpoint()

	classHeritage := parseClassHeritage(i, p.only(pYield|pAwait))

	if !i.acceptOneOfTypes(token.BraceOpen) {
		i.restore(chck)
		return nil
	}

	classBody := parseClassBody(i, p.only(pYield|pAwait))

	if !i.acceptOneOfTypes(token.BraceClose) {
		i.restore(chck)
		return nil
	}

	return &ast.ClassTail{
		ClassHeritage: classHeritage,
		ClassBody:     classBody,
	}
}

func parseClassHeritage(i *isolate, p param) *ast.ClassHeritage {
	chck := i.checkpoint()

	if !i.acceptOneOfTypes(token.Extends) {
		i.restore(chck)
		return nil
	}

	if leftHandSideExpression := parseLeftHandSideExpression(i, p.only(pYield|pAwait)); leftHandSideExpression != nil {
		return &ast.ClassHeritage{
			LeftHandSideExpression: leftHandSideExpression,
		}
	}

	i.restore(chck)
	return nil
}

func parseClassBody(i *isolate, p param) *ast.ClassBody {
	return &ast.ClassBody{
		ClassElementList: parseClassElementList(i, p.only(pYield|pAwait)),
	}
}

func parseClassElementList(i *isolate, p param) *ast.ClassElementList {
	var decls []*ast.ClassElement

	first := parseClassElement(i, p.only(pYield|pAwait))
	if first == nil {
		return nil
	}
	decls = append(decls, first)

	for {
		next := parseClassElement(i, p.only(pYield|pAwait))
		if next == nil {
			break
		}
		decls = append(decls, next)
	}

	return &ast.ClassElementList{
		ClassElements: decls,
	}
}

func parseClassElement(i *isolate, p param) *ast.ClassElement {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.SemiColon) {
		return &ast.ClassElement{}
	}

	static := i.acceptOneOfTypes(token.Static)
	if methodDefinition := parseMethodDefinition(i, p.only(pYield|pAwait)); methodDefinition != nil {
		return &ast.ClassElement{
			Static:           static,
			MethodDefinition: methodDefinition,
		}
	}

	i.restore(chck)
	return nil
}

func parseFunctionDeclaration(i *isolate, p param) *ast.FunctionDeclaration {
	chck := i.checkpoint()

	if !i.acceptOneOfTypes(token.Function) {
		i.restore(chck)
		return nil
	}

	var bindingIdentifier *ast.BindingIdentifier
	if p.is(pDefault) {
		bindingIdentifier = parseBindingIdentifier(i, p.only(pYield|pAwait))
	}

	if i.acceptOneOfTypes(token.ParOpen) {
		if formalParameters := parseFormalParameters(i, 0); formalParameters != nil {
			if i.acceptOneOfTypes(token.ParClose) &&
				i.acceptOneOfTypes(token.BraceOpen) {
				if functionBody := parseFunctionBody(i, 0); functionBody != nil {
					if i.acceptOneOfTypes(token.BraceClose) {
						return &ast.FunctionDeclaration{
							BindingIdentifier: bindingIdentifier,
							FormalParameters:  formalParameters,
							FunctionBody:      functionBody,
						}
					}
				}
			}
		}
	}

	i.restore(chck)
	return nil
}
