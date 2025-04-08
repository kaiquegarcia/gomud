package cli

import (
	"context"
	"gomud/internal/errs"
	"gomud/pkg/envs"
	"gomud/pkg/infra"
	"gomud/pkg/logs"
	"os"
	"time"
)

func Run() error {
	i := infra.NewInfrastructure(
		envs.Get("LOG_LEVEL", logs.LevelDebug),
	)

	if err := i.ConnectDatabase(); err != nil {
		return err
	}

	defer i.CloseDatabase()

	deps, err := dependencies(i)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(
		i.CreateContext(),
		72*time.Hour,
	)

	defer cancel()

	args := os.Args
	command, err := parseCommand(ctx, args)
	if err != nil {
		return err
	}

	switch command {
	case CmdStudent:
		op, err := parseStudentOperation(ctx, args)
		if err != nil {
			return err
		}

		switch op {
		case OpStudentImport:
			args, err := parseArgs(ctx, args, 3, StudentImportExpectedArgs)
			if err != nil {
				return err
			}

			return runOpStudentImport(ctx, deps, args)
		}
	}

	return errs.ErrCommandNotImplemented
}
