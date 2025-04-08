package cli

import (
	"context"
	"fmt"
	"gomud/internal/errs"
	"gomud/internal/services/enc"
	"gomud/internal/usecase/student"
	"gomud/pkg/logs"
	"gomud/pkg/maps"
	"os"
	"strings"
	"time"
)

var (
	defaultPlanID         = 2
	defaultExpirationDate = time.Now().AddDate(1, 0, 0)
	defaultPassword       = fmt.Sprintf("dacomud%d", time.Now().Year())
)

func runOpStudentImport(
	ctx context.Context,
	deps *deps,
	args arguments,
) error {
	logger := logs.LoggerFromCtx(ctx)
	institution, ok := args["institution"]
	if !ok {
		logger.Error("missing institution argument")
		return errs.ErrMissingInstitutionArgument
	}

	planID := maps.Get(args, "plan_id", defaultPlanID)
	expirationDate := maps.GetTime(args, "expiration_date", defaultExpirationDate, time.DateOnly)
	encPassword, err := enc.VersaPassword(maps.Get(args, "password", defaultPassword))
	if err != nil {
		logger.Error("could not encrypt password into Versa format", logs.Error(err))
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		logger.Error("could not get current working directory")
		return err
	}

	ctx = logs.CtxWithLogger(ctx, logger)
	return deps.studentUsecases.Import(ctx, &student.ImportInput{
		CsvFilePath:       fmt.Sprintf("%s\\%s.csv", wd, strings.ReplaceAll(institution, " ", "_")),
		InstitutionName:   institution,
		PlanID:            planID,
		ExpirationDate:    expirationDate,
		EncryptedPassword: encPassword,
	})
}
