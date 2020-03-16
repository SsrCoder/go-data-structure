package main

import (
	"fmt"
	"github.com/SsrCoder/go-data-structure/list"
)

func main() {
	l := list.NewArrayList()
	i := 0
	l.Add(0)
	l.Add(1)
	l.Add(2)
	if err := l.Get(1, &i); err != nil {
		panic(err)
	} else {
		fmt.Println(i)
	}
	fmt.Println(l.IndexOf(1))
}
