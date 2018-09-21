package main

import "fmt"

type MyReader struct{}

func (source MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}

	return len(b), nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	//reader.Validate(MyReader{})
	r := MyReader{}
	b := make([]byte, 128)

	r.Read(b)
	fmt.Println(string(b))
}

