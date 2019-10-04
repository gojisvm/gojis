package lexer

func acceptEscapeSequence(l *Lexer) state {
	switch l.peek() {
	case 'u':
		return acceptUnicodeEscapeSequence(l)
	case 'x':
		return acceptHexEscapeSequence(l)
	}

	return errorf("Expected escape sequence, but did not find valid escape token, got '%s'.", string(l.peek()))
}

func acceptCharacterEscapeSequence(l *Lexer) state {
	if !(l.accept(SingleEscapeCharacter) || l.accept(NonEscapeCharacter)) {
		return tokenMismatch(SingleEscapeCharacter, NonEscapeCharacter)
	}
	return nil
}

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

func acceptLineTerminatorSequence(l *Lexer) state {
	if l.accept(_CR) {
		l.accept(_LF)
		return nil
	}
	if !l.accept(LineTerminator) {
		return tokenMismatch(LineTerminator)
	}
	return nil
}
