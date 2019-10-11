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
		if l.accept(zero) {
			if decimalDigit.Matches(l.peek()) {
				return unexpectedToken
			}
			return nil
		}

		if singleEscapeCharacter.Matches(r) || nonEscapeCharacter.Matches(r) {
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
	if !(l.accept(singleEscapeCharacter) || l.accept(nonEscapeCharacter)) {
		return tokenMismatch(singleEscapeCharacter, nonEscapeCharacter)
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
	if !l.accept(lowerX) {
		return tokenMismatch(lowerX)
	}

	for i := 0; i < 2; i++ {
		if !l.accept(hexDigit) {
			return tokenMismatch(hexDigit)
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
	if !l.accept(lowerU) {
		return tokenMismatch(lowerU)
	}

	if l.accept(braceOpen) {
		n := l.acceptMultiple(hexDigit)
		if n > 6 {
			// actually if MV(HexDigits) > 0x10FFFF (unicode.MaxRune)
			return errorf("Must be a code point (<= 0x10FFFF)")
		}
		if !l.accept(braceClose) {
			if n < 6 {
				return tokenMismatch(braceClose, hexDigit)
			}
			return tokenMismatch(braceClose)
		}
	} else {
		for i := 0; i < 4; i++ {
			if !l.accept(hexDigit) {
				return tokenMismatch(hexDigit)
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
	if l.accept(carriageReturn) {
		l.accept(lineFeed)
		return nil
	}
	if !l.accept(lineTerminator) {
		return tokenMismatch(lineTerminator)
	}
	return nil
}

func acceptRegularExpressionBody(l *Lexer) state {
	panic("TODO")
}

func acceptRegularExpressionFlags(l *Lexer) state {
	panic("TODO")
}
