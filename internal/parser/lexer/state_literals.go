package lexer

import (
	"github.com/gojisvm/gojis/internal/parser/lexer/matcher"
	"github.com/gojisvm/gojis/internal/parser/token"
)

// Defined literals
const (
	keywordNull  = "null"
	keywordTrue  = "true"
	keywordFalse = "false"
)

func lexNull(l *Lexer) state {
	if !l.acceptWord(keywordNull) {
		return unexpectedWord(keywordNull)
	}
	l.emit(token.Null)
	return lexToken
}

func lexBoolean(l *Lexer) state {
	switch l.peek() {
	case 't':
		if !l.acceptWord(keywordTrue) {
			return unexpectedWord(keywordTrue)
		}
	case 'f':
		if !l.acceptWord(keywordFalse) {
			return unexpectedWord(keywordFalse)
		}
	}
	l.emit(token.Boolean)
	return lexToken
}

func lexNumericLiteral(l *Lexer) state {
	if l.accept(zero) {
		// decimal starting with 0 or
		// 0b binary or
		// 0o octal or
		// 0x hex
		switch l.peek() {
		case 'b', 'B':
			return lexBinaryIntegerLiteral
		case 'o', 'O':
			return lexOctalIntegerLiteral
		case 'x', 'X':
			return lexHexIntegerLiteral
		}
	}
	return lexDecimalLiteral
}

func lexDecimalLiteral(l *Lexer) state {
	if l.accept(dot) {
		if l.acceptMultiple(decimalDigit) < 1 {
			return tokenMismatch(decimalDigit)
		}
		return lexOptionalExponentPart
	}
	return lexDecimalIntegerLiteral
}

func lexDecimalIntegerLiteral(l *Lexer) state {
	if l.eof() {
		l.emit(token.DecimalLiteral)
		return lexToken
	}

	// zero has already been accepted in lexNumericLiteral
	if !l.accept(nonZeroDigit) {
		return tokenMismatch(nonZeroDigit)
	}
	l.acceptMultiple(decimalDigit)

	if l.accept(dot) {
		l.acceptMultiple(decimalDigit)
	}
	return lexOptionalExponentPart
}

func lexOptionalExponentPart(l *Lexer) state {
	if !l.accept(exponentIndicator) {
		// no exponent, emit token
		l.emit(token.DecimalLiteral)
		return lexToken
	}

	l.accept(sign) // optional sign
	if l.acceptMultiple(decimalDigit) < 1 {
		return tokenMismatch(decimalDigit)
	}

	l.emit(token.DecimalLiteral)
	return lexToken
}

func lexBinaryIntegerLiteral(l *Lexer) state {
	if !l.accept(binaryIndicator) {
		return tokenMismatch(binaryIndicator)
	}
	if l.acceptMultiple(binaryDigit) < 1 {
		return tokenMismatch(binaryDigit)
	}

	l.emit(token.BinaryIntegerLiteral)
	return lexToken
}

func lexOctalIntegerLiteral(l *Lexer) state {
	if !l.accept(octalIndicator) {
		return tokenMismatch(octalIndicator)
	}
	if l.acceptMultiple(octalDigit) < 1 {
		return tokenMismatch(octalDigit)
	}

	l.emit(token.OctalIntegerLiteral)
	return lexToken
}

func lexHexIntegerLiteral(l *Lexer) state {
	if !l.accept(hexIndicator) {
		return tokenMismatch(hexIndicator)
	}
	if l.acceptMultiple(hexDigit) < 1 {
		return tokenMismatch(hexDigit)
	}

	l.emit(token.HexIntegerLiteral)
	return lexToken
}

func lexStringLiteral(l *Lexer) state {
	var quote matcher.M
	var partial matcher.M

	switch l.peek() {
	case '"':
		quote = doubleQuote
		partial = doubleStringCharacterPartial
	case '\'':
		quote = singleQuote
		partial = singleStringCharacterPartial
	default:
		return tokenMismatch(doubleQuote, singleQuote)
	}

	if !l.accept(quote) {
		return tokenMismatch(quote)
	}

	for n := uint(1); n > 0; n = l.acceptMultiple(partial) {
		if l.accept(backslash) {
			if ok, _ := l.acceptEnclosed(acceptEscapeSequence); !ok { // #71: don't swallow error state
				if ok, _ := l.acceptEnclosed(acceptLineTerminatorSequence); !ok {
					return errorf("'%s' is not a valid escape token", string(l.peek()))
				}
			}
		}
	}

	if !l.accept(quote) {
		return tokenMismatch(quote)
	}

	l.emit(token.StringLiteral)
	return lexToken
}

func lexRegularExpressionLiteral(l *Lexer) state {
	if !l.accept(slash) {
		return tokenMismatch(slash)
	}

	if ok, errState := l.acceptEnclosed(acceptRegularExpressionBody); !ok {
		return errState
	}

	if !l.accept(slash) {
		return tokenMismatch(slash)
	}

	if ok, errState := l.acceptEnclosed(acceptRegularExpressionFlags); !ok {
		return errState
	}

	l.emit(token.RegularExpressionLiteral)
	return lexToken
}
