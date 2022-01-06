package util

var mid, index int
var cur int

func BinarySearch(list []int, item int) int {
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
