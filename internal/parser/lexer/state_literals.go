package lexer

import "github.com/gojisvm/gojis/internal/parser/token"

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
	if !l.accept(Zero) {
		if !l.accept(NonZeroDigit) {
			return tokenMismatch(NonZeroDigit)
		}
		l.acceptMultiple(DecimalDigit)
	}

	if l.accept(Dot) {
		l.acceptMultiple(DecimalDigit)
	}
	return lexOptionalExponentPart
}

func lexOptionalExponentPart(l *Lexer) state {
	if !l.accept(ExponentIndicator) {
		// no exponent, emit token
		l.emit(token.NumericLiteral)
		return lexToken
	}

	l.accept(Sign) // optional sign
	if l.acceptMultiple(DecimalDigit) < 1 {
		return tokenMismatch(DecimalDigit)
	}

	l.emit(token.NumericLiteral)
	return lexToken
}

func lexBinaryIntegerLiteral(l *Lexer) state {
	panic("TODO")
}

func lexOctalIntegerLiteral(l *Lexer) state {
	panic("TODO")
}

func lexHexIntegerLiteral(l *Lexer) state {
	panic("TODO")
}
