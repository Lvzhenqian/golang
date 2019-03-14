package main

import (
	"fmt"
	"golang/retiever/me"
	real2 "golang/retiever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = &me.Retriever{"hahaha"}
	inspect(r)

	r = &real2.Retriever{
		"Mozilla/5.0",
		time.Minute,
	}
	inspect(r)
	//fmt.Println(download(r))

	realRetriever := r.(*real2.Retriever)
	fmt.Println(realRetriever.TimeOut)

}

func inspect(r Retriever) {
	fmt.Println("inspecting",r)
	fmt.Printf("> %T %v\n",r,r)
	fmt.Print("> Type switch:")
	switch v := r.(type) {
	case *me.Retriever:
		fmt.Println("Contents:",v.Contents)

	case *real2.Retriever:
		fmt.Println("*Usagent:",v.UserAgent)
		fmt.Printf("%T %v\n",r,r)
	}
	fmt.Println()
}
