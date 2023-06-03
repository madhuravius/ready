package internal

import "github.com/urfave/cli/v2"

func NewApp() *cli.App {
	return &cli.App{
		Name:  "ready",
		Usage: "wait for a group of hosts and ports to be ready",
		Commands: []*cli.Command{
			{
				Name:    RunKey,
				Aliases: []string{""},
				Usage:   "start ready",
				Action: func(cCtx *cli.Context) error {
					if err := RunLoop(cCtx); err != nil {
						return err
					}
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:     DebugKey,
				Required: false,
				Value:    true,
				Usage:    "if enabled, will print out logs",
			},
			&cli.StringSliceFlag{
				Name:     HostPortsKey,
				Required: true,
				Usage:    "as a csv, specify a range of hosts and ports to check (ex: \"localhost:3000,test:1234\" )",
			},
			&cli.IntFlag{
				Value:    30,
				Name:     TimeoutKey,
				Required: false,
				Usage:    "as an integer, maximum number of seconds to wait and error if ready checks do not all pass by",
			},
		},
	}
}
