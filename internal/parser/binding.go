package parser

import (
	"github.com/gojisvm/gojis/internal/parser/ast"
	"github.com/gojisvm/gojis/internal/parser/token"
)

func parseBindingPattern(i *isolate, p param) *ast.BindingPattern {
	if objectBindingPattern := parseObjectBindingPattern(i, p.only(pYield|pAwait)); objectBindingPattern != nil {
		return &ast.BindingPattern{
			ObjectBindingPattern: objectBindingPattern,
		}
	} else if arrayBindingPattern := parseArrayBindingPattern(i, p.only(pYield|pAwait)); arrayBindingPattern != nil {
		return &ast.BindingPattern{
			ArrayBindingPattern: arrayBindingPattern,
		}
	}

	return nil
}

func parseObjectBindingPattern(i *isolate, p param) *ast.ObjectBindingPattern {
	chck := i.checkpoint()

	if !i.acceptOneOfTypes(token.BraceOpen) {
		i.restore(chck)
		return nil
	}

	if i.acceptOneOfTypes(token.BraceClose) {
		return &ast.ObjectBindingPattern{}
	} else if bindingRestProperty := parseBindingRestProperty(i, p.only(pYield|pAwait)); bindingRestProperty != nil {
		if i.acceptOneOfTypes(token.BraceClose) {
			return &ast.ObjectBindingPattern{
				BindingRestProperty: bindingRestProperty,
			}
		}
	} else if bindingPropertyList := parseBindingPropertyList(i, p.only(pYield|pAwait)); bindingPropertyList != nil {
		if i.acceptOneOfTypes(token.BraceClose) {
			return &ast.ObjectBindingPattern{
				BindingPropertyList: bindingPropertyList,
			}
		} else if i.acceptOneOfTypes(token.Comma) {
			bindingRestProperty := parseBindingRestProperty(i, p.only(pYield|pAwait)) // bindingRestProperty is optional
			return &ast.ObjectBindingPattern{
				BindingRestProperty: bindingRestProperty,
				BindingPropertyList: bindingPropertyList,
				Comma:               true,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseBindingRestProperty(i *isolate, p param) *ast.BindingRestProperty {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Ellipsis) {
		if bindingIdentifier := parseBindingIdentifier(i, p.only(pYield|pAwait)); bindingIdentifier != nil {
			return &ast.BindingRestProperty{
				BindingIdentifier: bindingIdentifier,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseBindingPropertyList(i *isolate, p param) *ast.BindingPropertyList {
	chck := i.checkpoint()

	var properties []*ast.BindingProperty

	first := parseBindingProperty(i, p.only(pYield|pAwait))
	if first == nil {
		i.restore(chck)
		return nil
	}
	properties = append(properties, first)

	for { // parse until there are no more binding properties parsable
		beforeComma := i.checkpoint()

		if !i.acceptOneOfTypes(token.Comma) {
			i.restore(chck)
			return nil
		}

		next := parseBindingProperty(i, p.only(pYield|pAwait|pReturn))
		if next == nil {
			i.restore(beforeComma) // comma was consumed, but no binding property, so reset to before comma
			break
		}
		properties = append(properties, next)
	}

	return &ast.BindingPropertyList{
		BindingProperties: properties,
	}
}

func parseBindingProperty(i *isolate, p param) *ast.BindingProperty {
	chck := i.checkpoint()

	if singleNameBinding := parseSingleNameBinding(i, p.only(pYield|pAwait)); singleNameBinding != nil {
		return &ast.BindingProperty{
			SingleNameBinding: singleNameBinding,
		}
	} else if propertyName := parsePropertyName(i, p.only(pYield|pAwait)); propertyName != nil {
		if i.acceptOneOfTypes(token.Colon) {
			if bindingElement := parseBindingElement(i, p.only(pYield|pAwait)); bindingElement != nil {
				return &ast.BindingProperty{
					PropertyName:   propertyName,
					BindingElement: bindingElement,
				}
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseBindingElement(i *isolate, p param) *ast.BindingElement {
	chck := i.checkpoint()

	if singleNameBinding := parseSingleNameBinding(i, p.only(pYield|pAwait)); singleNameBinding != nil {
		return &ast.BindingElement{
			SingleNameBinding: singleNameBinding,
		}
	} else if bindingPattern := parseBindingPattern(i, p.only(pYield|pAwait)); bindingPattern != nil {
		initializer := parseInitializer(i, p.only(pYield|pAwait).add(pIn)) // initializer is optional
		return &ast.BindingElement{
			BindingPattern: bindingPattern,
			Initializer:    initializer,
		}
	}

	i.restore(chck)
	return nil
}

func parseBindingRestElement(i *isolate, p param) *ast.BindingRestElement {
	chck := i.checkpoint()

	if i.acceptOneOfTypes(token.Ellipsis) {
		if bindingIdentifier := parseBindingIdentifier(i, p.only(pYield|pAwait)); bindingIdentifier != nil {
			return &ast.BindingRestElement{
				BindingIdentifier: bindingIdentifier,
			}
		} else if bindingPattern := parseBindingPattern(i, p.only(pYield|pAwait)); bindingPattern != nil {
			return &ast.BindingRestElement{
				BindingPattern: bindingPattern,
			}
		}
	}

	i.restore(chck)
	return nil
}

func parseBindingElementList(i *isolate, p param) *ast.BindingElementList {
	chck := i.checkpoint()

	var elems []*ast.BindingElisionElement

	first := parseBindingElisionElement(i, p.only(pYield|pAwait))
	if first == nil {
		i.restore(chck)
		return nil
	}
	elems = append(elems, first)

	for { // parse until there are no more binding properties parsable
		beforeComma := i.checkpoint()

		if !i.acceptOneOfTypes(token.Comma) {
			i.restore(chck)
			return nil
		}

		next := parseBindingElisionElement(i, p.only(pYield|pAwait|pReturn))
		if next == nil {
			i.restore(beforeComma) // comma was consumed, but no binding element, so reset to before comma
			break
		}
		elems = append(elems, next)
	}

	return &ast.BindingElementList{
		BindingElisionElement: elems,
	}
}

func parseBindingElisionElement(i *isolate, p param) *ast.BindingElisionElement {
	chck := i.checkpoint()

	_ = parseElision(i, 0) // elision is optional

	if bindingElement := parseBindingElement(i, p.only(pYield|pAwait)); bindingElement != nil {
		return &ast.BindingElisionElement{
			BindingElement: bindingElement,
		}
	}

	i.restore(chck)
	return nil
}

func parseArrayBindingPattern(i *isolate, p param) *ast.ArrayBindingPattern {
	chck := i.checkpoint()

	if !i.acceptOneOfTypes(token.BracketOpen) {
		i.restore(chck)
		return nil
	}

	if bindingElementList := parseBindingElementList(i, p.only(pYield|pAwait)); bindingElementList != nil {
		if i.acceptOneOfTypes(token.BracketClose) {
			return &ast.ArrayBindingPattern{
				BindingElementList: bindingElementList,
			}
		} else if i.acceptOneOfTypes(token.Comma) {
			if i.acceptOneOfTypes(token.BracketClose) {
				return &ast.ArrayBindingPattern{
					BindingElementList: bindingElementList,
					Comma:              true,
				}
			} else if elision := parseElision(i, 0); elision != nil {
				if i.acceptOneOfTypes(token.BracketClose) {
					return &ast.ArrayBindingPattern{
						BindingElementList: bindingElementList,
						Elision:            elision,
						Comma:              true,
					}
				} else if bindingRestElement := parseBindingRestElement(i, p.only(pYield|pAwait)); bindingRestElement != nil {
					if i.acceptOneOfTypes(token.BracketClose) {
						return &ast.ArrayBindingPattern{
							BindingElementList: bindingElementList,
							BindingRestElement: bindingRestElement,
							Elision:            elision,
							Comma:              true,
						}
					}
				}
			} else if bindingRestElement := parseBindingRestElement(i, p.only(pYield|pAwait)); bindingRestElement != nil {
				if i.acceptOneOfTypes(token.BracketClose) {
					return &ast.ArrayBindingPattern{
						BindingElementList: bindingElementList,
						BindingRestElement: bindingRestElement,
						Comma:              true,
					}
				}
			}
		}
	} else if elision := parseElision(i, 0); elision != nil {
		if bindingRestElement := parseBindingRestElement(i, p.only(pYield|pAwait)); bindingRestElement != nil {
			if i.acceptOneOfTypes(token.BracketClose) {
				return &ast.ArrayBindingPattern{
					BindingRestElement: bindingRestElement,
					Elision:            elision,
				}
			}
		}
	} else if bindingRestElement := parseBindingRestElement(i, p.only(pYield|pAwait)); bindingRestElement != nil {
		if i.acceptOneOfTypes(token.BracketClose) {
			return &ast.ArrayBindingPattern{
				BindingRestElement: bindingRestElement,
			}
		}
	} else if i.acceptOneOfTypes(token.BracketClose) {
		return &ast.ArrayBindingPattern{}
	}

	i.restore(chck)
	return nil
}

func parseSingleNameBinding(i *isolate, p param) *ast.SingleNameBinding {
	chck := i.checkpoint()

	if bindingIdentifier := parseBindingIdentifier(i, p.only(pYield|pAwait)); bindingIdentifier != nil {
		if initializer := parseInitializer(i, p.only(pYield|pAwait).add(pIn)); initializer != nil {
			return &ast.SingleNameBinding{
				BindingIdentifier: bindingIdentifier,
				Initializer:       initializer,
			}
		}
	}

	i.restore(chck)
	return nil
}
