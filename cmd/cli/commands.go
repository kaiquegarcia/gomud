package cli

import (
	"context"
	"fmt"
	"gomud/internal/errs"
	"gomud/pkg/logs"
)

type (
	Command   string
	Operation string
)

var (
	CmdStudent Command = "student"
)

func parseCommand(ctx context.Context, args []string) (Command, error) {
	logger := logs.LoggerFromCtx(ctx)
	if len(args) < 2 {
		logger.Error("missing command argument")
		return "", errs.ErrInvalidArgumentSize
	}

	cmd := Command(args[1])
	switch cmd {
	case CmdStudent:
		return cmd, nil
	default:
		logger.Error(fmt.Sprintf("invalid command '%s'", cmd))
		return "", errs.ErrInvalidCommand
	}
}
