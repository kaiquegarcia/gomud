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
	PlanInscriptionStoreRepository interface {
		List(ctx context.Context, p *db.Pagination) ([]*entity.PlanInscriptionStore, error)
		Fetch(ctx context.Context, ID int) (*entity.PlanInscriptionStore, error)
		Create(ctx context.Context, e *entity.PlanInscriptionStore) error
		Update(ctx context.Context, e *entity.PlanInscriptionStore) error
		Delete(ctx context.Context, ID int) error
		Transaction(ctx context.Context, f func() error) error
	}

	planInscriptionStoreRepository struct {
		base
	}
)

func NewPlanInscriptionStoreRepository(db infra.DB) PlanInscriptionStoreRepository {
	return &planInscriptionStoreRepository{
		base: base{
			db:    db,
			table: "plan_planinscriptionstore",
		},
	}
}

func (r *planInscriptionStoreRepository) List(ctx context.Context, p *db.Pagination) ([]*entity.PlanInscriptionStore, error) {
	result, err := r.list(ctx, p)
	if err != nil {
		return nil, err
	}

	return db.ScanResult[entity.PlanInscriptionStore](result)
}

func (r *planInscriptionStoreRepository) Fetch(ctx context.Context, ID int) (*entity.PlanInscriptionStore, error) {
	result, err := r.fetch(ctx, ID)
	if err != nil {
		return nil, err
	}

	list, err := db.ScanResult[entity.PlanInscriptionStore](result)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, errs.ErrPlanInscriptionStoreNotFound
	}

	return list[0], nil
}

func (r *planInscriptionStoreRepository) Create(ctx context.Context, e *entity.PlanInscriptionStore) error {
	e.UpdatedAt = time.Now()
	result, err := r.insert(ctx, e)
	if err != nil {
		return err
	}

	e.ID = int(result.InsertId)
	e.CreatedAt = e.UpdatedAt
	return nil
}

func (r *planInscriptionStoreRepository) Update(ctx context.Context, e *entity.PlanInscriptionStore) error {
	e.UpdatedAt = time.Now()
	_, err := r.update(ctx, e, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *planInscriptionStoreRepository) Delete(ctx context.Context, ID int) error {
	_, err := r.delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
