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
// CharacterEscapeSequence is specified in 11.8.4.
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

func acceptIdentifierPart(l *Lexer) state {
	for l.acceptMultiple(identifierPartPartial) >= 0 {
		if l.accept(backslash) {
			errState := acceptUnicodeEscapeSequence(l)
			if errState != nil {
				return errState
			}
		} else {
			break
		}
	}
	return nil
}

func acceptNotEscapeSequence(l *Lexer) state {
	// 0 DecimalDigit
	if l.accept(zero) {
		if !l.accept(decimalDigit) {
			return tokenMismatch(decimalDigit)
		}
		return nil
	}

	// DecimalDigit but not 0
	if l.accept(nonZeroDigit) {
		return nil
	}

	// x [nla HexDigit]
	if l.negativeLookahead(1, hexDigit) &&
		l.accept(lowerX) {
		return nil
	}

	// x HexDigit [nla HexDigit]
	if l.negativeLookahead(2, hexDigit) &&
		l.acceptSequence(lowerX, hexDigit) {
		return nil
	}

	if l.negativeLookahead(1, hexDigit) && l.negativeLookahead(1, braceOpen) &&
		l.accept(lowerU) {
		return nil
	}

	if l.negativeLookahead(2, hexDigit) &&
		l.acceptSequence(lowerU, hexDigit) {
		return nil
	}

	if l.negativeLookahead(3, hexDigit) &&
		l.acceptSequence(lowerU, hexDigit, hexDigit) {
		return nil
	}

	if l.negativeLookahead(4, hexDigit) &&
		l.acceptSequence(lowerU, hexDigit, hexDigit, hexDigit) {
		return nil
	}

	if l.negativeLookahead(2, hexDigit) &&
		l.acceptSequence(lowerU, braceOpen) {
		return nil
	}

	if l.negativeLookahead(3, hexDigit) && // #79: look ahead of NotCodePoint
		l.acceptSequence(lowerU, braceOpen) {
		if ok, errState := l.acceptEnclosed(acceptNotCodePoint); !ok {
			return errState
		}
		return nil
	}

	if l.negativeLookahead(3, hexDigit) && l.negativeLookahead(3, braceClose) && // #79: look ahead of CodePoint
		l.acceptSequence(lowerU, braceOpen) {
		if ok, errState := l.acceptEnclosed(acceptCodePoint); !ok {
			return errState
		}
		return nil
	}

	return errorf("Expected not escape sequence, but got '%s'", string(l.peek()))
}

func acceptLineContinuation(l *Lexer) state {
	if !l.accept(backslash) {
		return tokenMismatch(backslash)
	}
	if ok, errState := l.acceptEnclosed(acceptLineTerminatorSequence); !ok {
		return errState
	}
	return nil
}

func acceptNotCodePoint(l *Lexer) state {
	if mv := l.mvHex(l.acceptMultiple(hexDigit)); mv <= 0x10FFFF {
		return errorf("MV of hex digits must be > 0x10FFFF, but was %x", mv)
	}
	return nil
}

func acceptCodePoint(l *Lexer) state {
	if mv := l.mvHex(l.acceptMultiple(hexDigit)); mv > 0x10FFFF {
		return errorf("MV of hex digits must be <= 0x10FFFF, but was %x", mv)
	}
	return nil
}
