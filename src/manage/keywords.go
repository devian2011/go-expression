package manage

type Keyword string

const (
	Fn     Keyword = "fn"
	For    Keyword = "for"
	Each   Keyword = "each"
	In     Keyword = "in"
	If     Keyword = "if"
	ElseIf Keyword = "elseif"
	Else   Keyword = "else"
)

type ConstructionType string

const (
	Function = iota
	Variable
	Return
	
	VarInt
	VarFloat
	VarString
	VarArray
	VarObject
)
