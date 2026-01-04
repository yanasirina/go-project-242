package main

import (
	"code/internal/app/cli"
	"log/slog"
)

func main() {
	pathSizeCLI := cli.NewPathSizeCLI()

	err := pathSizeCLI.Run()
	if err != nil {
		slog.Error(err.Error())
	}
}
