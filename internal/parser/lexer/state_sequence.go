package lexer

// acceptEscapeSequence accepts an escape sequence from the given
// lexer. If everything went ok, it will return nil as state, otherwise the
// state will not be nil and will emit an error token. If the returned state is
// not nil, it must be returned by the caller.
//
//	errState := acceptEscapeSequence(l)
//	if errState != nil {
//		return errState
//	}
//
// EscapeSequence and other escape sequences are specified in 11.8.4.
func acceptEscapeSequence(l *Lexer) state {
	switch r := l.peek(); r {
	case 'u':
		return acceptUnicodeEscapeSequence(l)
	case 'x':
		return acceptHexEscapeSequence(l)
	default:
		if l.accept(Zero) {
			if DecimalDigit.Matches(l.peek()) {
				return unexpectedToken
			}
			return nil
		}

		if SingleEscapeCharacter.Matches(r) || NonEscapeCharacter.Matches(r) {
			return acceptCharacterEscapeSequence(l)
		}
	}
	return errorf("Expected escape sequence, but did not find valid escape token, got '%s'.", string(l.peek()))
}

// acceptCharacterEscapeSequence accepts a character escape sequence from the given
// lexer. If everything went ok, it will return nil as state, otherwise the
// state will not be nil and will emit an error token. If the returned state is
// not nil, it must be returned by the caller.
//
//	errState := acceptCharacterEscapeSequence(l)
//	if errState != nil {
//		return errState
//	}
//
// characterEscapeSequence is specified in 11.8.4.
func acceptCharacterEscapeSequence(l *Lexer) state {
	if !(l.accept(SingleEscapeCharacter) || l.accept(NonEscapeCharacter)) {
		return tokenMismatch(SingleEscapeCharacter, NonEscapeCharacter)
	}
	return nil
}

// acceptHexEscapeSequence accepts a hex escape sequence from the given
// lexer. If everything went ok, it will return nil as state, otherwise the
// state will not be nil and will emit an error token. If the returned state is
// not nil, it must be returned by the caller.
//
//	errState := acceptHexEscapeSequence(l)
//	if errState != nil {
//		return errState
//	}
//
// HexEscapeSequence is specified in 11.8.4.
func acceptHexEscapeSequence(l *Lexer) state {
	if !l.accept(LowerX) {
		return tokenMismatch(LowerX)
	}

	for i := 0; i < 2; i++ {
		if !l.accept(HexDigit) {
			return tokenMismatch(HexDigit)
		}
	}

	return nil
}

// acceptUnicodeEscapeSequence accepts a unicode escape sequence from the given
// lexer. If everything went ok, it will return nil as state, otherwise the
// state will not be nil and will emit an error token. If the returned state is
// not nil, it must be returned by the caller.
//
//	errState := acceptUnicodeEscapeSequence(l)
//	if errState != nil {
//		return errState
//	}
//
// UnicodeEscapeSequence is specified in 11.8.4.
func acceptUnicodeEscapeSequence(l *Lexer) state {
	if !l.accept(LowerU) {
		return tokenMismatch(LowerU)
	}

	if l.accept(BraceOpen) {
		n := l.acceptMultiple(HexDigit)
		if n > 6 {
			// actually if MV(HexDigits) > 0x10FFFF (unicode.MaxRune)
			return errorf("Must be a code point (<= 0x10FFFF)")
		}
		if !l.accept(BraceClose) {
			if n < 6 {
				return tokenMismatch(BraceClose, HexDigit)
			}
			return tokenMismatch(BraceClose)
		}
	} else {
		for i := 0; i < 4; i++ {
			if !l.accept(HexDigit) {
				return tokenMismatch(HexDigit)
			}
		}
	}

	return nil
}

// acceptLineTerminatorEscapeSequence accepts a line terminator escape sequence from the given
// lexer. If everything went ok, it will return nil as state, otherwise the
// state will not be nil and will emit an error token. If the returned state is
// not nil, it must be returned by the caller.
//
//	errState := acceptLineTerminatorEscapeSequence(l)
//	if errState != nil {
//		return errState
//	}
//
// LineTerminatorEscapeSequence is specified in 11.8.4.
func acceptLineTerminatorSequence(l *Lexer) state {
	if l.accept(CarriageReturn) {
		l.accept(LineFeed)
		return nil
	}
	if !l.accept(LineTerminator) {
		return tokenMismatch(LineTerminator)
	}
	return nil
}

func acceptRegularExpressionBody(l *Lexer) state {
	panic("TODO")
}

func acceptRegularExpressionFlags(l *Lexer) state {
	panic("TODO")
}
