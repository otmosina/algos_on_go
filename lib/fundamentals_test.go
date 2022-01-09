package lib

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverse(t *testing.T) {
	inputString := "abc"
	expected := "cba"
	require.Equal(t, expected, Reverse(inputString))
}

func TestFizzBuzz(t *testing.T) {
	var r string
	var e error
	r, e = FizzBuzz(1, 0)
	require.NotEmpty(t, e)
	require.Empty(t, r)

	r, e = FizzBuzz(7, 7)

	require.Empty(t, e)
	require.Equal(t, "7", r)

	r, e = FizzBuzz(1, 5)

	require.Empty(t, e)
	require.Equal(t, "1 2 Fizz 4 Buzz", r)

	r, e = FizzBuzz(11, 20)

	require.Empty(t, e)
	require.Equal(t, "11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz", r)
}
