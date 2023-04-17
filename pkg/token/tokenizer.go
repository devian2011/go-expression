package token

import (
	"bufio"
	"expressionlng/pkg/manage"
	"io"
	"os"
)

var sMap = map[rune]int{
	32:  manage.WhiteSpace,
	33:  manage.Not,
	34:  manage.DoubleQuote,
	38:  manage.And,
	39:  manage.SingleQuote,
	40:  manage.OpenBracket,
	41:  manage.ClosedBracket,
	42:  manage.Multiply,
	43:  manage.Plus,
	44:  manage.Comma,
	45:  manage.Minus,
	47:  manage.Division,
	59:  manage.Semicolon,
	60:  manage.Lt,
	61:  manage.Equal,
	62:  manage.Gt,
	91:  manage.OpenSquare,
	93:  manage.ClosedSquare,
	94:  manage.Xor,
	123: manage.OpenCurlyBracket,
	124: manage.Or,
	125: manage.ClosedCurlyBracket,
}

type Token struct {
	tType int
	val   rune
	text  []rune
}

func (t *Token) GetVal() rune {
	return t.val
}

func (t *Token) GetText() []rune {
	return t.text
}

func (t *Token) IsManagedSymbol() bool {
	return manage.Text != t.tType
}

func (t *Token) NotEmpty() bool {
	return len(t.text) > 0 || t.val > 0
}

type Reader interface {
	Next() bool
	HasErrors() bool
	GetError() error
	Current() rune
}

type Stack struct {
	curIndex uint64
	Tokens   []*Token
	curr     *Token
}

func (s *Stack) current() *Token {
	if s.curr == nil {
		s.curr = &Token{}
	}

	return s.curr
}

func (s *Stack) closeNotEmptyToken() {
	if s.current().NotEmpty() {
		s.current().tType = manage.Text
		s.currentClose()
	}
}

func (s *Stack) currentClose() {
	s.Tokens = append(s.Tokens, s.curr)
	s.curr = nil
}

func ParseFile(filePath string) (*Stack, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	reader := &fileReader{
		r:    bufio.NewReader(file),
		curr: 0,
		err:  nil,
	}

	return Parse(reader)
}

func ParseByBytes(code []byte) (*Stack, error) {
	return Parse(fromBytes(code))
}

func ParseByString(code string) (*Stack, error) {
	return Parse(fromString(code))
}

func ParseByReader(reader io.Reader) (*Stack, error) {
	return Parse(fromReader(reader))
}

func Parse(reader Reader) (*Stack, error) {
	stack := &Stack{
		Tokens: make([]*Token, 0, 10000),
		curr:   nil,
	}

	for reader.Next() {
		rn := reader.Current()

		if 7 <= rn && rn <= 13 { // - \r \n \t \a
			continue
		}

		if code, exists := sMap[rn]; exists {
			stack.closeNotEmptyToken()
			stack.current().val = rn
			stack.current().tType = code
			stack.currentClose()
			continue
		}

		stack.current().text = append(stack.current().text, rn)
	}

	if stack.current().NotEmpty() {
		stack.currentClose()
	}

	if reader.HasErrors() {
		return nil, reader.GetError()
	}

	return stack, nil
}
