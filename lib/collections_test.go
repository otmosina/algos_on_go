package lib

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeSensored(t *testing.T) {
	testCases := []struct {
		inputS   string
		inputW   []string
		expected string
	}{
		{
			inputS:   "When you play the game of thrones, you win or you die",
			inputW:   []string{"die"},
			expected: "When you play the game of thrones, you win or you $#%!",
		},
		{
			inputS:   "chicken chicken? chicken! chicken",
			inputW:   []string{"chicken"},
			expected: "$#%! chicken? chicken! $#%!",
		},

		{
			inputS:   "chicken chicken? chicken! ? chicken",
			inputW:   []string{"?", "chicken"},
			expected: "$#%! chicken? chicken! $#%! $#%!",
		},
	}

	for i := range testCases {
		tc := testCases[i]
		require.Equal(t, tc.expected, MakeSensored(tc.inputS, tc.inputW))
	}

}

// class CompareVersionsTest < Minitest::Test
//   def test_compare_versions
//     assert { compare_versions('0.2', '0.1') == 1 }
//     assert { compare_versions('0.1', '0.2') == -1 }
//     assert { compare_versions('4.2', '4.2').zero? }
//     assert { compare_versions('0.2', '0.12') == -1 }
//     assert { compare_versions('3.2', '4.12') == -1 }
//     assert { compare_versions('3.2', '2.12') == 1 }
//   end
// end

func TestCompareVersion(t *testing.T) {
	testCases := []struct {
		inputS1, inputS2 string
		expected         int8
	}{
		{
			inputS1:  "0.2",
			inputS2:  "0.1",
			expected: 1,
		},
		{
			inputS1:  "0.1",
			inputS2:  "0.2",
			expected: -1,
		},
		{
			inputS1:  "4.2",
			inputS2:  "4.2",
			expected: 0,
		},
		{
			inputS1:  "0.2",
			inputS2:  "0.12",
			expected: -1,
		},
		{
			inputS1:  "3.2",
			inputS2:  "4.12",
			expected: -1,
		},
		{
			inputS1:  "3.2",
			inputS2:  "2.12",
			expected: 1,
		},
		// TODO: handle error scenario
		// {
		// 	inputS1:   "Error",
		// 	inputS2: "2.12",
		// 	expected: 1,
		// },
	}
	for i := range testCases {
		tc := testCases[i]
		require.Equal(t, tc.expected, CompareVersion(tc.inputS1, tc.inputS2))
	}
}
func TestBuildQueryString(t *testing.T) {
	testCases := []struct {
		input    map[string]string
		expected string
	}{
		{
			input:    map[string]string{},
			expected: "",
		},
		{
			input:    map[string]string{"page": "1"},
			expected: "page=1",
		},
		{
			input:    map[string]string{"per": "10", "page": "1"},
			expected: "page=1&per=10",
		},
		{
			input:    map[string]string{"a": "10", "s": "Wow", "d": "3.2", "z": "89"},
			expected: "a=10&d=3.2&s=Wow&z=89",
		},
		{
			input:    map[string]string{"param": "all", "param1": "true"},
			expected: "param=all&param1=true",
		},
	}

	for i := range testCases {
		tc := testCases[i]
		require.Equal(t, tc.expected, BuildQueryString(tc.input))
	}
}
