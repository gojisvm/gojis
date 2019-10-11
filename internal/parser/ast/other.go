package ast

type Initializer struct {
	AssignmentExpression *AssignmentExpression
}

type Elision struct {
	Elision *Elision
}
