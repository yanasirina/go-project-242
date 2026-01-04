package cli

import (
	"code/internal/app/actions"
	"code/internal/pkg/errors"
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

type PathSizeCLI struct {
	command *cli.Command
}

func NewPathSizeCLI() *PathSizeCLI {
	cmd := &cli.Command{
		Name:  "Get Path Size",
		Usage: "Command is used to get size of provided file or directory. Command expects path as an argument.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    actions.HumanFlagName,
				Value:   false,
				Usage:   "print human readable format",
				Aliases: []string{"H"},
			},
			&cli.BoolFlag{
				Name:    actions.ShowAllFilesFlag,
				Value:   false,
				Usage:   "show all files (including hidden files)",
				Aliases: []string{"a"},
			},
		},
		Action: actions.PathSizeAction,
	}

	return &PathSizeCLI{command: cmd}
}

func (p *PathSizeCLI) Run() error {
	if err := p.command.Run(context.Background(), os.Args); err != nil {
		return errors.Wrap(err, "failed to run command")
	}

	return nil
}
