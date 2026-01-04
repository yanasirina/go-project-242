package main

import (
	"log/slog"

	"code/internal/app/cli"
)

func main() {
	cmd := cli.NewCLICommand()

	err := cli.RunCMD(cmd)
	if err != nil {
		slog.Error(err.Error())
	}
}
