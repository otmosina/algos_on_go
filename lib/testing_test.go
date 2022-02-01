package lib

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPop(t *testing.T) {
	stack := NewStack([]string{"ruby", "go"})
	el := stack.Pop()
	require.Equal(t, "go", el)
	f := reflect.DeepEqual(stack.ToArray(), []string{"ruby"})
	require.True(t, f)
}
