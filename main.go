package main

import (
	"fmt"
	"learn/algos_on_go/lib"
)

var seqSize = 10000000

func main() {
	fmt.Println("Hello world!")

	//Work with Binary search
	// memStore := util.NewMemStore()
	// if len(memStore.Arr) == 0 {
	// 	fmt.Println("Sequence not read from file...We run generation")
	// 	util.GenerateSequence(seqSize, memStore)
	// }
	// fmt.Println(util.BinarySearch(memStore.Arr, 915))
	// fmt.Println(util.SelectionSort([]int{3, 2, 1, 321, 2, 1}))
	//^^^^^^^^^^^^^^^^^^^^^^^^^^

	fmt.Println(lib.Reverse("asd"))
	fmt.Println(lib.Finonacci(5))

}
