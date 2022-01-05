package util

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type memStore struct {
	Arr []int
}

func NewMemStore() *memStore {
	return &memStore{}
}

func (m *memStore) add(item int) {
	m.Arr = append(m.Arr, item)
}

func (m *memStore) writeToFile() {

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
}
