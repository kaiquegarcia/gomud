package infra

import (
	"crypto/tls"

	"github.com/go-mysql-org/go-mysql/client"
	"github.com/go-mysql-org/go-mysql/mysql"
)

var (
	DefaultDatabaseHost     = "127.0.0.1:3306"
	DefaultDatabaseUser     = "root"
	DefaultDatabasePassword = ""
	DefaultDatabaseSchema   = "test"
)

type DB interface {
	Begin() error
	BeginTx(readOnly bool, txIsolation string) error
	Close() error
	Commit() error
	Execute(command string, args ...interface{}) (*mysql.Result, error)
	ExecuteMultiple(query string, perResultCallback client.ExecPerResultCallback) (*mysql.Result, error)
	ExecuteSelectStreaming(command string, result *mysql.Result, perRowCallback client.SelectPerRowCallback, perResultCallback client.SelectPerResultCallback) error
	FieldList(table string, wildcard string) ([]*mysql.Field, error)
	IsInTransaction() bool
	Prepare(query string) (*client.Stmt, error)
	Quit() error
	Rollback() error
	SetAttributes(attributes map[string]string)
	SetAutoCommit() error
	SetCharset(charset string) error
	SetCollation(collation string) error
	SetQueryAttributes(attrs ...mysql.QueryAttribute) error
	SetTLSConfig(config *tls.Config)
	UseSSL(insecureSkipVerify bool)
}
