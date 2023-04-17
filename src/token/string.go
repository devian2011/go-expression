package token

import (
	"bufio"
	"bytes"
	"io"
)

type stringReader struct {
	reader *bufio.Reader
	curr   rune
	err    error
}

func fromReader(reader io.Reader) *stringReader {
	return &stringReader{reader: bufio.NewReader(reader)}
}

func fromString(code string) *stringReader {
	return &stringReader{reader: bufio.NewReader(bytes.NewBufferString(code))}
}

func fromBytes(code []byte) *stringReader {
	return &stringReader{reader: bufio.NewReader(bytes.NewReader(code))}
}

func (sr *stringReader) Next() bool {
	if sr.curr, _, sr.err = sr.reader.ReadRune(); sr.err != nil {
		if sr.err == io.EOF {
			sr.err = nil
			return false
		} else {
			return false
		}
	}

	return true
}

func (sr *stringReader) HasErrors() bool {
	return sr.err != nil
}

func (sr *stringReader) Current() rune {
	return sr.curr
}

func (sr *stringReader) GetError() error {
	return sr.err
}
