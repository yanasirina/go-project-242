package cli

import (
	"code/internal/pkg/errors"
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

const (
	HumanFlagName    = "human"
	ShowAllFilesFlag = "all"
)

func NewCLICommand() *cli.Command {
	cmd := &cli.Command{
		Name:  "Get Path Size",
		Usage: "Command is used to get size of provided file or directory. Command expects path as an argument.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    HumanFlagName,
				Value:   false,
				Usage:   "human-readable sizes (auto-select unit)",
				Aliases: []string{"H"},
			},
			&cli.BoolFlag{
				Name:    ShowAllFilesFlag,
				Value:   false,
				Usage:   "include hidden files and directories",
				Aliases: []string{"a"},
			},
		},
		Action: RunPathSizeAction,
	}

	return cmd
}

func RunCMD(c *cli.Command) error {
	if err := c.Run(context.Background(), os.Args); err != nil {
		return errors.Wrap(err, "failed to run command")
	}

	return nil
}
