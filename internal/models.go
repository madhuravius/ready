package internal

import "time"

type HostPort struct {
	host string
	port int
}

type RunConfig struct {
	debug          bool
	hostPortsReady map[HostPort]bool
	started        time.Time
	timeout        time.Duration
}
