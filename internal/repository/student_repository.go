package repository

import (
	"context"
	"fmt"
	"gomud/internal/entity"
	"gomud/internal/errs"
	"gomud/internal/services/db"
	"gomud/pkg/infra"
	"strings"
	"time"
)

type (
	StudentRepository interface {
		List(ctx context.Context, p *db.Pagination) ([]*entity.Student, error)
		GetIDsFromEmails(ctx context.Context, emails []interface{}) (map[string]int, error)
		Fetch(ctx context.Context, ID int) (*entity.Student, error)
		Create(ctx context.Context, e *entity.Student) error
		Update(ctx context.Context, e *entity.Student) error
		Delete(ctx context.Context, ID int) error
		Transaction(ctx context.Context, f func() error) error
	}

	studentRepository struct {
		base
	}
)

func NewStudentRepository(db infra.DB) StudentRepository {
	return &studentRepository{
		base: base{
			db:    db,
			table: "stu_student",
		},
	}
}

func (r *studentRepository) List(ctx context.Context, p *db.Pagination) ([]*entity.Student, error) {
	result, err := r.list(ctx, p)
	if err != nil {
		return nil, err
	}

	return db.ScanResult[entity.Student](result)
}

func (r *studentRepository) GetIDsFromEmails(ctx context.Context, emails []interface{}) (map[string]int, error) {
	sqlKeys := strings.Repeat(", ?", len(emails))[2:]
	sql := fmt.Sprintf("SELECT ID, email FROM %s WHERE email IN (%s) AND status=1", r.base.table, sqlKeys)
	result, err := r.query(ctx, sql, emails...)
	if err != nil {
		return nil, err
	}

	students, err := db.ScanResult[entity.Student](result)
	if err != nil {
		return nil, err
	}

	output := make(map[string]int)
	for _, student := range students {
		output[student.Email] = student.ID
	}

	return output, nil
}

func (r *studentRepository) Fetch(ctx context.Context, ID int) (*entity.Student, error) {
	result, err := r.fetch(ctx, ID)
	if err != nil {
		return nil, err
	}

	list, err := db.ScanResult[entity.Student](result)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, errs.ErrStudentNotFound
	}

	return list[0], nil
}

func (r *studentRepository) Create(ctx context.Context, e *entity.Student) error {
	e.UpdatedAt = time.Now()
	result, err := r.insert(ctx, e)
	if err != nil {
		return err
	}

	e.ID = int(result.InsertId)
	e.CreatedAt = e.UpdatedAt
	return nil
}

func (r *studentRepository) Update(ctx context.Context, e *entity.Student) error {
	e.UpdatedAt = time.Now()
	_, err := r.update(ctx, e, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *studentRepository) Delete(ctx context.Context, ID int) error {
	_, err := r.delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
