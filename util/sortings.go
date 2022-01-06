package util

func FindSmallest(arr []int) int {
	var smallestIndex int
	for i := 1; i <= len(arr)-1; i++ {
		if arr[i] < arr[smallestIndex] {
			smallestIndex = i
		}
	}
	return smallestIndex
}

func SelectionSort(arr []int) []int {
	var smallestIndex int
	var newArr []int
	for len(arr) > 0 {
		smallestIndex = FindSmallest(arr)
		newArr = append(newArr, arr[smallestIndex])
		arr = RemoveIndex(arr, smallestIndex)
	}
	return newArr
}

func RemoveIndex(arr []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, arr[:index]...)
	return append(ret, arr[index+1:]...)
}
