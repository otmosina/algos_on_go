package util

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

type memStore struct {
	Arr []int
}

func NewMemStore() *memStore {
	store := &memStore{}
	store.readFromFile()
	return store
}

func (m *memStore) add(item int) {
	m.Arr = append(m.Arr, item)
}

func (m *memStore) writeToFile() {

}

func (m *memStore) readFromFile() {
	dat, err := os.ReadFile("sorted_list.json")
	// Check(err)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var seq []int
	err = json.Unmarshal(dat, &seq)
	Check(err)
	m.Arr = seq

}

func (m *memStore) Print() {
	fmt.Println(m.Arr)
}

func (m *memStore) Sort() {
	sort.Ints(m.Arr)
}

func randomInt(max int) int {
	// return rand.Intn(1000000)
	return rand.Intn(max)
}

func GenerateSequence(l int, m *memStore) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < l; i++ {
		m.add(randomInt(l))
	}
	m.Sort()
	WriteToFile(m.Arr)
}
