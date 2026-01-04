package actions

import (
	"code/internal/app/service"
	"code/internal/pkg/errors"
	"context"
	"log/slog"

	"github.com/urfave/cli/v3"
)

func PathSizeAction(ctx context.Context, cmd *cli.Command) error {
	filename := cmd.Args().Get(0)
	if filename == "" {
		return ErrBadArguments
	}

	size, err := service.GetPathSize(filename)
	if err != nil {
		return errors.Wrapf(err, "get path size of %s failed", filename)
	}

	slog.InfoContext(ctx, "Successfully got path size:", slog.Int64("size", size))

	return nil
}
