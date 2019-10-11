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

func acceptRegularExpressionBody(l *Lexer) state {
	if ok, errState := l.acceptEnclosed(acceptRegularExpressionFirstChar); !ok {
		return errState
	}
	if ok, errState := l.acceptEnclosed(acceptRegularExpressionChars); !ok {
		return errState
	}
	return nil
}

func acceptRegularExpressionFirstChar(l *Lexer) state {
	if !l.accept(regularExpressionFirstCharPartial) {
		if ok, errState := l.acceptEnclosed(acceptRegularExpressionClassOrBackslashSequence); !ok {
			return errState
		}
	}
	return nil
}

func acceptRegularExpressionChars(l *Lexer) state {
	// zero or more regular expression char
	var ok bool
	for {
		ok, _ = l.acceptEnclosed(acceptRegularExpressionChar)
		if !ok {
			break
		}
	}
	return nil
}

func acceptRegularExpressionChar(l *Lexer) state {
	if !l.accept(regularExpressionCharPartial) {
		if ok, errState := l.acceptEnclosed(acceptRegularExpressionClassOrBackslashSequence); !ok {
			return errState
		}
	}
	return nil
}

func acceptRegularExpressionClassOrBackslashSequence(l *Lexer) state {
	if ok, _ := l.acceptEnclosed(acceptRegularExpressionClass); !ok {
		if ok, _ := l.acceptEnclosed(acceptRegularExpressionBackslashSequence); !ok {
			return errorf("Expected regular expression class or backslash sequence, got '%s'", string(l.peek()))
		}
	}
	return nil
}

func acceptRegularExpressionBackslashSequence(l *Lexer) state {
	if !l.accept(backslash) {
		return tokenMismatch(backslash)
	}
	if !l.accept(regularExpressionNonTerminator) {
		return tokenMismatch(regularExpressionNonTerminator)
	}
	return nil
}

func acceptRegularExpressionClass(l *Lexer) state {
	panic("TODO")
}

func acceptRegularExpressionClassChars(l *Lexer) state {
	// zero or more regular expression class char
	var ok bool
	for {
		ok, _ = l.acceptEnclosed(acceptRegularExpressionClassChar)
		if !ok {
			break
		}
	}
	return nil
}

func acceptRegularExpressionClassChar(l *Lexer) state {
	if !l.accept(regularExpressionClassCharPartial) {
		if ok, errState := l.acceptEnclosed(acceptRegularExpressionBackslashSequence); !ok {
			return errState
		}
	}
	return nil
}

func acceptRegularExpressionFlags(l *Lexer) state {
	// zero or more IdentifierPart
	var ok bool
	for {
		ok, _ = l.acceptEnclosed(acceptIdentifierPart)
		if !ok {
			break
		}
	}
	return nil
}
