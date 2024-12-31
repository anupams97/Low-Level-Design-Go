package main

import "fmt"

func main() {
	l := lousMenu{}
	l.addItem("soup", 100, "manchow")
	l.addItem("roti", 40, "butter")
	l.addItem("paneer", 120, "Kadai")
	display(&l)

}

func display(l menu) {
	it := l.createIterator()
	for it.hasNext() {
		fmt.Println(*it.next())
	}
}
