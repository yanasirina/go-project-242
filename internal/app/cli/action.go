package cli

import (
	"code/internal/app/handler"
	"code/internal/pkg/errors"
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func RunPathSizeAction(_ context.Context, cmd *cli.Command) error {
	filePath := cmd.Args().Get(0)
	if filePath == "" {
		return ErrBadArguments
	}

	pathSizeHandler := handler.PathSizeHandler{
		Arguments: handler.CommandArguments{
			Path: filePath,
		},
		Flags: handler.CommandFlags{
			HumanizeSize:    cmd.Bool(HumanFlagName),
			ShowHiddenFiles: cmd.Bool(ShowAllFilesFlag),
		},
	}

	size, err := pathSizeHandler.GetFormatedSize()
	if err != nil {
		return errors.Wrapf(err, "get path size of %s failed", filePath)
	}

	fmt.Println(filePath, "-", size) //nolint:all

	return nil
}
