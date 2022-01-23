package lib

import (
	"errors"
	"fmt"
	"math/big"
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
		if i%15 == 0 {
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

//про большие циферки https://blog.devgenius.io/big-int-in-go-handling-large-numbers-is-easy-157cb272dd4f
func Fibonacci(n int) (*big.Int, error) {
	var err error
	var fib [3]*big.Int //{0, 1}
	fib[1] = big.NewInt(0)
	fib[2] = big.NewInt(1)

	if n <= 0 {
		return big.NewInt(-1), errors.New("N <= 0")
	}

	if n == 1 {
		return big.NewInt(0), nil
	}

	if n == 2 {
		return big.NewInt(1), nil
	}

	for i := 3; i <= n; i++ {
		next := big.NewInt(0).Add(fib[1], fib[2])
		fib[0] = fib[1]
		fib[1] = fib[2]
		fib[2] = next
	}
	return fib[len(fib)-1], err
}
