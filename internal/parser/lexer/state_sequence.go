package lexer

func acceptEscapeSequence(l *Lexer) state {
	panic("TODO")
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
