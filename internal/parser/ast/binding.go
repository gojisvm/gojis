package ast

// BindingElement represents a BindingElement node.
type BindingElement struct {
	SingleNameBinding *SingleNameBinding
	BindingPattern    *BindingPattern
	Initializer       *Initializer
}

// SingleNameBinding represents a SingleNameBinding node.
type SingleNameBinding struct {
	BindingIdentifier *BindingIdentifier
	Initializer       *Initializer
}

// BindingRestElement represents a BindingRestElement node.
type BindingRestElement struct {
	BindingIdentifier *BindingIdentifier
	BindingPattern    *BindingPattern
}

// BindingPattern represents a BindingPattern node.
type BindingPattern struct {
	ObjectBindingPattern *ObjectBindingPattern
	ArrayBindingPattern  *ArrayBindingPattern
}

// ObjectBindingPattern represents a ObjectBindingPattern node.
type ObjectBindingPattern struct {
	BindingRestProperty *BindingRestProperty
	BindingPropertyList *BindingPropertyList
	Comma               bool
}

// ArrayBindingPattern represents a ArrayBindingPattern node.
type ArrayBindingPattern struct {
	Elision            *Elision
	BindingRestElement *BindingRestElement
	BindingElementList *BindingElementList
	Comma              bool
}

// BindingRestProperty represents a BindingRestProperty node.
type BindingRestProperty struct {
	BindingIdentifier *BindingIdentifier
}

// BindingElementList represents a BindingElementList node.
type BindingElementList struct {
	BindingElisionElement []*BindingElisionElement
	Comma                 bool
}

// BindingElisionElement represents a BindingElisionElement node.
type BindingElisionElement struct {
	Elision        *Elision
	BindingElement *BindingElement
}

// BindingPropertyList represents a BindingPropertyList node.
type BindingPropertyList struct {
	BindingProperties []*BindingProperty
}

// BindingProperty represents a BindingProperty node.
type BindingProperty struct {
	SingleNameBinding *SingleNameBinding
	PropertyName      *PropertyName
	BindingElement    *BindingElement
}

// ForBinding represents a ForBinding node.
type ForBinding struct {
	BindingIdentifier *BindingIdentifier
	BindingPattern    *BindingPattern
}

// BindingList represents a BindingList node.
type BindingList struct {
	LexicalBindings []*LexicalBinding
}

// LexicalBinding represents a LexicalBinding node.
type LexicalBinding struct {
	BindingIdentifier *BindingIdentifier
	BindingPattern    *BindingPattern
	Initializer       *Initializer
}
