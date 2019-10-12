package ast

type IdentifierReference struct {
	Identifier *Identifier
	Yield      bool
	Await      bool
}

type BindingIdentifier struct {
	Identifier *Identifier
	Yield      bool
	Await      bool
}

type LabelIdentifier struct {
	Identifier *Identifier
	Yield      bool
	Await      bool
}

type Identifier struct {
	IdentifierName string
}
