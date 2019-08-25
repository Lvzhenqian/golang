package main

import (
	"fmt"
)

func ssort(a []int) {
	for i:=0;i<len(a);i++{
		for j:= i+1; j<len(a);j++{
			if a[i] > a[j]{
				a[i],a[j] = a[j],a[i]
			}
		}
	}
}

func main() {
	b := [...]int{8,7,5,4,3,10,15}
	ssort(b[:])
	fmt.Println(b)
}
