package internal

import (
	"fmt"
	"net"
	"os"
)

func startWebServer(port int) net.Listener {
	l, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1: %d", port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	return l
}
