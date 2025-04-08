package db

import (
	"fmt"
	"gomud/internal/errs"
	"reflect"
	"slices"
	"strings"
	"time"
)

type (
	Fields struct {
		Keys   []string
		Values []any
	}

	SkipRule string
)

const (
	SkipInsert SkipRule = "insert"
	SkipUpdate SkipRule = "update"
)

func ExtractFields(entity any, skipRule SkipRule) (*Fields, error) {
	ps := reflect.ValueOf(entity)
	s := ps.Elem()
	if s.Kind() != reflect.Struct {
		return nil, errs.ErrDestMustBeAddressable
	}

	output := &Fields{
		Keys:   make([]string, 0),
		Values: make([]any, 0),
	}
	destType := s.Type()
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

		rules := strings.Split(dbKey, ";")
		if slices.Contains(rules, fmt.Sprintf("skip-%s", skipRule)) {
			continue
		}

		output.Keys = append(output.Keys, rules[0])
		if field.CanInt() {
			output.Values = append(output.Values, field.Int())
		} else if field.CanFloat() {
			output.Values = append(output.Values, field.Float())
		} else if fieldTime, ok := field.Interface().(time.Time); ok {
			layout := fieldTags.Get("layout")
			if layout == "" {
				layout = DefaultTimeLayout
			}

			output.Values = append(output.Values, fieldTime.Format(layout))
		} else {
			output.Values = append(output.Values, field.String())
		}
	}

	return output, nil
}
