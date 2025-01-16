package goplayground

import (
	"fmt"
	"slices"
)

func main() {
	var a []int
	fmt.Println("this is a slice ", a == nil)

	b := make([]string, 3)
	fmt.Printf("I'm a slice with length : %d and capacity %d\n", len(b), cap(b))

	c := []string{"abc"}
	c = append(c, "a", "b")
	fmt.Println(c)

	c = append(c, []string{"xyz"}...)
	fmt.Println(c)

	//can also be copied
	copyc := make([]string, 2)

	// only the first two will be copied because the copyc has capacity of 2
	copy(copyc, c)
	fmt.Println(copyc)

	//changing c won't change copy c
	c[1] = "qwer"
	fmt.Println(c)
	fmt.Println(copyc)

	// slice operator
	fmt.Println(c[1:])

	// both z and c modified
	z := c[1:]
	z[1] = "mmm"
	fmt.Println(z)
	fmt.Println(c)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Check if two slices are equal
	slices.Equal([]int{1, 2, 3}, []int{1, 2, 3}) // true

	// Check if a slice contains an element
	slices.Contains([]string{"apple", "banana", "cherry"}, "banana") // true

	// Find the index of an element in a slice
	slices.Index([]int{10, 20, 30, 40}, 30) // 2

	// Get the minimum and maximum values from a slice
	slices.Min([]int{5, 10, 3, 8}) // 3
	slices.Max([]int{5, 10, 3, 8}) // 10

	// Sort a slice in ascending order
	nums := []int{4, 1, 3, 2}
	slices.Sort(nums) // [1, 2, 3, 4]

	// Reverse the order of elements in a slice
	words := []string{"Go", "is", "awesome"}
	slices.Reverse(words) // ["awesome", "is", "Go"]

	// Sort a slice in descending order using a custom function
	slices.SortFunc([]int{5, 2, 8, 3}, func(a, b int) int { return b - a }) // [8, 5, 3, 2]

	// Insert an element at a specific index
	slices.Insert([]int{1, 2, 4, 5}, 2, 3) // [1, 2, 3, 4, 5]

	// Delete an element at a specific index
	k := slices.Delete([]int{1, 2, 3, 4, 5}, 2, 3) // [1, 2, 4, 5]

	// Clone a slice (create a new copy)
	cloned := slices.Clone([]int{1, 2, 3}) // [1, 2, 3]

	fmt.Println(k, cloned)

}
