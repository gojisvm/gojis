package parser

func (p *ECMAScriptLexer) negativeLookahead(rules ...int) bool {
	return contains(p.GetInputStream().LA(1), rules)
}

func (p *ECMAScriptParser) negativeLookahead(rule ...int) bool {
	p.GetTokenStream().Get(p.GetCurrentToken().GetTokenIndex() + 1).GetTokenType()
	return false
}

func (p *ECMAScriptLexer) not(rules ...int) bool {
	return false
}

func (p *ECMAScriptLexer) codePoint(ecp string) bool {
	return false
}

func (p *ECMAScriptLexer) notCodePoint(ecp string) bool {
	return false
}

func contains(needle int, haystack []int) bool {
	for _, straw := range haystack {
		if straw == needle {
			return true
		}
	}
	return false
}
