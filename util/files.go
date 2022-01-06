package util

import (
	"encoding/json"
	"os"
)

func WriteToFile(arr []int) {
	f, err := os.Create("sorted_list.json")
	Check(err)
	defer f.Close()

	b, err := json.Marshal(arr)
	Check(err)
	f.WriteString(string(b))
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
