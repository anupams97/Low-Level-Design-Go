package _go

import (
	"container/list"
	"fmt"
)

func mainQ() {
	l := list.New() // Create a new linked list

	// Add elements to the front
	l.PushFront(5) // List: [5]
	l.PushFront(1) // List: [1, 5]

	// Add elements to the back
	l.PushBack(10) // List: [1, 5, 10]
	l.PushBack(20) // List: [1, 5, 10, 20]

	// Get front and back elements
	fmt.Println("Front:", l.Front().Value) // Output: 1
	fmt.Println("Back:", l.Back().Value)   // Output: 20

	// Remove front element
	l.Remove(l.Front()) // List: [5, 10, 20]

	// Remove last element
	l.Remove(l.Back()) // List: [5, 10]

	// Print remaining elements
	fmt.Print("Remaining List: ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}
