package ast

type BindingElement struct {
	SingleNameBinding *SingleNameBinding
	BindingPattern    *BindingPattern
	Initializer       *Initializer
}

type SingleNameBinding struct {
	BindingIdentifier *BindingIdentifier
	Initializer       *Initializer
}

type BindingRestElement struct {
	BindingIdentifier *BindingIdentifier
	BindingPattern    *BindingPattern
}

type BindingPattern struct {
	ObjectBindingPattern *ObjectBindingPattern
	ArrayBindingPattern  *ArrayBindingPattern
}

type ObjectBindingPattern struct {
	BindingRestProperty *BindingRestProperty
	BindingPropertyList *BindingPropertyList
	Comma               bool
}

type ArrayBindingPattern struct {
	Elision            *Elision
	BindingRestElement *BindingRestElement
	BindingElementList *BindingElementList
	Comma              bool
}

type BindingRestProperty struct {
	BindingIdentifier *BindingIdentifier
}

type BindingElementList struct {
	BindingElisionElement []*BindingElisionElement
	Comma                 bool
}

type BindingElisionElement struct {
	Elision        *Elision
	BindingElement *BindingElement
}

type BindingPropertyList struct {
	BindingProperties []*BindingProperty
}

type BindingProperty struct {
	SingleNameBinding *SingleNameBinding
	PropertyName      *PropertyName
	BindingElement    *BindingElement
}

type ForBinding struct {
	BindingIdentifier *BindingIdentifier
	BindingPattern    *BindingPattern
}

type BindingList struct {
	LexicalBindings []*LexicalBinding
}

type LexicalBinding struct {
	BindingIdentifier *BindingIdentifier
	BindingPattern    *BindingPattern
	Initializer       *Initializer
}
