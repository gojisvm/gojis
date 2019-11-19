package lexer

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
	if !l.accept(bracketOpen) {
		return tokenMismatch(bracketOpen)
	}

	if ok, err := l.acceptEnclosed(acceptRegularExpressionChars); !ok {
		return err
	}

	if !l.accept(bracketClose) {
		return tokenMismatch(bracketClose)
	}

	return nil
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
	for l.acceptMultiple(identifierPartPartial) > 0 {
		if l.accept(backslash) {
			// next sequence must be a unicode escape sequence
			errState := acceptUnicodeEscapeSequence(l)
			if errState != nil {
				return errState
			}
		}
	}
	return nil
}
