package ast

type Script struct {
	ScriptBody *ScriptBody
}

type ScriptBody struct {
	StatementList *StatementList
}
