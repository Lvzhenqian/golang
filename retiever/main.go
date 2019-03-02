package main

import (
	"fmt"
	"golang/retiever/me"
	real2 "golang/retiever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = me.Retriever{"hahaha"}
	fmt.Printf("%T %v\n",r,r)

	r = real2.Retriever{}
	fmt.Printf("%T %v\n",r,r)
	//fmt.Println(download(r))


}
