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
		i.restore(chck)
		return nil
	}
	decls = append(decls, first)

	for { // parse until there are no more variable declarations parsable
		beforeComma := i.checkpoint()

		if !i.acceptOneOfTypes(token.Comma) {
			i.restore(beforeComma)
			break
		}

		next := parseVariableDeclaration(i, p.only(pYield|pAwait|pReturn))
		if next == nil {
			i.restore(beforeComma) // comma was consumed, but no variable declaration, so reset to before comma
			i.fatal("expecting variable declaration after comma")
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

func parseLexicalDeclaration(i *isolate, p param) *ast.LexicalDeclaration {
	chck := i.checkpoint()

	if letOrConst := parseLetOrConst(i, 0); letOrConst != nil {
		if bindingList := parseBindingList(i, p.only(pIn|pYield|pAwait)); bindingList != nil {
			if i.acceptOneOfTypes(token.SemiColon) {
				return &ast.LexicalDeclaration{
					LetOrConst:  letOrConst,
					BindingList: bindingList,
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseLetOrConst(i *isolate, p param) *ast.LetOrConst {
	if i.acceptOneOfTypes(token.Let) {
		return &ast.LetOrConst{
			Let: true,
		}
	} else if i.acceptOneOfTypes(token.Const) {
		return &ast.LetOrConst{
			Const: true,
		}
	}

	return nil
}

func parseMethodDefinition(i *isolate, p param) *ast.MethodDefinition {
	chck := i.checkpoint()

	if propertyName := parsePropertyName(i, p.only(pYield|pAwait)); propertyName != nil {
		if i.acceptOneOfTypes(token.ParOpen) {
			if uniqueFormalParameters := parseUniqueFormalParameters(i, 0); uniqueFormalParameters != nil {
				if i.acceptOneOfTypes(token.ParClose) &&
					i.acceptOneOfTypes(token.BraceOpen) {
					if functionBody := parseFunctionBody(i, 0); functionBody != nil {
						return &ast.MethodDefinition{
							PropertyName:           propertyName,
							UniqueFormalPatameters: uniqueFormalParameters,
							FunctionBody:           functionBody,
						}
					}
				}
			}
		}
	} else if generatorMethod := parseGeneratorMethod(i, p.only(pYield|pAwait)); generatorMethod != nil {
		return &ast.MethodDefinition{
			GeneratorMethod: generatorMethod,
		}
	} else if asyncMethod := parseAsyncMethod(i, p.only(pYield|pAwait)); asyncMethod != nil {
		return &ast.MethodDefinition{
			AsyncMethod: asyncMethod,
		}
	} else if asyncGeneratorMethod := parseAsyncGeneratorMethod(i, p.only(pYield|pAwait)); asyncGeneratorMethod != nil {
		return &ast.MethodDefinition{
			AsyncGeneratorMethod: asyncGeneratorMethod,
		}
	} else if i.acceptOneOfTypes(token.Get) {
		if propertyName := parsePropertyName(i, p.only(pYield|pAwait)); propertyName != nil {
			if i.acceptOneOfTypes(token.ParOpen) &&
				i.acceptOneOfTypes(token.ParClose) &&
				i.acceptOneOfTypes(token.BraceOpen) {
				if functionBody := parseFunctionBody(i, 0); functionBody != nil {
					if i.acceptOneOfTypes(token.BraceClose) {
						return &ast.MethodDefinition{
							PropertyName: propertyName,
							FunctionBody: functionBody,
						}
					}
				}
			}
		}
	} else if i.acceptOneOfTypes(token.Set) {
		if propertyName := parsePropertyName(i, p.only(pYield|pAwait)); propertyName != nil {
			if i.acceptOneOfTypes(token.ParOpen) {
				if propertySetParameterList := parsePropertySetParameterList(i, 0); propertySetParameterList != nil {
					if i.acceptOneOfTypes(token.ParClose) &&
						i.acceptOneOfTypes(token.BraceOpen) {
						if functionBody := parseFunctionBody(i, 0); functionBody != nil {
							if i.acceptOneOfTypes(token.BraceClose) {
								return &ast.MethodDefinition{
									PropertyName:             propertyName,
									PropertySetParameterList: propertySetParameterList,
									FunctionBody:             functionBody,
								}
							}
						}
					}
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parsePropertySetParameterList(i *isolate, p param) *ast.PropertySetParameterList {
	if formalParameter := parseFormalParameter(i, 0); formalParameter != nil {
		return &ast.PropertySetParameterList{
			FormalParameter: formalParameter,
		}
	}

	return nil
}

func parseUniqueFormalParameters(i *isolate, p param) *ast.UniqueFormalParameters {
	if formalParameters := parseFormalParameters(i, p.only(pYield|pAwait)); formalParameters != nil {
		return &ast.UniqueFormalParameters{
			FormalParameters: formalParameters,
		}
	}
	return nil
}

func parseGeneratorMethod(i *isolate, p param) *ast.GeneratorMethod {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Asterisk) {
		if propertyName := parsePropertyName(i, p.only(pYield|pAwait)); propertyName != nil {
			if i.acceptOneOfTypes(token.ParOpen) {
				if uniqueFormalParameters := parseUniqueFormalParameters(i, pYield); uniqueFormalParameters != nil {
					if i.acceptOneOfTypes(token.ParClose) &&
						i.acceptOneOfTypes(token.BraceOpen) {
						if generatorBody := parseGeneratorBody(i, 0); generatorBody != nil {
							if i.acceptOneOfTypes(token.BraceClose) {
								return &ast.GeneratorMethod{
									PropertyName:           propertyName,
									UniqueFormalParameters: uniqueFormalParameters,
									GeneratorBody:          generatorBody,
								}
							}
						}
					}
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseGeneratorDeclaration(i *isolate, p param) *ast.GeneratorDeclaration {
	chck := i.checkpoint()

	if p.is(pDefault) {
		if i.acceptOneOfTypes(token.Function) &&
			i.acceptOneOfTypes(token.Asterisk) &&
			i.acceptOneOfTypes(token.ParOpen) {
			if formalParameters := parseFormalParameters(i, pYield); formalParameters != nil {
				if i.acceptOneOfTypes(token.ParClose) &&
					i.acceptOneOfTypes(token.BraceOpen) {
					if generatorBody := parseGeneratorBody(i, 0); generatorBody != nil {
						return &ast.GeneratorDeclaration{
							FormalParameters: formalParameters,
							GeneratorBody:    generatorBody,
						}
					}
				}
			}
		}
	}

	i.restore(chck)

	if i.acceptOneOfTypes(token.Function) &&
		i.acceptOneOfTypes(token.Asterisk) {
		if bindingIdent := parseBindingIdentifier(i, p.only(pYield|pAwait)); bindingIdent != nil {
			if i.acceptOneOfTypes(token.ParOpen) {
				if formalParameters := parseFormalParameters(i, pYield); formalParameters != nil {
					if i.acceptOneOfTypes(token.ParClose) &&
						i.acceptOneOfTypes(token.BraceOpen) {
						if generatorBody := parseGeneratorBody(i, 0); generatorBody != nil {
							return &ast.GeneratorDeclaration{
								BindingIdentifier: bindingIdent,
								FormalParameters:  formalParameters,
								GeneratorBody:     generatorBody,
							}
						}
					}
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseGeneratorBody(i *isolate, p param) *ast.GeneratorBody {
	if functionBody := parseFunctionBody(i, pYield); functionBody != nil {
		return &ast.GeneratorBody{
			FunctionBody: functionBody,
		}
	}
	return nil
}

func parseAsyncFunctionDeclaration(i *isolate, p param) *ast.AsyncFunctionDeclaration {
	chck := i.checkpoint()

	if p.is(pDefault) {
		if i.acceptOneOfTypes(token.Async) {
			if i.negativeLookahead(token.LineTerminator) { // negative lookahead
				if i.acceptOneOfTypes(token.Function) {
					if i.acceptOneOfTypes(token.ParOpen) {
						if formalParameters := parseFormalParameters(i, pAwait); formalParameters != nil {
							if i.acceptOneOfTypes(token.ParClose) {
								if i.acceptOneOfTypes(token.BraceOpen) {
									if asyncFunctionBody := parseAsyncFunctionBody(i, 0); asyncFunctionBody != nil {
										if i.acceptOneOfTypes(token.BraceClose) {
											return &ast.AsyncFunctionDeclaration{
												FormalParameters:  formalParameters,
												AsyncFunctionBody: asyncFunctionBody,
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	if i.acceptOneOfTypes(token.Async) {
		if i.negativeLookahead(token.LineTerminator) { // negative lookahead
			if i.acceptOneOfTypes(token.Function) {
				if bindingIdent := parseBindingIdentifier(i, p.only(pYield|pAwait)); bindingIdent != nil {
					if i.acceptOneOfTypes(token.ParOpen) {
						if formalParameters := parseFormalParameters(i, pAwait); formalParameters != nil {
							if i.acceptOneOfTypes(token.ParClose) {
								if i.acceptOneOfTypes(token.BraceOpen) {
									if asyncFunctionBody := parseAsyncFunctionBody(i, 0); asyncFunctionBody != nil {
										if i.acceptOneOfTypes(token.BraceClose) {
											return &ast.AsyncFunctionDeclaration{
												BindingIdentifier: bindingIdent,
												FormalParameters:  formalParameters,
												AsyncFunctionBody: asyncFunctionBody,
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseAsyncFunctionBody(i *isolate, p param) *ast.AsyncFunctionBody {
	if functionBody := parseFunctionBody(i, pAwait); functionBody != nil {
		return &ast.AsyncFunctionBody{
			FunctionBody: functionBody,
		}
	}

	return nil
}

func parseAsyncGeneratorDeclaration(i *isolate, p param) *ast.AsyncGeneratorDeclaration {
	chck := i.checkpoint()

	if p.is(pDefault) {
		if i.acceptOneOfTypes(token.Async) {
			if i.negativeLookahead(token.LineTerminator) { // negative lookahead
				if i.acceptOneOfTypes(token.Function) {
					if i.acceptOneOfTypes(token.Asterisk) {
						if i.acceptOneOfTypes(token.ParOpen) {
							if formalParameters := parseFormalParameters(i, pYield|pAwait); formalParameters != nil {
								if i.acceptOneOfTypes(token.ParClose) {
									if i.acceptOneOfTypes(token.BraceOpen) {
										if asyncGeneratorBody := parseAsyncGeneratorBody(i, 0); asyncGeneratorBody != nil {
											if i.acceptOneOfTypes(token.BraceClose) {
												return &ast.AsyncGeneratorDeclaration{
													FormalParameters:   formalParameters,
													AsyncGeneratorBody: asyncGeneratorBody,
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	if i.acceptOneOfTypes(token.Async) {
		if i.negativeLookahead(token.LineTerminator) { // negative lookahead
			if i.acceptOneOfTypes(token.Function) {
				if i.acceptOneOfTypes(token.Asterisk) {
					if bindingIdent := parseBindingIdentifier(i, p.only(pYield|pAwait)); bindingIdent != nil {
						if i.acceptOneOfTypes(token.ParOpen) {
							if formalParameters := parseFormalParameters(i, pYield|pAwait); formalParameters != nil {
								if i.acceptOneOfTypes(token.ParClose) {
									if i.acceptOneOfTypes(token.BraceOpen) {
										if asyncGeneratorBody := parseAsyncGeneratorBody(i, 0); asyncGeneratorBody != nil {
											if i.acceptOneOfTypes(token.BraceClose) {
												return &ast.AsyncGeneratorDeclaration{
													BindingIdentifier:  bindingIdent,
													FormalParameters:   formalParameters,
													AsyncGeneratorBody: asyncGeneratorBody,
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseAsyncGeneratorBody(i *isolate, p param) *ast.AsyncGeneratorBody {
	if functionBody := parseFunctionBody(i, pYield|pAwait); functionBody != nil {
		return &ast.AsyncGeneratorBody{
			FunctionBody: functionBody,
		}
	}

	return nil
}

func parseAsyncGeneratorMethod(i *isolate, p param) *ast.AsyncGeneratorMethod {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Async) {
		if i.negativeLookahead(token.LineTerminator) { // negative lookahead
			if i.acceptOneOfTypes(token.Asterisk) {
				if propertyName := parsePropertyName(i, p.only(pYield|pAwait)); propertyName != nil {
					if i.acceptOneOfTypes(token.ParOpen) {
						if uniqueFormalParameters := parseUniqueFormalParameters(i, pYield|pAwait); uniqueFormalParameters != nil {
							if i.acceptOneOfTypes(token.ParClose) {
								if i.acceptOneOfTypes(token.BraceOpen) {
									if asyncGeneratorBody := parseAsyncGeneratorBody(i, 0); asyncGeneratorBody != nil {
										if i.acceptOneOfTypes(token.BraceClose) {
											return &ast.AsyncGeneratorMethod{
												PropertyName:           propertyName,
												UniqueFormalParameters: uniqueFormalParameters,
												AsyncGeneratorBody:     asyncGeneratorBody,
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseAsyncMethod(i *isolate, p param) *ast.AsyncMethod {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Async) {
		if i.negativeLookahead(token.LineTerminator) { // negative lookahead
			if propertyName := parsePropertyName(i, p.only(pYield|pAwait)); propertyName != nil {
				if i.acceptOneOfTypes(token.ParOpen) {
					if uniqueFormalParameters := parseUniqueFormalParameters(i, pAwait); uniqueFormalParameters != nil {
						if i.acceptOneOfTypes(token.ParClose) {
							if i.acceptOneOfTypes(token.BraceOpen) {
								if asyncFunctionBody := parseAsyncFunctionBody(i, 0); asyncFunctionBody != nil {
									if i.acceptOneOfTypes(token.BraceClose) {
										return &ast.AsyncMethod{
											PropertyName:           propertyName,
											UniqueFormalParameters: uniqueFormalParameters,
											AsyncFunctionBody:      asyncFunctionBody,
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	i.restore(chck)
	return nil
}
