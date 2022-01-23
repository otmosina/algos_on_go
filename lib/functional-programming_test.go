package lib

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountByYears(t *testing.T) {
	// expectArr  := []lib.User{lib.User{ "Bronn', gender: 'male', birthday: '1973-03-23' }}

	jsonArr := `[
		{
			"name": "Bronn",
			"gender": "male",
			"birthday": "1973-03-23"
		},
		{
			"name": "Reigar",
			"gender": "male",
			"birthday": "1973-11-03"
		},
		{
			"name": "Eiegon",
			"gender": "male",
			"birthday": "1963-11-03"
		},
		{
			"name": "Sansa",
			"gender": "female",
			"birthday": "2012-11-03"
		},
		{
			"name": "Jon",
			"gender": "male",
			"birthday": "1980-11-03"
		},
		{
			"name": "Robb",
			"gender": "male",
			"birthday": "1980-05-14"
		},
		{
			"name": "Tisha",
			"gender": "female",
			"birthday": "2012-11-03"
		},
		{
			"name": "Rick",
			"gender": "male",
			"birthday": "2012-11-03"
		},
		{
			"name": "Joffrey",
			"gender": "male",
			"birthday": "1999-11-03"
		},
		{
			"name": "Edd",
			"gender": "male",
			"birthday": "1973-11-03"
		}
	]`

	byt := []byte(jsonArr)
	var exampleArr []User
	err := json.Unmarshal(byt, &exampleArr)
	if err != nil {
		fmt.Println(err.Error())
	}

	expected := make(map[string]int)
	expected["1973"] = 3
	expected["1963"] = 1
	expected["1980"] = 2
	expected["2012"] = 1
	expected["1999"] = 1
	res := CountByYears(exampleArr)

	require.Equal(t, expected, res)

}

func TestGetSameParity(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedOutput []int
	}{
		{
			input:          []int{1, 1, 1, 1},
			expectedOutput: []int{1, 1, 1, 1},
		},
		{
			input:          []int{1, 2, 3},
			expectedOutput: []int{1, 3},
		},
		{
			input:          []int{2, 10, 15, 20},
			expectedOutput: []int{2, 10, 20},
		},
		{
			input:          []int{12345, 32154, 112332},
			expectedOutput: []int{12345},
		},
	}
	for _, tc := range testCases {
		require.Equal(t, tc.expectedOutput, GetSameParity(tc.input))
	}
}

func TestAnagramFilter(t *testing.T) {

	// def test_anagramm_filter
	//   assert { anagramm_filter('laser', %w[lazing lazy lacer]) == [] }
	//   assert { anagramm_filter('abba', %w[aabb abcd bbaa dada]) == %w[aabb bbaa] }
	//   assert { anagramm_filter('racer', %w[crazer carer racar caers racer]) == %w[carer racer] }
	// end
	type inputStruct struct {
		target string
		words  []string
	}
	testCases := []struct {
		input          inputStruct
		expectedOutput []string
	}{
		{
			input:          inputStruct{target: "laser", words: []string{"lazing", "lazy", "lacer"}},
			expectedOutput: []string{},
		},
		{
			input:          inputStruct{target: "abba", words: []string{"aabb", "abcd", "bbaa", "dada"}},
			expectedOutput: []string{"aabb", "bbaa"},
		},
		{
			input:          inputStruct{target: "racer", words: []string{"crazer", "carer", "racar", "caers", "racer"}},
			expectedOutput: []string{"carer", "racer"},
		},
	}
	for _, tc := range testCases {
		require.Equal(t, tc.expectedOutput, AnagramFilter(tc.input.target, tc.input.words))
	}
}
