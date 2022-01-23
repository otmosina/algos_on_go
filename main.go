package main

import (
	"fmt"
	"learn/algos_on_go/lib"
	"learn/algos_on_go/util"
)

var seqSize = 10000000

func variadic(a ...int) {
	for i := range a {
		fmt.Println(i)
	}
}

func main() {
	variadic(1, 2)
	//os.Exit(0)
	memStore := util.NewMemStore()
	if len(memStore.Arr) == 0 {
		fmt.Println("Sequence not read from file...We run generation")
		util.GenerateSequence(seqSize, memStore)
	}
	const lookupItem = 915
	fmt.Printf("Lookup element %d in arr %v... \n", lookupItem, memStore.Arr[0:10])
	fmt.Printf("Position is %d \n", util.BinarySearch(memStore.Arr, lookupItem))
	fmt.Println(util.SelectionSort([]int{3, 2, 1, 321, 2, 1}))
	//^^^^^^^^^^^^^^^^^^^^^^^^^^

	fmt.Println(lib.FizzBuzz(1, 10))
	fmt.Println(lib.Reverse("Reverse This String"))
	fmt.Println(lib.Fibonacci(100))
	util.PrintMemUsage()

}
