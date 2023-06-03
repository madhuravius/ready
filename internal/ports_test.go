package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	// TestPort - using this port for all tests
	TestPort = 22000
	// TestPortFailure - using this port for fail tests
	TestPortFailure = 22001
)

func TestExtractHostAndPortFromStringSuccess(t *testing.T) {
	host, port, err := extractHostAndPortFromString("localhost:1234")
	assert.Nil(t, err)
	assert.Equal(t, host, "localhost")
	assert.Equal(t, port, 1234)
}

func TestExtractHostAndPortFromStringFailureBadInput(t *testing.T) {
	host, port, err := extractHostAndPortFromString("localhost")
	assert.Equal(t, err, ErrorHostPortBadFormat)
	assert.Equal(t, host, "")
	assert.Equal(t, port, 0)
}

func TestExtractHostAndPortFromStringFailureBadInt(t *testing.T) {
	host, port, err := extractHostAndPortFromString("localhost:lol")
	assert.NotNil(t, err)
	assert.Equal(t, host, "")
	assert.Equal(t, port, 0)
}

func TestIsPortActiveFalse(t *testing.T) {
	// potentially flaky test, expects nothing to be running on this port
	assert.False(t, isPortActive("localhost", TestPort))
}

func TestIsPortActiveTrue(t *testing.T) {
	defer startWebServer(TestPort).Close()
	assert.True(t, isPortActive("localhost", TestPort))
}
