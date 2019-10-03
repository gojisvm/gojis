package lexer

import "github.com/gojisvm/gojis/internal/parser/token"

// state is a definition of a state. The state returned will be the one that is
// executed next.
type state func(*Lexer) state

// Below this point, all lexer states are defined.

func lexScript(l *Lexer) state {
	return lexToken
}

func lexToken(l *Lexer) state {
	switch r := l.peek(); r {
	default:
		// handle all cases that cannot be expressed in a switch block
	}
	return unexpectedToken
}

// =================
// == identifiers ==
// =================

func lexIdentifierName(l *Lexer) state {
	return lexIdentifierStart
}

func lexIdentifierStart(l *Lexer) state {
	if !l.accept(IdentifierStartPartial) {
		if l.accept(Backslash) {
			// next sequence must be a unicode escape sequence
			errState := acceptUnicodeEscapeSequence(l)
			if errState != nil {
				return errState
			}
		}
	}
	return lexIdentifierPart
}

func lexIdentifierPart(l *Lexer) state {
	for l.acceptMultiple(IdentifierPartPartial) >= 0 {
		if l.accept(Backslash) {
			errState := acceptUnicodeEscapeSequence(l)
			if errState != nil {
				return errState
			}
		} else {
			break
		}
	}

	l.emit(token.IdentifierName)
	return lexToken
}

// ==============
// == comments ==
// ==============

func lexComment(l *Lexer) state {
	if !l.accept(Slash) {
		return tokenMismatch(Slash)
	}

	switch r := l.next(); r {
	case '/':
		return lexSingleLineComment
	case '*':
		return lexMultiLineComment
	}
	l.unread()
	return tokenMismatch(Asterisk, Slash)
}

func lexSingleLineComment(l *Lexer) state {
	l.acceptMultiple(SingleLineCommentChar)
	l.emit(token.SingleLineComment)
	return lexToken
}

func lexMultiLineComment(l *Lexer) state {
	for {
		l.acceptMultiple(MultiLineNotAsteriskChar)
		if !l.accept(Asterisk) {
			return tokenMismatch(Asterisk)
		}
		if l.accept(Slash) {
			break
		}
	}
	l.emit(token.MultiLineComment)
	return lexToken
}

// ===============
// == sequences ==
// ===============

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
