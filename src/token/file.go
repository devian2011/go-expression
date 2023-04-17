package token

import (
	"bufio"
	"io"
)

type fileReader struct {
	r    *bufio.Reader
	curr rune
	err  error
}

func (r *fileReader) Next() bool {
	if r.curr, _, r.err = r.r.ReadRune(); r.err != nil {
		if r.err == io.EOF {
			r.err = nil
			return false
		} else {
			return false
		}
	}

	return true
}

func (r *fileReader) HasErrors() bool {
	return r.err != nil
}

func (r *fileReader) Current() rune {
	return r.curr
}

func (r *fileReader) GetError() error {
	return r.err
}
