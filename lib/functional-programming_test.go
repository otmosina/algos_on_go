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
	// def test_get_same_parity
	//   assert { [] == get_same_parity([]) }
	//   assert { [1, 1, 1, 1] == get_same_parity([1, 1, 1, 1]) }
	//   assert { [1, 3] == get_same_parity([1, 2, 3]) }
	//   assert { [2, 10, 20] == get_same_parity([2, 10, 15, 20]) }
	//   assert { [12_345] == get_same_parity([12_345, 32_154, 112_332]) }
	// end
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
