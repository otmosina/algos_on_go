package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	arr := []int32{1, 3, 5, 7, 9}
	// for i, x := range arr {
	// 	fmt.Printf("")
	// }
	fmt.Println(binarySearch(arr, 1))
}

var mid, index int
var cur int32

func binarySearch(list []int32, item int32) int {
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
