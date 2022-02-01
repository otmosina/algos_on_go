package lib

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

var TestStack *Stack

func TestMain(m *testing.M) {

	TestStack = NewStack([]string{"ruby", "go"})
	fmt.Println("HALLO  IAM TESTMAIN")
	extVal := m.Run()
	os.Exit(extVal)
}

func TestPop(t *testing.T) {
	el := TestStack.Pop()
	require.Equal(t, "go", el)
	f := reflect.DeepEqual(TestStack.ToArray(), []string{"ruby"})
	require.True(t, f)
}

func TestPush(t *testing.T) {
	TestStack.Push("go")
	f := reflect.DeepEqual(TestStack.ToArray(), []string{"ruby", "go"})
	require.True(t, f)
}

func TestEmpty(t *testing.T) {
	// _ = TestStack.Pop()
	// _ = TestStack.Pop()
	TestStack.Clear()
	// f := reflect.DeepEqual(TestStack.ToArray(), []string{"ruby", "go"})
	require.True(t, TestStack.Empty())
}
