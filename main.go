package main

import (
	"fmt"
	"learn/algos_on_go/util"
)

var seqSize = 10000000

func main() {
	fmt.Println("Hello world!")
	// arr := []int{1, 3, 5, 7, 9}
	// for i, x := range arr {
	// 	fmt.Printf("")
	// }
	// fmt.Println(binarySearch(arr, 1))

	memStore := util.NewMemStore()
	if len(memStore.Arr) == 0 {
		fmt.Println("Sequence not read from file...We run generation")
		util.GenerateSequence(seqSize, memStore)
	}

	// memStore.Print()
	fmt.Println(util.BinarySearch(memStore.Arr, 915))

}
