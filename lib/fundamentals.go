package lib

import (
	"errors"
	"fmt"
	"strings"
)

func Reverse(str string) string {
	var result []byte
	lastIndex := len(str) - 1
	for lastIndex >= 0 {
		result = append(result, str[lastIndex])
		lastIndex--
	}
	return string(result)
}

func FizzBuzz(start, stop int) (string, error) {
	var result []string
	if start > stop {
		err := errors.New("Start > Stop. Error!")
		return strings.Join(result, ""), err
	}

	for i := start; i <= stop; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, " FizzBuzz")
		} else if i%5 == 0 {
			result = append(result, " Buzz")

		} else if i%3 == 0 {
			result = append(result, " Fizz")
		} else {
			result = append(result, fmt.Sprintf(" %d", i))
		}
	}
	return strings.TrimSpace(strings.Join(result, "")), nil
}

func Fibonacci(n int) (int, error) {
	var err error
	fib := []int{0, 1}

	if n <= 0 {
		return 0, errors.New("N <= 0")
	}

	if n == 1 {
		return 0, nil
	}

	if n == 2 {
		return 1, nil
	}

	for i := 3; i <= n; i++ {
		fib = append(fib, fib[len(fib)-1]+fib[len(fib)-2])
	}
	// Finonacci(2)

	return fib[len(fib)-1], err
}
