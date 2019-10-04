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
		l.emit(token.NumericLiteral, token.DecimalLiteral)
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
		l.emit(token.NumericLiteral, token.DecimalLiteral)
		return lexToken
	}

	l.accept(Sign) // optional sign
	if l.acceptMultiple(DecimalDigit) < 1 {
		return tokenMismatch(DecimalDigit)
	}

	l.emit(token.NumericLiteral, token.DecimalLiteral)
	return lexToken
}

func lexBinaryIntegerLiteral(l *Lexer) state {
	if !l.accept(BinaryIndicator) {
		return tokenMismatch(BinaryIndicator)
	}
	if l.acceptMultiple(BinaryDigit) < 1 {
		return tokenMismatch(BinaryDigit)
	}

	l.emit(token.NumericLiteral, token.BinaryIntegerLiteral)
	return lexToken
}

func lexOctalIntegerLiteral(l *Lexer) state {
	if !l.accept(OctalIndicator) {
		return tokenMismatch(OctalIndicator)
	}
	if l.acceptMultiple(OctalDigit) < 1 {
		return tokenMismatch(OctalDigit)
	}

	l.emit(token.NumericLiteral, token.OctalIntegerLiteral)
	return lexToken
}

func lexHexIntegerLiteral(l *Lexer) state {
	if !l.accept(HexIndicator) {
		return tokenMismatch(HexIndicator)
	}
	if l.acceptMultiple(HexDigit) < 1 {
		return tokenMismatch(HexDigit)
	}

	l.emit(token.NumericLiteral, token.HexIntegerLiteral)
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

	for l.acceptMultiple(partial) > 0 {
		if l.accept(Backslash) {
			errState := acceptEscapeSequence(l)
			if errState != nil {
				errState = acceptLineTerminatorSequence(l)
				if errState != nil {
					return errState
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
