package ast

type Declaration struct {
	HoistableDeclaration *HoistableDeclaration
	ClassDeclaration     *ClassDeclaration
	LexicalDeclaration   *LexicalDeclaration
}

type HoistableDeclaration struct {
	FunctionDeclaration       *FunctionDeclaration
	GeneratorDeclaration      *GeneratorDeclaration
	AsyncFunctionDeclaration  *AsyncFunctionDeclaration
	AsyncGeneratorDeclaration *AsyncGeneratorDeclaration
}

type FunctionDeclaration struct {
	BindingIdentifier *BindingIdentifier
	FormalParameters  *FormalParameters
	FunctionBody      *FunctionBody
}

type GeneratorDeclaration struct {
	BindingIdentifier *BindingIdentifier
	FormalParameters  *FormalParameters
	GeneratorBody     *GeneratorBody
}

type AsyncFunctionDeclaration struct {
}

type AsyncGeneratorDeclaration struct {
}

type ClassDeclaration struct {
}

type LexicalDeclaration struct {
}
