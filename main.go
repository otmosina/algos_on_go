package main

import (
	"learn/algos_on_go/lib"
	"learn/algos_on_go/util"
	"math/big"
)

var seqSize = 10000000

func main() {
	// fmt.Println("Hello world!")

	//Work with Binary search
	// memStore := util.NewMemStore()
	// if len(memStore.Arr) == 0 {
	// 	fmt.Println("Sequence not read from file...We run generation")
	// 	util.GenerateSequence(seqSize, memStore)
	// }
	// fmt.Println(util.BinarySearch(memStore.Arr, 915))
	// fmt.Println(util.SelectionSort([]int{3, 2, 1, 321, 2, 1}))
	//^^^^^^^^^^^^^^^^^^^^^^^^^^

	// fmt.Println(lib.Reverse("asd"))

	//прочитать это https://blog.devgenius.io/big-int-in-go-handling-large-numbers-is-easy-157cb272dd4f

	// fmt.Println(lib.Fibonacci(30))
	// fmt.Println()
	lib.Fibonacci(1000000)

	a := big.NewInt(0)
	b := big.NewInt(1)

	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(999), nil)

	for a.Cmp(&limit) < 0 {
		a.Add(a, b)
		a, b = b, a

	}
	util.PrintMemUsage()

	//	fmt.Println(a)
	// fmt.Printf("limit", limit)
	// fmt.Println("-----------------------------")
	// fmt.Printf("limit", limit.ProbablyPrime(20))

}
