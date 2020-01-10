package main

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type A struct {
	Name string
	Versions string
	Order 	map[string]string
}

type B struct {
	Name string
	Versions string
	Order  map[string]string
}

func WriteFile(filepath string) error {
	a := A{
		Name:     "lv",
		Versions: "a1",
		Order: map[string]string{"lv":"haha"},
	}

	b := B{
		Name:     "lb",
		Versions: "b1",
		Order: map[string]string{"b1":"bb"},
	}
	var buf bytes.Buffer
	encodes := yaml.NewEncoder(&buf)

	encodes.Encode(a)
	encodes.Encode(b)

	if f ,err:= os.OpenFile(filepath,os.O_CREATE|os.O_RDWR,0755);err != nil{
		return err
	}else {
		defer f.Close()
		f.Write(buf.Bytes())
	}
	return nil
}

func ReadFile(filepath string) error {
	var (
		a A
		b B
	)
	f, e := os.Open(filepath)
	if e != nil{
		return e
	}
	defer f.Close()
	dc := yaml.NewDecoder(f)
	dc.Decode(&a)
	dc.Decode(&b)
	fmt.Println(a,b)
	return nil
}


func main() {
	ReadFile("./test.yaml")
}
