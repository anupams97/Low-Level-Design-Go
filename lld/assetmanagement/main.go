package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	//1, 0, 1, 1, 1, 0, 0
	//1, 1, 1, 1
	//0, 0, 1, 1, 0
	//
	a := []int{0, 0, 1, 1, 0}
	fmt.Println(findMaxLen(a))
}

func findMaxLen(nums []int) int {
	hashmap := map[int]int{0: -1}
	ps := 0
	maxLen := 0
	for i, num := range nums {
		if num == 0 {
			ps += -1
		} else {
			ps += 1
		}
		if idx, ok := hashmap[ps]; ok {
			maxLen = max(maxLen, i-idx)
		} else {
			hashmap[ps] = i
		}
	}
	return maxLen
}
