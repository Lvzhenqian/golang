package me

import "fmt"

type Retriever struct {
	Contents string
}


func (r *Retriever) Get(url string) string {
	return r.Contents
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever:{Contents=%s}",r.Contents)
}