package db

import (
	"gomud/internal/errs"
	"reflect"
	"strings"
	"time"

	"github.com/go-mysql-org/go-mysql/mysql"
)

var (
	DefaultTimeLayout = time.RFC3339
)

func scanRow(
	resultSet *mysql.Resultset,
	row int,
	dest any,
) error {
	ps := reflect.ValueOf(dest)
	s := ps.Elem()
	if s.Kind() != reflect.Struct {
		return errs.ErrDestMustBeAddressable
	}

	destType := s.Type()
	dbFields := resultSet.FieldNames
	for i := range s.NumField() {
		field := s.Field(i)
		if !field.IsValid() || !field.CanSet() {
			continue
		}

		fieldTags := destType.Field(i).Tag
		dbKey := fieldTags.Get("db")
		if dbKey == "" {
			continue
		}

		dbKey = strings.Split(dbKey, ";")[0]
		column, ok := dbFields[dbKey]
		if !ok {
			continue
		}

		v, err := resultSet.GetValue(row, column)
		if err != nil {
			return err
		}

		dbValue := reflect.ValueOf(v)
		if dbValue.CanInt() {
			field.SetInt(dbValue.Int())
		} else if dbValue.CanFloat() {
			field.SetFloat(dbValue.Float())
		} else if _, ok := field.Interface().(time.Time); ok {
			layout := fieldTags.Get("layout")
			if layout == "" {
				layout = DefaultTimeLayout
			}

			parsedDbValue, err := time.Parse(layout, string(dbValue.Bytes()))
			if err != nil {
				return err
			}

			field.Set(reflect.ValueOf(parsedDbValue))
		} else {
			field.SetString(string(dbValue.Bytes()))
		}
	}

	return nil
}
