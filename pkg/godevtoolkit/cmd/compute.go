package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	computeClient "github.com/nekruzvatanshoev/godevtoolkit/pkg/gcp/computeengine"
	"github.com/urfave/cli/v2"
)

func computeCmd() *cli.Command {
	return &cli.Command{
		Name:        ComputeCmdName,
		Aliases:     []string{"ce"},
		Description: ComputeCmdDescription,
		Usage:       ComputeCmdUsage,
		Subcommands: []*cli.Command{
			computeDescribeCmd(),
			computeListCmd(),
			computeResetCmd(),
		},
	}
}

func computeDescribeCmd() *cli.Command {
	return &cli.Command{
		Name:    ActionDescribeCmdName,
		Aliases: []string{"desc"},
		Action:  computeDescribe(context.Background()),
	}
}

func computeDescribe(ctx context.Context) cli.ActionFunc {
	return cli.ActionFunc(func(c *cli.Context) error {
		instanceId := c.Args().First()
		zone := c.Args().Get(1)
		projectId := os.Getenv("GCP_PROJECT_ID")
		if instanceId == "" {
			log.Fatal(errors.New("instanceId needs to be specified"))
		}
		if projectId == "" {
			log.Fatal(errors.New("GCP_PROJECT_ID needs to be specified"))
		}
		if zone == "" {
			log.Fatal(errors.New("zone needs to be specified"))
		}
		instance, err := computeClient.GetInstance(projectId, zone, instanceId)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(AsJSON(instance))
		return nil
	})
}

func computeListCmd() *cli.Command {
	return &cli.Command{
		Name:    ActionListCmdName,
		Aliases: []string{"ls"},
		Action:  computeList(context.Background()),
	}
}

func computeList(ctx context.Context) cli.ActionFunc {
	return cli.ActionFunc(func(c *cli.Context) error {
		zone := c.Args().First()
		projectId := os.Getenv("GCP_PROJECT_ID")
		if projectId == "" {
			log.Fatal(errors.New("GCP_PROJECT_ID needs to be specified"))
		}
		if zone == "" {
			log.Fatal(errors.New("zone needs to be specified"))
		}

		instances, err := computeClient.ListInstances(projectId, zone)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(AsJSON(instances))
		return nil
	})
}

func computeResetCmd() *cli.Command {
	return &cli.Command{
		Name:    ActionResetCmdName,
		Aliases: []string{"rs"},
		Action:  computeReset(context.Background()),
	}
}

func computeReset(ctx context.Context) cli.ActionFunc {
	return cli.ActionFunc(func(c *cli.Context) error {
		instanceId := c.Args().First()
		zone := c.Args().Get(1)
		projectId := os.Getenv("GCP_PROJECT_ID")
		if instanceId == "" {
			log.Fatal(errors.New("instanceId needs to be specified"))
		}
		if projectId == "" {
			log.Fatal(errors.New("GCP_PROJECT_ID needs to be specified"))
		}
		if zone == "" {
			log.Fatal(errors.New("zone needs to be specified"))
		}
		operation, err := computeClient.ResetInstance(projectId, zone, instanceId)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(AsJSON(operation))
		return nil
	})
}
