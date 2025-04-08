package cli

import (
	"gomud/internal/repository"
	"gomud/internal/usecase/student"
	"gomud/pkg/infra"
)

type deps struct {
	studentUsecases student.StudentUsecases
}

func dependencies(i infra.Infrastructure) (*deps, error) {
	db, err := i.Database()
	if err != nil {
		return nil, err
	}

	institutionRepository := repository.NewInstitutionRepository(db)
	studentRepository := repository.NewStudentRepository(db)
	planInscriptionRepository := repository.NewPlanInscriptionRepository(db)
	planInscriptionStoreRepository := repository.NewPlanInscriptionStoreRepository(db)

	studentUsecases := student.NewStudentUsecases(
		institutionRepository,
		studentRepository,
		planInscriptionRepository,
		planInscriptionStoreRepository,
	)

	return &deps{
		studentUsecases: studentUsecases,
	}, nil
}
