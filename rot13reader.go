package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	r io.Reader
}

func (source rot13Reader) Read(b []byte) (int, error) {
	readSize, readError := source.r.Read(b)

	for i := 0; i < readSize; i++ {
		letter := rune(b[i])

		if unicode.IsLower(letter) {
			b[i] = byte('a') + ((b[i] - byte('a') + 13) % 26)
		} else if unicode.IsUpper(letter) {
			b[i] = byte('A') + ((b[i] - byte('A') + 13) % 26)
		}
	}

	return readSize, readError
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

