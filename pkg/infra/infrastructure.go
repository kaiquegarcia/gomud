package infra

import (
	"context"
	"gomud/internal/errs"
	"gomud/pkg/envs"
	"gomud/pkg/logs"

	"github.com/go-mysql-org/go-mysql/client"
	"github.com/google/uuid"
)

type (
	Infrastructure interface {
		InstanceID() uuid.UUID
		Database() (DB, error)
		ConnectDatabase() error
		CloseDatabase() error
		Logger() logs.Logger
		CreateContext() context.Context
	}

	infrastructure struct {
		instanceID         uuid.UUID
		minLevelToReport   logs.Level
		databaseConnection DB
		rootLogger         logs.Logger
	}
)

func NewInfrastructure(
	minLevelToReport logs.Level,
) Infrastructure {
	return &infrastructure{
		instanceID:       uuid.New(),
		minLevelToReport: minLevelToReport,
	}
}

func (i *infrastructure) InstanceID() uuid.UUID {
	return i.instanceID
}

func (i *infrastructure) ConnectDatabase() error {
	logger := i.Logger()
	conn, err := client.Connect(
		envs.Get("DB_HOST", DefaultDatabaseHost),
		envs.Get("DB_USER", DefaultDatabaseUser),
		envs.Get("DB_PASS", DefaultDatabasePassword),
		envs.Get("DB_SCHEMA", DefaultDatabaseSchema),
	)

	if err != nil {
		logger.Error("could not connect database", logs.Error(err))
		return errs.ErrCouldNotConnectDatabase
	}

	if err := conn.SetCharset("latin1"); err != nil {
		logger.Warn("could not use charset 'latin1'")
	}

	i.databaseConnection = conn
	return nil
}

func (i *infrastructure) CloseDatabase() error {
	if i.databaseConnection == nil {
		return nil
	}

	if err := i.databaseConnection.Close(); err != nil {
		i.Logger().Error("could not close database", logs.Error(err))
		return err
	}

	return nil
}

func (i *infrastructure) Database() (DB, error) {
	if i.databaseConnection == nil {
		i.Logger().Error("could not return the database connection as it's not connected")
		return nil, errs.ErrDatabaseNotConnected
	}

	return i.databaseConnection, nil
}

func (i *infrastructure) Logger() logs.Logger {
	if i.rootLogger == nil {
		i.rootLogger = logs.NewLogger(i.minLevelToReport)
	}

	return i.rootLogger
}

func (i *infrastructure) CreateContext() context.Context {
	return logs.CtxWithLogger(
		context.Background(),
		i.Logger(),
	)
}
