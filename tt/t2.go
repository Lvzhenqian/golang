package main

import "fmt"

func main() {
	for i := 0; i < 2; i++ {
		var name string
		fmt.Print("Input Name:")
		n, err := fmt.Scanf("%s", &name)
		fmt.Println(n, err, name)
	}
}
