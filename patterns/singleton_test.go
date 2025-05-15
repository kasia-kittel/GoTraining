package patterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {

	s1 := GetInstance()

	assert.IsType(t, &single{}, s1)
	assert.Equal(t, 1, s1.getCallCounter())

	s2 := GetInstance()

	assert.IsType(t, &single{}, s2)
	assert.Equal(t, 2, s2.getCallCounter())

	assert.Equal(t, s1, s2)

}
