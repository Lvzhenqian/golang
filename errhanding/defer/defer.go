package main

import (
	"fmt"
)

func tryDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	panic("err")
	fmt.Println("3")
}

func main() {
	tryDefer()
}

