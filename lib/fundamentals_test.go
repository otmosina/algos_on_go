package lib

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		name       string
		input      int
		expected   int
		assertFunc func(t *testing.T, expected int, actual int, err error)
	}{
		{
			name:     "happyCase2",
			input:    2,
			expected: 1,
			assertFunc: func(t *testing.T, expected int, actual int, err error) {
				require.NoError(t, err)
				require.Equal(t, expected, actual)
			},
		},

		{
			name:     "happyCase10",
			input:    10,
			expected: 34,
			assertFunc: func(t *testing.T, expected int, actual int, err error) {
				require.NoError(t, err)
				require.Equal(t, expected, actual)
			},
		},

		{
			name:     "happyCase15",
			input:    15,
			expected: 377,
			assertFunc: func(t *testing.T, expected int, actual int, err error) {
				require.NoError(t, err)
				require.Equal(t, expected, actual)
			},
		},

		{
			name:     "happyCase25",
			input:    25,
			expected: 46368,
			assertFunc: func(t *testing.T, expected int, actual int, err error) {
				require.NoError(t, err)
				require.Equal(t, expected, actual)
			},
		},
		{
			name:     "happyCase0",
			input:    1,
			expected: 0,
			assertFunc: func(t *testing.T, expected int, actual int, err error) {
				require.NoError(t, err)
				require.Equal(t, expected, actual)
			},
		},
		{
			name:     "errCase-100",
			input:    -100,
			expected: -1,
			assertFunc: func(t *testing.T, expected int, actual int, err error) {
				require.Error(t, err)
				require.Equal(t, expected, actual)
			},
		},
		{
			name:     "errCase0",
			input:    0,
			expected: -1,
			assertFunc: func(t *testing.T, expected int, actual int, err error) {
				require.Error(t, err)
				require.Equal(t, expected, actual)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		fib, err := Fibonacci(tc.input)
		// t.Run(tc.name, tc.assertFunc(t, tc.expected, fib, err))
		t.Run(tc.name, func(t *testing.T) {
			tc.assertFunc(t, tc.expected, fib, err)
		})
		// require.NoError(t, err)
		// require.Equal(t, tc.expected, fib)
	}

}

func TestReverse(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "abc",
			expected: "cba",
		},

		{
			input:    "a",
			expected: "a",
		},
	}

	for i := range testCases {
		tc := testCases[i]
		require.Equal(t, tc.expected, Reverse(tc.input))
	}
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
