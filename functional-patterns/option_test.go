package patterns

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {

	testHost := "testHost"
	testPort := 8080
	testTimeOut := time.Second

	server := NewServer(WithHost(testHost), WithPort(testPort), WithTimeout(testTimeOut))

	assert.Equal(t, testHost, server.host)
	assert.Equal(t, testPort, server.port)
	assert.Equal(t, testTimeOut, server.timeOut)
}
