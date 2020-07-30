package parser

import (
	"github.com/gojisvm/gojis/internal/parser/matcher"
	"github.com/gojisvm/gojis/internal/parser/token"
)

type lexFn func(*isolate) bool

var (
	lexFuncs = map[token.Type]lexFn{
		// please keep this sorted lexicographically ascending (a-z)
		token.AndAssign:                lexWord("&="),
		token.Arrow:                    lexWord("=>"),
		token.Assign:                   lexWord("="),
		token.Asterisk:                 lexWord("*"),
		token.Async:                    lexWord("async"),
		token.Await:                    lexWord("await"),
		token.BinaryIntegerLiteral:     lexBinaryIntegerLiteral,
		token.BitwiseAnd:               lexWord("&"),
		token.BitwiseOr:                lexWord("|"),
		token.BitwiseXor:               lexWord("^"),
		token.Boolean:                  lexBoolean,
		token.BraceClose:               lexWord("}"),
		token.BraceOpen:                lexWord("{"),
		token.BracketClose:             lexWord("]"),
		token.BracketOpen:              lexWord("["),
		token.Break:                    lexWord("break"),
		token.Case:                     lexWord("case"),
		token.Catch:                    lexWord("catch"),
		token.Class:                    lexWord("class"),
		token.Colon:                    lexWord(":"),
		token.Comma:                    lexWord(","),
		token.CommonToken:              func(*isolate) bool { return false },
		token.Const:                    lexWord("const"),
		token.Continue:                 lexWord("continue"),
		token.Debugger:                 lexWord("debugger"),
		token.DecimalLiteral:           lexDecimalLiteral,
		token.Default:                  lexWord("default"),
		token.Delete:                   lexWord("delete"),
		token.DivAssign:                lexWord("/="),
		token.Do:                       lexWord("do"),
		token.Dot:                      lexWord("."),
		token.Ellipsis:                 lexWord("..."),
		token.Else:                     lexWord("else"),
		token.Enum:                     lexWord("enum"),
		token.Equals:                   lexWord("=="),
		token.ExclamationMark:          lexWord("!"),
		token.Export:                   lexWord("export"),
		token.Extends:                  lexWord("extends"),
		token.Finally:                  lexWord("finally"),
		token.For:                      lexWord("for"),
		token.Function:                 lexWord("function"),
		token.Get:                      lexWord("get"),
		token.GreaterThan:              lexWord(">"),
		token.GreaterThanOrEqualTo:     lexWord(">="),
		token.HexIntegerLiteral:        lexHexIntegerLiteral,
		token.IdentifierName:           lexIdentifierName,
		token.If:                       lexWord("if"),
		token.Implements:               lexWord("implements"),
		token.Import:                   lexWord("import"),
		token.In:                       lexWord("in"),
		token.Instanceof:               lexWord("instanceof"),
		token.Interface:                lexWord("interface"),
		token.LeftShift:                lexWord("<<"),
		token.LeftShiftAssign:          lexWord("<<="),
		token.LessThan:                 lexWord("<"),
		token.LessThanOrEqualTo:        lexWord("<="),
		token.Let:                      lexWord("let"),
		token.LineTerminator:           lexMatcher(lineTerminator),
		token.LogicalAnd:               lexWord("&&"),
		token.LogicalOr:                lexWord("||"),
		token.Minus:                    lexWord("-"),
		token.MinusAssign:              lexWord("-="),
		token.Modulo:                   lexWord("%"),
		token.ModuloAssign:             lexWord("%="),
		token.MultiLineComment:         func(*isolate) bool { return false },
		token.MultiplyAssign:           lexWord("*="),
		token.NewT:                     lexWord("new"),
		token.NoSubstitutionTemplate:   func(*isolate) bool { return false },
		token.NotEquals:                lexWord("!="),
		token.Null:                     lexWord("null"),
		token.OctalIntegerLiteral:      lexOctalIntegerLiteral,
		token.OrAssign:                 lexWord("|="),
		token.Package:                  lexWord("package"),
		token.ParClose:                 lexWord(")"),
		token.ParOpen:                  lexWord("("),
		token.Plus:                     lexWord("+"),
		token.PlusAssign:               lexWord("+="),
		token.Power:                    lexWord("**"),
		token.PowerAssign:              lexWord("**="),
		token.Private:                  lexWord("private"),
		token.Protected:                lexWord("protected"),
		token.Public:                   lexWord("public"),
		token.Punctuator:               lexPunctuator,
		token.QuestionMark:             lexWord("?"),
		token.RegularExpressionBody:    func(*isolate) bool { return false },
		token.RegularExpressionFlags:   func(*isolate) bool { return false },
		token.Return:                   lexWord("return"),
		token.RightShift:               lexWord(">>"),
		token.RightShiftAssign:         lexWord(">>="),
		token.SemiColon:                lexWord(";"),
		token.Set:                      lexWord("set"),
		token.SignedInteger:            lexSignedInteger,
		token.SingleLineComment:        func(*isolate) bool { return false },
		token.Slash:                    lexWord("/"),
		token.Static:                   lexWord("static"),
		token.StrictEquals:             lexWord("==="),
		token.StrictNotEquals:          lexWord("!=="),
		token.StringLiteral:            lexStringLiteral,
		token.Super:                    lexWord("super"),
		token.Switch:                   lexWord("switch"),
		token.Target:                   lexWord("target"),
		token.TemplateHead:             func(*isolate) bool { return false },
		token.TemplateMiddle:           func(*isolate) bool { return false },
		token.TemplateTail:             func(*isolate) bool { return false },
		token.This:                     lexWord("this"),
		token.Throw:                    lexWord("throw"),
		token.Tilde:                    lexWord("~"),
		token.Try:                      lexWord("try"),
		token.Typeof:                   lexWord("typeof"),
		token.UnsignedRightShift:       lexWord(">>>"),
		token.UnsignedRightShiftAssign: lexWord(">>>="),
		token.UpdateMinus:              lexWord("--"),
		token.UpdatePlus:               lexWord("++"),
		token.Var:                      lexWord("var"),
		token.Void:                     lexWord("void"),
		token.While:                    lexWord("while"),
		token.Whitespace:               lexMatcher(whiteSpace),
		token.With:                     lexWord("with"),
		token.XorAssign:                lexWord("^="),
		token.Yield:                    lexWord("yield"),
	}
)

