package ast

type ParseNode interface {
	Evaluate() interface{}
}
