package main

import (
	"gomud/cmd/cli"
	"gomud/internal/errs"
	"gomud/pkg/envs"
	"os"
)

var (
	Exit = os.Exit
)

func main() {
	envs.Load()
	if err := cli.Run(); err != nil {
		Exit(errs.ExitCode(err))
		return
	}

	Exit(0)
}
