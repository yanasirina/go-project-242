package cli

import (
	"context"
	"fmt"

	"code"

	"github.com/urfave/cli/v3"
)

func RunPathSizeAction(_ context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() != 1 {
		return ErrBadArguments
	}

	filePath := cmd.Args().Get(0)

	size, err := code.GetPathSize(
		filePath,
		cmd.Bool(RecursiveFlag),
		cmd.Bool(HumanFlagName),
		cmd.Bool(ShowAllFilesFlag),
	)
	if err != nil {
		return fmt.Errorf("get path size of %s failed: %w", filePath, err)
	}

	fmt.Println(filePath, size)

	return nil
}
