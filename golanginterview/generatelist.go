package main

import (
	"fmt"
	"log"
)

func main() {
	l, t := getListAndTarget()
	list, ok := l.([]int)
	if !ok {
		log.Printf("not a list")
	}
	target, ok := t.(int)
	if !ok {
		log.Printf("not a int")
	}
	targetSlice := gSlice(list, target)
	fmt.Println(targetSlice)
	return
}
func gSlice(list []int, target int) []int {
	result := make([]int, 0, len(list)) // type len and capacity
	for _, num := range list {
		if num > target {
			result = append(result, num)
		}
	}
	return result
}

func getListAndTarget() (interface{}, interface{}) {
	return []int{1, 2, 3, 4, 5, 6, 7, 8}, 4
}
