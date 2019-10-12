package parser

type param uint8

const (
	pUnknown       = 0
	pAwait   param = 1 << iota
	pDefault
	pIn
	pReturn
	pTagged
	pYield
)

func (p param) is(o param) bool {
	return p&o == o
}

func (p param) add(o param) param {
	return p | o
}

func (p param) only(o param) param {
	return p & o
}
