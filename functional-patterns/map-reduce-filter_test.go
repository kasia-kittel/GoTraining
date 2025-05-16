package patterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapFn(t *testing.T) {

	testArray1 := []string{"one", "two", "three"}

	fn := func(s string) int {
		return len(s)
	}

	result := mapFn[string, int](testArray1, fn)
	expected := []int{3, 3, 5}

	assert.Equal(t, expected, result)
}

func TestFilterFn(t *testing.T) {

	testArray1 := []string{"one", "two", "three"}

	fn := func(s string) bool {
		return len(s) > 3
	}

	result := filterFn(testArray1, fn)
	expected := []string{"three"}

	assert.Equal(t, expected, result)
}

func TestReduceFn(t *testing.T) {

	testArray1 := []string{"one", "two", "three"} //11

	fn := func(i int, s string) int {
		return i + len(s)
	}

	result := reduceFn[string, int](testArray1, fn)
	expected := 11

	assert.Equal(t, expected, result)
}
