package _go

import "fmt"

func mainA() {
	var a [2]int
	fmt.Println(a)
	a[1] = 5
	fmt.Println(a)

	b := [3]int{1, 2, 3}
	fmt.Println(b)
	c := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(c)

	d := [...]int{10, 20, 10: 40}
	fmt.Println(d)

	x := [][]int{{1, 5, 3}, {4, 5, 6}}
	fmt.Println(x)

	e := [][]int{{1, 5: 3}, {4, 5: 6}}
	fmt.Println(e)
}
