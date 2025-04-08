package repository

import (
	"context"
	"gomud/internal/entity"
	"gomud/internal/errs"
	"gomud/internal/services/db"
	"gomud/pkg/infra"
	"time"
)

type (
	PlanInscriptionRepository interface {
		List(ctx context.Context, p *db.Pagination) ([]*entity.PlanInscription, error)
		Fetch(ctx context.Context, ID int) (*entity.PlanInscription, error)
		Create(ctx context.Context, e *entity.PlanInscription) error
		Update(ctx context.Context, e *entity.PlanInscription) error
		Delete(ctx context.Context, ID int) error
		Transaction(ctx context.Context, f func() error) error
	}

	planInscriptionRepository struct {
		base
	}
)

func NewPlanInscriptionRepository(db infra.DB) PlanInscriptionRepository {
	return &planInscriptionRepository{
		base: base{
			db:    db,
			table: "plan_planinscription",
		},
	}
}

func (r *planInscriptionRepository) List(ctx context.Context, p *db.Pagination) ([]*entity.PlanInscription, error) {
	result, err := r.list(ctx, p)
	if err != nil {
		return nil, err
	}

	return db.ScanResult[entity.PlanInscription](result)
}

func (r *planInscriptionRepository) Fetch(ctx context.Context, ID int) (*entity.PlanInscription, error) {
	result, err := r.fetch(ctx, ID)
	if err != nil {
		return nil, err
	}

	list, err := db.ScanResult[entity.PlanInscription](result)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, errs.ErrPlanInscriptionNotFound
	}

	return list[0], nil
}

func (r *planInscriptionRepository) Create(ctx context.Context, e *entity.PlanInscription) error {
	e.UpdatedAt = time.Now()
	result, err := r.insert(ctx, e)
	if err != nil {
		return err
	}

	e.ID = int(result.InsertId)
	e.CreatedAt = e.UpdatedAt
	return nil
}

func (r *planInscriptionRepository) Update(ctx context.Context, e *entity.PlanInscription) error {
	e.UpdatedAt = time.Now()
	_, err := r.update(ctx, e, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *planInscriptionRepository) Delete(ctx context.Context, ID int) error {
	_, err := r.delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
