package main

import (
	"bufio"
	"fmt"
	"os"
)

func fileer(f string) *os.File {
	file,_ := os.Create(f)
	//defer file.Close()
	return file
}


func main() {
	w := fileer("d:/1.txt")
	defer w.Close()
	buf := bufio.NewWriter(w)
	_,err:=fmt.Fprintln(buf,"test")
	if err != nil {
		panic(err)
	}
	errs:=buf.Flush()
	if errs!=nil{
		panic(errs)
	}

}