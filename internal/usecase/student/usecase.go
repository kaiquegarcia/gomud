package student

import (
	"context"
	"gomud/internal/repository"
)

type (
	StudentUsecases interface {
		Import(context.Context, *ImportInput) error
	}

	studentUsecases struct {
		institutionRepository          repository.InstitutionRepository
		studentRepository              repository.StudentRepository
		planInscriptionRepository      repository.PlanInscriptionRepository
		planInscriptionStoreRepository repository.PlanInscriptionStoreRepository
	}
)

func NewStudentUsecases(
	institutionRepository repository.InstitutionRepository,
	studentRepository repository.StudentRepository,
	planInscriptionRepository repository.PlanInscriptionRepository,
	planInscriptionStoreRepository repository.PlanInscriptionStoreRepository,
) StudentUsecases {
	return &studentUsecases{
		institutionRepository:          institutionRepository,
		studentRepository:              studentRepository,
		planInscriptionRepository:      planInscriptionRepository,
		planInscriptionStoreRepository: planInscriptionStoreRepository,
	}
}
