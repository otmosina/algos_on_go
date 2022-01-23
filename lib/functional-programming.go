package lib

import (
	"strings"
)

type User struct {
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}

type request []User

const targetSex = "male"

func CountByYears(users request) map[string]int {
	result := make(map[string]int, len(users))

	for _, el := range users {
		bd := strings.Split(el.Birthday, "-")[0]
		if el.Gender != targetSex {
			continue
		}
		if _, ok := result[bd]; ok {
			result[bd]++
		} else {
			result[bd] = 1
		}

	}
	return result

}

func GetSameParity(list []int) (result []int) {
	oddFirst := list[0] % 2
	for _, item := range list {
		currentOdness := item % 2
		if oddFirst == currentOdness {
			result = append(result, item)
		}
	}
	return
}

func isAnagramm(word1, word2 string) bool {
	hash := make(map[byte]int)
	bytes1 := []byte(word1)
	for _, b := range bytes1 {
		if _, ok := hash[b]; ok {
			hash[b]++
		} else {
			hash[b] = 1
		}
	}

	bytes2 := []byte(word2)
	for _, b := range bytes2 {
		if _, ok := hash[b]; ok {
			hash[b]--
		} else {
			return false
		}
	}

	for _, v := range hash {
		if v != 0 {
			return false
		}
	}
	return true
}

func AnagramFilter(target string, words []string) []string {
	result := make([]string, 0)
	for _, word := range words {
		if isAnagramm(target, word) {
			result = append(result, word)
		}
	}
	return result
}
