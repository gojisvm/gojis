package token

// Stash is a data structure that allows to read and unread tokens. When
// unreading tokens, the last read token will be put back onto the stash.
type Stash struct {
	refill func() (Token, bool)

	cached bool
	cache  Token

	done bool
}

func NewStash(refill func() (Token, bool)) *Stash {
	return &Stash{
		refill: refill,
	}
}

func (s *Stash) Read() (t Token, ok bool) {
	s.ensureCache()
	if s.done {
		return Token{}, false
	}

	s.cached = false
	t = s.cache
	ok = true
	return
}

func (s *Stash) Unread() {
	s.cached = true
}

func (s *Stash) ensureCache() {
	if s.cached {
		// noop if there's a token in cache
		return
	}

	token, ok := s.refill()
	if !ok {
		s.done = true
		return
	}

	s.cached = true
	s.cache = token
}
