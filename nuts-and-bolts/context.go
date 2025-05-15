package nutsandbolts

import (
 "context"
 "time"
)

func performTask(ctx context.Context, durationSeconds int) {
	time.After( time.Duration(durationSeconds) * time.Second)
}

type Key string
type Envelope struct {
	value string
}

func performTaskWithValue[T any](ctx context.Context) T {
	// type assertion needed
	v := ctx.Value(Key("Data")).(T)
	return v
}