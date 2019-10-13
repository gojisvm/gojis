package parser

import "github.com/gojisvm/gojis/internal/parser/ast"

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
