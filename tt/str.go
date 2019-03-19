package main

import (
	"fmt"
	"io"
	"strings"
)

type intGen func() int

func (r intGen) Read(p []byte) (n int, err error) {
	next := r()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i <= 10; i++ {
		fmt.Println(f())
	}

}
