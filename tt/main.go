package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

func main() {
	input := []byte("foo\x00bar")
	buffer := new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.StdEncoding,buffer)
	encoder.Write(input)
	fmt.Println(string(buffer.Bytes()))
}