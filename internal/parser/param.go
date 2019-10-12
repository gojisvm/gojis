package parser

type param uint8

const (
	paramUnknown param = 1 << iota
	paramYield
	paramAwait
	paramDefault
	paramReturn
	paramTagged
	paramIn
)

func (p param) is(o param) bool {
	return p&o == o
}

func (p param) add(o param) param {
	return p | o
}
