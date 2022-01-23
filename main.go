package main

import (
	"encoding/json"
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

	fmt.Println("Count By Years Task:")

	u1 := lib.User{"Andrey", "male", "1990-14-01"}
	uArr := []lib.User{u1, lib.User{"Andrey2", "male", "1990-14-01"}, lib.User{"Oli", "female", "1990-14-01"}}
	fmt.Println(uArr)
	resultMap := lib.CountByYears(uArr)

	output, err := json.Marshal(resultMap)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(output))
	util.PrintMemUsage()

}
