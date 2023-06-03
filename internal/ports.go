package internal

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
)

func extractHostAndPortFromString(hostPortValue string) (string, int, error) {
	hostPortSplit := strings.Split(hostPortValue, ":")

	if len(hostPortSplit) < 2 {
		return "", 0, ErrorHostPortBadFormat
	}

	port, err := strconv.Atoi(hostPortSplit[1])
	if err != nil {
		return "", 0, err
	}

	return hostPortSplit[0], port, nil
}

func isPortActive(host string, port int) bool {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, strconv.Itoa(port)), PortCheckTimeout)
	if err != nil {
		return false
	}
	if conn != nil {
		defer func(conn net.Conn) {
			_ = conn.Close()
		}(conn)
	}
	return conn != nil
}

func (r *RunConfig) portsFoundCount() int {
	count := 0
	for _, status := range r.hostPortsReady {
		if status {
			count += 1
		}
	}
	return count
}

func (r *RunConfig) allPortsFound() bool {
	for _, status := range r.hostPortsReady {
		if !status {
			return false
		}
	}
	return true
}

func (r *RunConfig) checkPorts() {
	wg := sync.WaitGroup{}
	for hostPort, status := range r.hostPortsReady {
		if status {
			// already checked to be there, continue
			continue
		}

		// loop through remaining ports to see if any are remaining
		wg.Add(1)
		go func(wg *sync.WaitGroup, hostPort HostPort, r *RunConfig) {
			defer wg.Done()
			if isPortActive(hostPort.host, hostPort.port) {
				r.hostPortsReady[hostPort] = true
				r.printDebugStatement(fmt.Sprintf("Found %s:%d to be available.", hostPort.host, hostPort.port))
			}
		}(&wg, hostPort, r)
	}
	wg.Wait()
}
