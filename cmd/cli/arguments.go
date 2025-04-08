package cli

import (
	"context"
	"gomud/internal/errs"
	"gomud/pkg/logs"
	"slices"
	"strings"
)

type arguments map[string]string

func parseArgs(
	ctx context.Context,
	args []string,
	initialIndex int,
	allowedArguments []string,
) (arguments, error) {
	opArgs := make(arguments)
	if len(args) < (initialIndex + 1) {
		return opArgs, nil
	}

	for _, a := range args[initialIndex:] {
		a = strings.Replace(a, "--", "", 1)
		argVal := strings.Split(a, "=")
		a = argVal[0]
		if slices.Contains(allowedArguments, a) {
			opArgs[a] = strings.Join(argVal[1:], "=")
		} else {
			logs.LoggerFromCtx(ctx).
				Error("unexpected argument", logs.Field("argument", a))
			return nil, errs.ErrInvalidArgument
		}
	}

	return opArgs, nil
}
