package lib

func Reverse(str string) string {
	var result []byte
	lastIndex := len(str) - 1
	for lastIndex >= 0 {
		result = append(result, str[lastIndex])
		lastIndex--
	}
	return string(result)
}
