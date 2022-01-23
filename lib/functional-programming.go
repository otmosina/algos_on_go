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
