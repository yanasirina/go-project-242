package main

import (
	"log/slog"

	"github.com/yanasirina/go-project-242/internal/app/cli"
)

func main() {
	cmd := cli.NewCLICommand()

	err := cli.RunCMD(cmd)
	if err != nil {
		slog.Error(err.Error())
	}
}
