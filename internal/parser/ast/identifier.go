package ast

// IdentifierReference represents a IdentifierReference node.
type IdentifierReference struct {
	Identifier *Identifier
	Yield      bool
	Await      bool
}

// BindingIdentifier represents a BindingIdentifier node.
type BindingIdentifier struct {
	Identifier *Identifier
	Yield      bool
	Await      bool
}

// LabelIdentifier represents a LabelIdentifier node.
type LabelIdentifier struct {
	Identifier *Identifier
	Yield      bool
	Await      bool
}

// Identifier represents a Identifier node.
type Identifier struct {
	IdentifierName string
}
