package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {
	js := `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

	value := gjson.Get(js,"name.last")

	fmt.Println(value)
}

