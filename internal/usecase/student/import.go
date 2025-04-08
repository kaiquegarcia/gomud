package student

import (
	"context"
	"encoding/csv"
	"gomud/internal/entity"
	"gomud/pkg/envs"
	"gomud/pkg/logs"
	"io"
	"os"
	"time"
)

const DefaultAuthorID = 0

type ImportInput struct {
	CsvFilePath       string
	InstitutionName   string
	EncryptedPassword string
	PlanID            int
	ExpirationDate    time.Time
}

func (u *studentUsecases) Import(ctx context.Context, input *ImportInput) error {
	logger := logs.LoggerFromCtx(ctx)
	studentsToImport, err := u.extractStudentsFromCsv(ctx, input.CsvFilePath, input.EncryptedPassword)
	if err != nil {
		return err
	}

	emails := make([]interface{}, 0)
	for _, student := range studentsToImport {
		emails = append(emails, student.Email)
	}

	studentsToUpdate, err := u.studentRepository.GetIDsFromEmails(ctx, emails)
	if err != nil {
		return err
	}

	err = u.studentRepository.Transaction(ctx, func() error {
		authorID := envs.Get("AUTHOR_ID", DefaultAuthorID)
		institution := &entity.Institution{
			AuthorID:    authorID,
			Name:        input.InstitutionName,
			Description: "",
		}
		if err := u.institutionRepository.Create(ctx, institution); err != nil {
			return err
		}

		for _, student := range studentsToImport {
			student.InstitutionID = institution.ID
			if ID, ok := studentsToUpdate[student.Email]; ok {
				student.ID = ID
				err := u.studentRepository.Update(ctx, student)
				if err != nil {
					return err
				}
			} else {
				err := u.studentRepository.Create(ctx, student)
				if err != nil {
					return err
				}
			}

			if err := u.planInscriptionRepository.Create(ctx, &entity.PlanInscription{
				AuthorID:     authorID,
				PlanID:       input.PlanID,
				StudentID:    student.ID,
				VoucherID:    0,
				ChargeID:     0,
				PlanCost:     25.9,
				PlanDuration: "+1 year",
				Discount:     25.9,
				DiscountTXT:  "R$ 25,90",
				TotalCost:    0.0,
				StartedAt:    time.Now(),
				FinishedAt:   input.ExpirationDate,
				Status:       1,
			}); err != nil {
				return err
			}

			if err := u.planInscriptionStoreRepository.Create(ctx, &entity.PlanInscriptionStore{
				AuthorID:      authorID,
				PlanID:        input.PlanID,
				Email:         student.Email,
				EmailContent:  "",
				DurationPrize: "years",
				Duration:      1,
				Charging:      0,
				Status:        1,
			}); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	logger.Info("import finished successfully")
	return nil
}

func (u *studentUsecases) extractStudentsFromCsv(
	ctx context.Context,
	csvFilePath string,
	encryptedPassword string,
) ([]*entity.Student, error) {
	logger := logs.LoggerFromCtx(ctx)
	stream, err := os.Open(csvFilePath)
	if err != nil {
		logger.Error("could not open CSV file", logs.Error(err))
		return nil, err
	}

	defer stream.Close()

	csvReader := csv.NewReader(stream)
	studentsToImport := make([]*entity.Student, 0)
	headers := make([]string, 0)
	line := 0
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			logger.Debug("end of file reached")
			break
		} else if err != nil {
			logger.Error("could not read CSV row", logs.Field("line", line))
			return nil, err
		}

		if line == 0 {
			headers = row
			line++
			continue
		}

		m := make(map[string]string)
		for column := range len(headers) {
			m[headers[column]] = row[column]
		}

		studentsToImport = append(studentsToImport, &entity.Student{
			Name:     m["Nome"],
			Email:    m["E-mail"],
			Login:    m["E-mail"],
			Password: encryptedPassword,
			Iam:      1,
			Status:   1,
		})
		line++
	}

	logger.Debug(
		"students extracted successfully",
		logs.Field("csv_lines", line),
	)
	return studentsToImport, nil
}
