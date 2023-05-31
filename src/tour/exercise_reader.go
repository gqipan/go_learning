package main

import (
	"io"
)

type MyReader struct{}

func (m MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'M':
		b = b + 13
	case 'M' < b && b <= 'Z':
		b = b - 13
	case 'a' <= b && b <= 'm':
		b = b + 13
	case 'm' < b && b <= 'z':
		b = b - 13
	}
	return b
}

func (rr rot13Reader) Read(b []byte) (int, error) {
	n, e := rr.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, e
}

func main() {

}
