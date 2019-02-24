package main

import (
	"fmt"
)

var resume chan int

func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count++
		}
	}()

	return yield
}

func generateInterger() int {
	return <-resume
}

func main() {

	//strings.Split()
	resume = integers()
	for i := 1; i <= 10; i++ {
		s := generateInterger()
		fmt.Printf("%d\n", s)
	}
}
