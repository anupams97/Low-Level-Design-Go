package main

import "fmt"

func main() {
	o := Originator{state: "hello"}
	c := Caretaker{history: make([]*Memento, 0)}
	fmt.Println(o.state)
	c.addMemento(o.createMemento())
	o.setState("new state")
	fmt.Println(o.state)

	o.restore(c.getLastMemento())
	fmt.Println(o.state)

}
