package cli

import (
	"context"
	"fmt"

	"code"
	"code/internal/pkg/errors"

	"github.com/urfave/cli/v3"
)

func RunPathSizeAction(_ context.Context, cmd *cli.Command) error {
	filePath := cmd.Args().Get(0)
	if filePath == "" {
		return ErrBadArguments
	}

	size, err := code.GetPathSize(
		filePath,
		cmd.Bool(RecursiveFlag),
		cmd.Bool(HumanFlagName),
		cmd.Bool(ShowAllFilesFlag),
	)
	if err != nil {
		return errors.Wrapf(err, "get path size of %s failed", filePath)
	}

	fmt.Println(filePath, "-", size)

	return nil
}
