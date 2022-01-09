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
