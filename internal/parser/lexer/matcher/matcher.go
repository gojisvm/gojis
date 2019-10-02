package matcher

import (
	"regexp"
	"strings"
)

// M is the definition of a character class, which can tell whether a rune is
// part of that character definition or not.
type M interface {
	// Matches determines, whether the given rune is part of the defined
	// character class.
	Matches(rune) bool
}

type functionMatcher func(rune) bool

func (m functionMatcher) Matches(r rune) bool { return m(r) }

type regexpMatcher struct {
	regexp *regexp.Regexp
}

func (m *regexpMatcher) Matches(r rune) bool {
	return m.regexp.MatchString(string(r))
}

// Matcher constructors

// New creates a new matcher from a given match function.
func New(matchFn functionMatcher) functionMatcher {
	return matchFn
}

// Negate negates whatever matcher is passed into it, meaning that the new
// matcher will return false when the passed matcher returned true and vice
// versa.
func Negate(m M) functionMatcher {
	return func(r rune) bool { return !m.Matches(r) }
}

// String creates a matcher that checks whether a rune is part of the given
// string.
func String(s string) functionMatcher {
	return func(r rune) bool { return strings.IndexRune(s, r) >= 0 }
}

// Regexp creates a matcher that checks whether a rune is matched by the given
// regular expression.
func Regexp(expr string) *regexpMatcher {
	return &regexpMatcher{
		regexp: regexp.MustCompile(expr),
	}
}