func (i *isolate) lexForToken(t token.Type) bool {
	fn, ok := lexFuncs[t]
	if !ok || fn == nil {
		panic("No lex function registered for token type " + t.String())
	}
	return fn(i)
}

func (i *isolate) nextRune() (rune, bool) {
	if i.done() {
		var r rune
		return r, false
	}

	r := i.input[i.pos]
	i.pos++
	return r, true
}

func (i *isolate) acceptMatcher(m matcher.M) bool {
	chck := i.checkpoint()

	next, ok := i.nextRune()
	if !ok {
		return false
	}
	if m.Matches(next) {
		return true
	}

	i.restore(chck)
	return false
}

func (i *isolate) acceptMultipleMatcher(m matcher.M) uint {
	var found uint
	for i.acceptMatcher(m) {
		found++
	}
	return found
}

func lexWord(word string) func(*isolate) bool {
	return func(i *isolate) bool {
		chck := i.checkpoint()
		for _, r := range word {
			next, ok := i.nextRune()
			if !ok {
				return false
			}

			if r != next {
				i.restore(chck)
				return false
			}
		}
		return true
	}
}

func lexMatcher(m matcher.M) func(*isolate) bool {
	return func(i *isolate) bool {
		return i.acceptMatcher(m)
	}
}

func lexBinaryIntegerLiteral(i *isolate) bool {
	chck := i.checkpoint()

	if i.acceptMatcher(zero) && i.acceptMatcher(binaryIndicator) {
		var found int
		for i.acceptMatcher(binaryDigit) {
			found++
		}
		if found > 0 {
			return true
		}
	}

	i.restore(chck)
	return false
}

func lexHexIntegerLiteral(i *isolate) bool {
	chck := i.checkpoint()

	if i.acceptMatcher(zero) && i.acceptMatcher(hexIndicator) {
		var found int
		for i.acceptMatcher(hexDigit) {
			found++
		}
		if found > 0 {
			return true
		}
	}

	i.restore(chck)
	return false
}

func lexIdentifierName(i *isolate) bool {
	var start bool
	if start = i.acceptMatcher(identifierStartPartial); !start {
		if i.acceptMatcher(backslash) {
			// next sequence must be a unicode escape sequence
			if !lexUnicodeEscapeSequence(i) {
				return false
			}
			start = true
		}
	}
	if !start {
		return false
	}

	for i.acceptMatcher(identifierPartPartial) {
		if i.acceptMatcher(backslash) {
			// next sequence must be a unicode escape sequence
			if !lexUnicodeEscapeSequence(i) {
				return false
			}
		}
	}

	return true
}

func lexUnicodeEscapeSequence(i *isolate) bool {
	if !i.acceptMatcher(lowerU) {
		return false
	}

	if i.acceptMatcher(braceOpen) {
		n := i.acceptMultipleMatcher(hexDigit)
		if n > 6 {
			return false
		}
		if !i.acceptMatcher(braceClose) {
			if n < 6 {
				return false
			}
			return false
		}
	} else {
		for cnt := 0; cnt < 4; cnt++ {
			if !i.acceptMatcher(hexDigit) {
				return false
			}
		}
	}

	return true
}

