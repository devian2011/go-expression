package parser

type Token interface {
	GetContains() []rune
}

type Stack interface {
}

type Construction struct {
	Type  int
	Token *Token
}

type AST struct {
}

type Parser struct {
}
