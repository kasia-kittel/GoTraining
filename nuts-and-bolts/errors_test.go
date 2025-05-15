package nutsandbolts

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestCantDoWith42(t *testing.T) {

	_, err := CantDoWith42(42)
	var ae *argError
	assert.True(t, errors.As(err, &ae), "The error should be of type argError")
}

func TestCantDoWith50(t *testing.T) {

	_, err := CantDoWith50(50)
	var ae *argError
	assert.False(t, errors.As(err, &ae), "The error should not be of type argError")
	assert.Error(t, err, "But it is still an error")
}

