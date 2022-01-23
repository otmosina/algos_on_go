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
