package repository

import (
	"context"
	"fmt"
	"gomud/internal/services/db"
	"gomud/pkg/infra"
	"gomud/pkg/logs"
	"strings"

	"github.com/go-mysql-org/go-mysql/mysql"
)

type base struct {
	db    infra.DB
	table string
}

func (r *base) query(ctx context.Context, sql string, args ...any) (*mysql.Result, error) {
	logger := logs.LoggerFromCtx(ctx)
	stmt, err := r.db.Prepare(sql)
	if err != nil {
		logger.Error("could not prepare SQL statement", logs.Error(err))
		return nil, err
	}

	result, err := stmt.Execute(args...)
	if err != nil {
		logger.Error("could not execute prepared SQL statement", logs.Error(err))
		return nil, err
	}

	return result, nil
}

func (r *base) list(ctx context.Context, p *db.Pagination) (*mysql.Result, error) {
	p.Sanitize()
	sql := fmt.Sprintf("SELECT * FROM %s ORDER BY registry_date ASC LIMIT ? OFFSET ?", r.table)
	return r.query(ctx, sql, p.Limit, p.Offset)
}

func (r *base) fetch(ctx context.Context, ID int) (*mysql.Result, error) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE ID=? LIMIT 1", r.table)
	return r.query(ctx, sql, ID)
}

func (r *base) insert(ctx context.Context, entity any) (*mysql.Result, error) {
	logger := logs.LoggerFromCtx(ctx)
	fields, err := db.ExtractFields(entity, db.SkipInsert)
	if err != nil {
		logger.Error("could not extract fields from entity", logs.Error(err))
		return nil, err
	}

	sqlKeys := strings.Join(fields.Keys, ", ")
	sqlValues := strings.Repeat(", ?", len(fields.Keys))[2:]
	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", r.table, sqlKeys, sqlValues)
	return r.query(ctx, sql, fields.Values...)
}

func (r *base) update(ctx context.Context, entity any, ID int) (*mysql.Result, error) {
	logger := logs.LoggerFromCtx(ctx)
	fields, err := db.ExtractFields(entity, db.SkipUpdate)
	if err != nil {
		logger.Error("could not extract fields from entity", logs.Error(err))
		return nil, err
	}

	sqlSet := make([]string, len(fields.Keys))
	for i := range len(fields.Keys) {
		sqlSet[i] = fmt.Sprintf("%s=?", fields.Keys[i])
	}

	fields.Values = append(fields.Values, ID)
	sql := fmt.Sprintf("UPDATE %s SET %s WHERE ID=?", r.table, sqlSet)
	return r.query(ctx, sql, fields.Values...)
}

func (r *base) delete(ctx context.Context, ID int) (*mysql.Result, error) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE ID=?", r.table)
	return r.query(ctx, sql, ID)
}

func (r *base) Transaction(ctx context.Context, f func() error) error {
	logger := logs.LoggerFromCtx(ctx)
	if err := r.db.Begin(); err != nil {
		return err
	}

	var functionError error = nil
	defer func() {
		panicErr := recover()
		if panicErr != nil {
			logger.Error("panic recovered during database transaction", logs.Field("error", panicErr))
			if err := r.db.Rollback(); err != nil {
				logger.Error("could not rollback after panic", logs.Error(err))
			}
			return
		}

		if functionError != nil {
			logger.Error("transaction failed", logs.Error(functionError))
			if err := r.db.Rollback(); err != nil {
				logger.Error("could not rollback after transaction failure", logs.Error(err))
			}
			return
		}

		if err := r.db.Commit(); err != nil {
			logger.Error("could not commit transaction", logs.Error(err))
		}
	}()

	functionError = f()
	return functionError
}
