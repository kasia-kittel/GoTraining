package nutsandbolts

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPerformTask(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go performTask(ctx, 5)

	select {
		case <-ctx.Done():
			assert.Error(t, ctx.Err(), "The task should time out with error")

		case <-time.After(3 * time.Second):
			assert.Fail(t, "Context never cancelled")			
    }

}

func TestPerformTaskWithValue(t *testing.T) {

	value := "value"
	envelope := Envelope{value}
	ctx := context.WithValue(context.Background(), Key("Data"), envelope)

	ret := performTaskWithValue[Envelope](ctx)

	assert.Equal(t, envelope, ret)
	assert.Equal(t, value, ret.value)

}