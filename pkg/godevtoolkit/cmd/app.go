package cmd

import "github.com/urfave/cli/v2"

func CreateApp() cli.App {
	return cli.App{
		Name: AppName,
		Commands: []*cli.Command{
			traceCmd(),
		},
		Description: AppDescription,
		Usage:       AppDescription,
		Version:     AppVersion,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     FlagNameGCPProjectId,
				Required: true,
			},
		},
	}
}
