package internal

import (
	"fmt"
	"log"
	"time"

	"github.com/urfave/cli/v2"
)

func NewRunConfig(debug bool, portsToCheck cli.StringSlice, started time.Time, timeout int) (*RunConfig, error) {
	portsOpen := make(map[HostPort]bool)
	for _, hostPortValue := range portsToCheck.Value() {
		host, port, err := extractHostAndPortFromString(hostPortValue)
		if err != nil {
			return nil, err
		}
		portsOpen[HostPort{
			host: host,
			port: port,
		}] = false

		if debug {
			log.Printf("Will check for %s:%d.\n", host, port)
		}
	}
	return &RunConfig{
		debug:          debug,
		hostPortsReady: portsOpen,
		started:        started,
		timeout:        time.Duration(timeout) * time.Second,
	}, nil
}

func (r *RunConfig) printDebugStatement(statement string) {
	if r.debug {
		log.Println(statement)
	}
}

func RunLoop(ctx *cli.Context) error {
	portsToCheck := ctx.Value(HostPortsKey).(cli.StringSlice)
	timeout := ctx.Value(TimeoutKey).(int)
	debug := ctx.Value(DebugKey).(bool)

	started := time.Now()
	r, err := NewRunConfig(debug, portsToCheck, started, timeout)
	if err != nil {
		return err
	}

	r.printDebugStatement(fmt.Sprintf("Starting ready checks (for up to %d seconds).", timeout))
	go func(r *RunConfig) {
		for {
			if r.allPortsFound() {
				r.printDebugStatement("Checked all ports. They're ready!")
				break
			}
			r.checkPorts()
			time.Sleep(PortCheckTimeout)
		}
	}(r)

	go func(r *RunConfig) {
		for {
			time.Sleep(DebugLogInterval)
			r.printDebugStatement(fmt.Sprintf("Still running ready checks (%d/%d)...", r.portsFoundCount(), len(r.hostPortsReady)))
		}
	}(r)

	stop := make(chan bool)
	for {
		if time.Now().After(started.Add(r.timeout)) {
			return ErrorTimeoutNotReady
		}
		select {
		case <-time.After(Heartbeat):
			if r.allPortsFound() {
				return nil
			}
		case <-stop:
			return nil
		}
	}
}
