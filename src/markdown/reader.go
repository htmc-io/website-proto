package markdown

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
)

type Reader interface {
	Readln() (string, error)
}

type readerImpl struct {
	*bufio.Scanner
	lineNum int
}

// NewReader is to creat a new Reader
func NewReader(path string) (Reader, error) {
	byts, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, e
	}

	scanner := bufio.NewScanner(bytes.NewReader(byts))

	return &readerImpl{scanner, 0}, nil
}

func (rd *readerImpl) Readln() (string, error) {
	if rd.Scan() {
		rd.lineNum++
		return rd.Text(), nil
	}

	if rd.Err() != nil {
		return "", rd.Err()
	}

	return "", io.EOF
}
