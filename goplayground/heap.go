package goplayground

import (
	"container/heap"
	"fmt"
)

// Define a custom type implementing heap.Interface
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-heap condition
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	h := &MinHeap{5, 3, 8, 1}
	heap.Init(h) // Convert to a min-heap

	heap.Push(h, 2)              // Insert 2
	fmt.Println("Min:", (*h)[0]) // Peek: Smallest element

	fmt.Println("Pop:", heap.Pop(h)) // Removes the smallest element (1)
}
