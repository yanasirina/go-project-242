package actions

import (
	"code/internal/app/service"
	"code/internal/pkg/errors"
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

const (
	HumanFlagName    = "human"
	ShowAllFilesFlag = "all"
)

func PathSizeAction(_ context.Context, cmd *cli.Command) error {
	includeHidden := cmd.Bool(ShowAllFilesFlag)
	humanize := cmd.Bool(HumanFlagName)

	filePath := cmd.Args().Get(0)
	if filePath == "" {
		return ErrBadArguments
	}

	size, err := service.GetPathSize(filePath, includeHidden)
	if err != nil {
		return errors.Wrapf(err, "get path size of %s failed", filePath)
	}

	formatedSize := formatSize(size, humanize)
	fmt.Println(filePath, "-", formatedSize) //nolint:all

	return nil
}
