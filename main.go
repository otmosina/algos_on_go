package main

import (
	"fmt"
	"learn/algos_on_go/util"
)

func main() {
	fmt.Println("Hello world!")
	// arr := []int{1, 3, 5, 7, 9}
	// for i, x := range arr {
	// 	fmt.Printf("")
	// }
	// fmt.Println(binarySearch(arr, 1))

	memStore := util.NewMemStore()
	util.GenerateSequence(100000000, memStore)
	// memStore.Print()
	fmt.Println(binarySearch(memStore.Arr, 999998))
}

var mid, index int
var cur int

func binarySearch(list []int, item int) int {
	result := -1
	low := 0
	high := len(list) - 1
	for low <= high {
		mid = int((high - low) / 2)
		index = low + mid
		cur = list[index]

		if cur == item {
			result = index
			break
		}
		if cur > item {
			high = index - 1
		} else {
			low = index + 1
		}
	}
	return result
}
