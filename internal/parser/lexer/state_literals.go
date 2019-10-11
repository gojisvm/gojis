package lexer

import (
	"github.com/gojisvm/gojis/internal/parser/lexer/matcher"
	"github.com/gojisvm/gojis/internal/parser/token"
)

// Defined literals
const (
	Null  = "null"
	True  = "true"
	False = "false"
)

func lexNull(l *Lexer) state {
	if !l.acceptWord(Null) {
		return unexpectedWord(Null)
	}
	l.emit(token.Null)
	return lexToken
}

func lexBoolean(l *Lexer) state {
	switch l.peek() {
	case 't':
		if !l.acceptWord(True) {
			return unexpectedWord(True)
		}
	case 'f':
		if !l.acceptWord(False) {
			return unexpectedWord(False)
		}
	}
	l.emit(token.Boolean)
	return lexToken
}

func lexNumericLiteral(l *Lexer) state {
	if l.accept(Zero) {
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
	if l.accept(Dot) {
		if l.acceptMultiple(DecimalDigit) < 1 {
			return tokenMismatch(DecimalDigit)
		}
		return lexOptionalExponentPart
	}
	return lexDecimalIntegerLiteral
}

func lexDecimalIntegerLiteral(l *Lexer) state {
	if l.IsEOF() {
		l.emit(token.DecimalLiteral)
		return lexToken
	}

	// zero has already been accepted in lexNumericLiteral
	if !l.accept(NonZeroDigit) {
		return tokenMismatch(NonZeroDigit)
	}
	l.acceptMultiple(DecimalDigit)

	if l.accept(Dot) {
		l.acceptMultiple(DecimalDigit)
	}
	return lexOptionalExponentPart
}

func lexOptionalExponentPart(l *Lexer) state {
	if !l.accept(ExponentIndicator) {
		// no exponent, emit token
		l.emit(token.DecimalLiteral)
		return lexToken
	}

	l.accept(Sign) // optional sign
	if l.acceptMultiple(DecimalDigit) < 1 {
		return tokenMismatch(DecimalDigit)
	}

	l.emit(token.DecimalLiteral)
	return lexToken
}

func lexBinaryIntegerLiteral(l *Lexer) state {
	if !l.accept(BinaryIndicator) {
		return tokenMismatch(BinaryIndicator)
	}
	if l.acceptMultiple(BinaryDigit) < 1 {
		return tokenMismatch(BinaryDigit)
	}

	l.emit(token.BinaryIntegerLiteral)
	return lexToken
}

func lexOctalIntegerLiteral(l *Lexer) state {
	if !l.accept(OctalIndicator) {
		return tokenMismatch(OctalIndicator)
	}
	if l.acceptMultiple(OctalDigit) < 1 {
		return tokenMismatch(OctalDigit)
	}

	l.emit(token.OctalIntegerLiteral)
	return lexToken
}

func lexHexIntegerLiteral(l *Lexer) state {
	if !l.accept(HexIndicator) {
		return tokenMismatch(HexIndicator)
	}
	if l.acceptMultiple(HexDigit) < 1 {
		return tokenMismatch(HexDigit)
	}

	l.emit(token.HexIntegerLiteral)
	return lexToken
}

func lexStringLiteral(l *Lexer) state {
	var quote matcher.M
	var partial matcher.M

	switch l.peek() {
	case '"':
		quote = DoubleQuote
		partial = DoubleStringCharacterPartial
	case '\'':
		quote = SingleQuote
		partial = SingleStringCharacterPartial
	default:
		return tokenMismatch(DoubleQuote, SingleQuote)
	}

	if !l.accept(quote) {
		return tokenMismatch(quote)
	}

	for n := uint(1); n > 0; n = l.acceptMultiple(partial) {
		if l.accept(Backslash) {
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
	if !l.accept(Slash) {
		return tokenMismatch(Slash)
	}

	if ok, errState := l.acceptEnclosed(acceptRegularExpressionBody); !ok {
		return errState
	}

	if !l.accept(Slash) {
		return tokenMismatch(Slash)
	}

	if ok, errState := l.acceptEnclosed(acceptRegularExpressionFlags); !ok {
		return errState
	}

	l.emit(token.RegularExpressionLiteral)
	return lexToken
}
