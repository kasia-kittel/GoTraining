package nutsandbolts

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


type timerStub struct {
	stubTime *time.Time
}

func (t *timerStub) Now() time.Time {
	return *t.stubTime
}

type writerMock struct {
	content string
}

func (w *writerMock) Write(p []byte) (n int, err error) {
	w.content = string(p)
	return 0, nil
}

func TestLog(t *testing.T) {

	time := time.Date(2025, time.April, 10, 11, 20, 0, 0, time.UTC)

	timerStub:= &timerStub{ &time }
	writerMock := &writerMock{}

	Log(timerStub, writerMock, "test message", "_prefix")

	expectedLogMsg := "_prefix - 2025-04-10 11:20:00 +0000 UTC: test message"

	assert.Equal(t, expectedLogMsg, writerMock.content)
}

