package main

import (
	"fmt"
)

func main() {
	s := "了看大戏"
	for _ , b :=range []byte(s){
		fmt.Printf("%X ",b)
	}
	fmt.Println()
	for i,ch := range  []rune(s) {
		fmt.Printf("(%d %c) ",i,ch)
	}
}
