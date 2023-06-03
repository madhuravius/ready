package internal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunLoopSuccess(t *testing.T) {
	defer startWebServer(TestPort).Close()
	app := NewApp()
	err := app.Run([]string{
		"--debugfalse",
		"--timeout=1",
		fmt.Sprintf("--host-ports=localhost:%d", TestPort),
		"run",
	})
	assert.Nil(t, err)
}

func TestRunLoopPartialSuccess(t *testing.T) {
	defer startWebServer(TestPort).Close()
	app := NewApp()
	err := app.Run([]string{
		"--debugfalse",
		"--timeout=1",
		fmt.Sprintf("--host-ports=localhost:%d,localhost:%d", TestPort, TestPortFailure),
		"run",
	})
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrorTimeoutNotReady)
}

func TestRunLoopFailure(t *testing.T) {
	app := NewApp()
	err := app.Run([]string{
		"--debugfalse",
		"--timeout=1",
		fmt.Sprintf("--host-ports=localhost:%d", TestPortFailure),
		"run",
	})
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrorTimeoutNotReady)
}
