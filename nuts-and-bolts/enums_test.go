package nutsandbolts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransition(t *testing.T) {

	tests := []struct {
		expectedState ServerState
		currentState ServerState
		msg string
	}{
		{expectedState: StateConnected, currentState: StateIdle, msg: "When idle should connect"},
		{expectedState: StateIdle, currentState: StateConnected, msg: "After connected is idle again"},
		{expectedState: StateError, currentState: StateError, msg: "Errors is error"},
	}

	for _, tt := range tests {
		newState := Transition(tt.currentState)
		assert.Equal(t, tt.expectedState, newState)
	}

    assert.Panics(t, func() {
        Transition(11)
    }, "When state not recognized it panics")

}
