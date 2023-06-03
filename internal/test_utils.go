package internal

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"net"
	"os"
)

func startWebServer(port int) net.Listener {
	l, err := net.Listen("tcp", fmt.Sprintf("localhost: %d", port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	return l
}

func setCommonContextTestValues(ctx *cli.Context) {
	ctx.Set(TimeoutKey, "1")
	ctx.Set(DebugKey, "false")
}
