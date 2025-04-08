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
	InstitutionRepository interface {
		List(ctx context.Context, p *db.Pagination) ([]*entity.Institution, error)
		Fetch(ctx context.Context, ID int) (*entity.Institution, error)
		Create(ctx context.Context, e *entity.Institution) error
		Update(ctx context.Context, e *entity.Institution) error
		Delete(ctx context.Context, ID int) error
		Transaction(ctx context.Context, f func() error) error
	}

	institutionRepository struct {
		base
	}
)

func NewInstitutionRepository(db infra.DB) InstitutionRepository {
	return &institutionRepository{
		base: base{
			db:    db,
			table: "stu_student_institution",
		},
	}
}

func (r *institutionRepository) List(ctx context.Context, p *db.Pagination) ([]*entity.Institution, error) {
	result, err := r.list(ctx, p)
	if err != nil {
		return nil, err
	}

	return db.ScanResult[entity.Institution](result)
}

func (r *institutionRepository) Fetch(ctx context.Context, ID int) (*entity.Institution, error) {
	result, err := r.fetch(ctx, ID)
	if err != nil {
		return nil, err
	}

	list, err := db.ScanResult[entity.Institution](result)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, errs.ErrInstitutionNotFound
	}

	return list[0], nil
}

func (r *institutionRepository) Create(ctx context.Context, e *entity.Institution) error {
	result, err := r.insert(ctx, e)
	if err != nil {
		return err
	}

	e.ID = int(result.InsertId)
	e.CreatedAt = time.Now()
	return nil
}

func (r *institutionRepository) Update(ctx context.Context, e *entity.Institution) error {
	_, err := r.update(ctx, e, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *institutionRepository) Delete(ctx context.Context, ID int) error {
	_, err := r.delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
