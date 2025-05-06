package main

import (
	"fmt"
	_go "github.com/anupams97/Low-Level-Design-Go/goplayground/go"
)

func main() {
	a := _go.Stack[int]{}
	a.Push(1)
	a.Push(2)
	val, ok := a.Pop()
	fmt.Println(val, ok)
}
