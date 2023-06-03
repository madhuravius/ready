package internal

import (
	"errors"
	"time"
)

const (
	// RunKey - main key for subcommand to run the app
	RunKey = "run"
	// DebugKey - enable debug mode that gives some more logs
	DebugKey = "debug"
	// HostPortsKey - flag to search on when trying to search against ports
	HostPortsKey = "host-ports"
	// TimeoutKey - flag to wait for application to error if not found
	TimeoutKey = "timeout"
	// DebugLogInterval - amount of time to wait in between each call
	DebugLogInterval = time.Second * 5
	// PortCheckTimeout - amount of time to check if a port is open
	PortCheckTimeout = time.Second
	// Heartbeat - regular heartbeat for all critical functions in core loop
	Heartbeat = time.Millisecond * 100
)

var (
	// ErrorTimeoutNotReady - error that gets raised when ready checks are not met in time
	ErrorTimeoutNotReady = errors.New("err: operation timed out, ready checks not all met")
	// ErrorHostPortBadFormat - error if bad input err
	ErrorHostPortBadFormat = errors.New("err: host port provided in bad format (not - host:port)")
)
