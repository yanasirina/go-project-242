package actions

import (
	"code/internal/app/service"
	"code/internal/pkg/errors"
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

const HumanFlagName = "human"

func PathSizeAction(_ context.Context, cmd *cli.Command) error {
	filePath := cmd.Args().Get(0)
	if filePath == "" {
		return ErrBadArguments
	}

	size, err := service.GetPathSize(filePath)
	if err != nil {
		return errors.Wrapf(err, "get path size of %s failed", filePath)
	}

	humanFlag := cmd.Bool(HumanFlagName)
	formatedSize := formatSize(size, humanFlag)
	fmt.Println(filePath, "-", formatedSize) //nolint:all

	return nil
}
