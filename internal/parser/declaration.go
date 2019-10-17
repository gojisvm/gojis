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

		if !i.acceptTypes(token.Comma) {
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
