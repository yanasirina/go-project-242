package cli

import (
	"code/internal/pkg/errors"
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

type PathSizeCLI struct {
	command cli.Command
}

func NewPathSizeCLI() *PathSizeCLI {
	command := cli.Command{}

	return &PathSizeCLI{command: command}
}

func (p *PathSizeCLI) Run() error {
	if err := p.command.Run(context.Background(), os.Args); err != nil {
		return errors.Wrap(err, "failed to run command")
	}
	return nil
}
