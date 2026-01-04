package main

import (
	"code/internal/app/cli"
	"log/slog"
)

func main() {
	cmd := cli.NewCLICommand()

	err := cli.RunCMD(cmd)
	if err != nil {
		slog.Error(err.Error())
	}
}
