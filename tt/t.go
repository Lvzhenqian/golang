package main

import (
	"fmt"
	"strings"
)

func main() {
	var o strings.Builder
	for _,s := range []string{"a","b","c"}{
		fmt.Fprintf(&o,"%s",s)
	}
	fmt.Println(o.String())
}
