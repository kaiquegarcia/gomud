package cli

import (
	"context"
	"fmt"
	"gomud/internal/errs"
	"gomud/pkg/logs"
)

var (
	OpStudentImport           Operation = "import"
	StudentImportExpectedArgs           = []string{"institution", "plan_id", "expiration_date", "password"}
)

func parseStudentOperation(ctx context.Context, args []string) (Operation, error) {
	logger := logs.LoggerFromCtx(ctx)
	if len(args) < 3 {
		logger.Error("missing operation argument")
		return "", errs.ErrInvalidArgumentSize
	}

	op := Operation(args[2])
	switch op {
	case OpStudentImport:
		return op, nil
	default:
		logger.Error(fmt.Sprintf("invalid operation '%s'", op))
		return "", errs.ErrInvalidOperation
	}
}
