package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"

	cloudTrceClient "github.com/nekruzvatanshoev/godevtoolkit/pkg/gcp/cloudtrace"
	"github.com/urfave/cli/v2"
)

func traceCmd() *cli.Command {
	return &cli.Command{
		Name:        TraceCmdName,
		Aliases:     []string{"tc"},
		Description: TraceCmdDescription,
		Usage:       TraceCmdUsage,
		Subcommands: []*cli.Command{
			traceDescribeCmd(),
			traceListmd(),
		},
	}
}

func traceDescribeCmd() *cli.Command {
	return &cli.Command{
		Name:    ActionDescribeCmdName,
		Aliases: []string{"desc"},
		Action:  traceDescribe(context.Background()),
	}
}

func traceDescribe(ctx context.Context) cli.ActionFunc {
	return cli.ActionFunc(func(c *cli.Context) error {
		traceId := c.Args().First()
		projectId := c.String(FlagNameGCPProjectId)
		if traceId == "" {
			log.Fatal(errors.New("traceId needs to be specified"))
		}
		trace, err := cloudTrceClient.GetTraceById(traceId, projectId)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(AsJSON(trace))
		return nil
	})
}

func traceListmd() *cli.Command {
	return &cli.Command{
		Name:   ActionListCmdName,
		Action: traceList(context.Background()),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    FlagNameTraceSpanName,
				Aliases: []string{"s"},
				Usage:   "Name of the span",
			},
			&cli.StringFlag{
				Name:    FlagNameTraceMethod,
				Aliases: []string{"m"},
				Usage:   "Method used (GET, POST, PUT, PATCH, etc.)",
			},
			&cli.StringFlag{
				Name:    FlagNameTraceLatency,
				Aliases: []string{"l"},
				Usage:   "Minimum latency (1s, 1m, 1h, 1d, etc.)",
			},
			&cli.StringFlag{
				Name:     FlagNameTraceStartTime,
				Aliases:  []string{"st"},
				Required: true,
				Usage:    "Start time from now (1m, 1h, 1d, etc.)",
			},

			&cli.StringFlag{
				Name:    FlagNameTraceServiceName,
				Aliases: []string{"sn"},
			},
		},
		UsageText: `
		godevtoolkit trace list -s <SpanName> -m <Method> -l <Latency> -st <Start time> -sn <Service Name>
		
		Examples:
			godevtoolkit trace list -s Span.get_entity -m GET, -l 60s -st 6h
		`,
	}
}

func traceList(ctx context.Context) cli.ActionFunc {
	return cli.ActionFunc(func(c *cli.Context) error {
		projectId := c.String(FlagNameGCPProjectId)
		spanName := c.String(FlagNameTraceSpanName)
		latency := c.String(FlagNameTraceLatency)
		method := c.String(FlagNameTraceMethod)
		startTime := c.String(FlagNameTraceStartTime)
		serviceName := c.String(FlagNameTraceServiceName)

		traces, err := cloudTrceClient.GetTracesByFilter(projectId, spanName, latency, method, startTime, serviceName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(AsJSON(traces))

		return nil
	})
}
