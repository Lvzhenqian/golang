package main

import (
	"fmt"
	"reflect"
)

type Empty struct{}

type Set map[interface{}]Empty

var (
	EmptyValue Empty
)

func (s *Set) Add(v interface{}) {
	c := *s
	c[v] = EmptyValue
}

func (s *Set) Delete(v interface{}) {
	delete(*s, v)
}

func (s *Set) Clear() {
	*s = NewSet()
}

func (s *Set) Len() int {
	return len(*s)
}

func NewSet() Set {
	return make(Set)
}

func (s *Set) ToSlice() (r []interface{}) {
	for k := range *s {
		r = append(r, k)
	}
	return
}



func main() {
	list := []interface{}{
		"a", "b", "a", "a", "c", 1, 1, 6, 3, 2,1,1,2,3,5,1,8,55,11,"abc","eee",
	}
	Sets := NewSet()

	for _, v := range list {
		Sets.Add(v)
	}
	a := Sets.ToSlice()

	tp := reflect.TypeOf(a)
	fmt.Println(a, tp)
}
