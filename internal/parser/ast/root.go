package ast

var _ ParseNode = (*Script)(nil) // ensure that Script implements ParseNode

type Script struct {
}

func (r *Script) Pos() int {
	return 0
}
