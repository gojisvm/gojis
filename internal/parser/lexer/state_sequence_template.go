package lexer

func acceptTemplateCharacter(l *Lexer) state {
	// $ [ lookahead != { ]
	if l.accept(dollar) && !braceOpen.Matches(l.peek()) {
		return nil
	}

	if l.accept(backslash) {
		// escape sequence OR not escape sequence
		if ok, _ := l.acceptEnclosed(acceptEscapeSequence); !ok {
			if ok, _ = l.acceptEnclosed(acceptNotEscapeSequence); !ok {
				// unread the backslash, because LineContinuation wants the backslash
				l.unread()
				if ok, _ = l.acceptEnclosed(acceptLineContinuation); !ok {
					return errorf("Expected escape sequence or not escape sequence or line terminator sequence, but got '%s'", string(l.peek()))
				}
				return nil
			}
			return nil
		}
		return nil
	}

	if l.accept(templateCharacterPartial) {
		return nil
	}

	if ok, errState := l.acceptEnclosed(acceptLineTerminatorSequence); !ok {
		return errState
	}

	return nil
}