func lexBoolean(i *isolate) bool {
	chck := i.checkpoint()
	if lexWord("true")(i) || lexWord("false")(i) {
		return true
	}
	i.restore(chck)
	return false
}

func lexPunctuator(i *isolate) bool {
	return lexWord("~")(i) ||
		lexWord("||")(i) ||
		lexWord("|=")(i) ||
		lexWord("|")(i) ||
		lexWord(">>>=")(i) ||
		lexWord(">>>")(i) ||
		lexWord(">>=")(i) ||
		lexWord(">>")(i) ||
		lexWord(">=")(i) ||
		lexWord(">")(i) ||
		lexWord("=>")(i) ||
		lexWord("===")(i) ||
		lexWord("==")(i) ||
		lexWord("=")(i) ||
		lexWord("<=")(i) ||
		lexWord("<<=")(i) ||
		lexWord("<<")(i) ||
		lexWord("<")(i) ||
		lexWord("+=")(i) ||
		lexWord("++")(i) ||
		lexWord("+")(i) ||
		lexWord("^=")(i) ||
		lexWord("^")(i) ||
		lexWord("%=")(i) ||
		lexWord("%")(i) ||
		lexWord("&=")(i) ||
		lexWord("&&")(i) ||
		lexWord("&")(i) ||
		lexWord("*=")(i) ||
		lexWord("**=")(i) ||
		lexWord("**")(i) ||
		lexWord("*")(i) ||
		lexWord("{")(i) ||
		lexWord("]")(i) ||
		lexWord("[")(i) ||
		lexWord(")")(i) ||
		lexWord("(")(i) ||
		lexWord(".")(i) ||
		lexWord("...")(i) ||
		lexWord("?")(i) ||
		lexWord("!==")(i) ||
		lexWord("!=")(i) ||
		lexWord("!")(i) ||
		lexWord(":")(i) ||
		lexWord(";")(i) ||
		lexWord(",")(i) ||
		lexWord("-=")(i) ||
		lexWord("-")(i) ||
		lexWord("--")(i)
}

func lexStringLiteral(i *isolate) bool {
	var quote matcher.M
	var partial matcher.M

	next, ok := i.nextRune()
	if !ok {
		return false
	}

	switch next {
	case '"':
		quote = doubleQuote
		partial = doubleStringCharacterPartial
	case '\'':
		quote = singleQuote
		partial = singleStringCharacterPartial
	default:
		return false
	}

	if !i.acceptMatcher(quote) {
		return false
	}

	for n := uint(1); n > 0; n = i.acceptMultipleMatcher(partial) {
		if i.acceptMatcher(backslash) {
			_ = lexLineTerminatorSequence // remove warning
			if !(lexEscapeSequence(i) || lexLineTerminatorSequence(i)) {
				return false
			}
		}
	}

	return i.acceptMatcher(quote)
}

func lexEscapeSequence(i *isolate) bool {
	panic("TODO")
}

func lexLineTerminatorSequence(i *isolate) bool {
	panic("TODO")
}

func lexDecimalLiteral(i *isolate) bool {
	if lexDecimalIntegerLiteral(i) {
		if i.acceptMatcher(dot) {
			_ = lexDecimalDigits(i)
		}
		_ = lexExponentPart(i)
		return true
	} else if i.acceptMatcher(dot) {
		if lexDecimalDigits(i) {
			_ = lexExponentPart(i)
			return true
		}
	}
	return false
}

func lexDecimalIntegerLiteral(i *isolate) bool {
	if i.acceptMatcher(zero) {
		return true
	}

	if i.acceptMatcher(nonZeroDigit) {
		_ = lexDecimalDigits(i) // the result doesn't matter, as decimal digits are optional here
	}

	return false
}

func lexDecimalDigits(i *isolate) bool {
	var foundAtLeastOne bool
	for i.acceptMatcher(decimalDigit) {
		foundAtLeastOne = true
	}
	return foundAtLeastOne
}

func lexExponentPart(i *isolate) bool {
	if i.acceptMatcher(exponentIndicator) {
		if lexSignedInteger(i) {
			return true
		}
	}
	return false
}

func lexSignedInteger(i *isolate) bool {
	_ = i.acceptMatcher(sign) // "+" or "-" is effectively optional
	return lexDecimalDigits(i)
}

func lexOctalIntegerLiteral(i *isolate) bool {
	if i.acceptMatcher(zero) {
		if i.acceptMatcher(octalIndicator) {
			return lexOctalDigits(i)
		}
	}
	return false
}

func lexOctalDigits(i *isolate) bool {
	var foundAtLeastOne bool
	for i.acceptMatcher(octalDigit) {
		foundAtLeastOne = true
	}
	return foundAtLeastOne
}
