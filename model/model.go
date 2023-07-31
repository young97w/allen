package model

import (
	"errors"
	"github.com/young97w/allen/internal"
	"reflect"
)

type Model struct {
	TableName string
	Fields    map[string]Field
	Columns   map[string]Field
}

type Field struct {
	FieldName string
	ColName   string
	Typ       reflect.Type
	Offset    uintptr
}

type ModelOpt func(model *Model) error

func WithTableName(tableName string) ModelOpt {
	return func(model *Model) error {
		if tableName == "" {
			return errors.New("orm: custom table name can't be empty")
		}
		model.TableName = tableName
		return nil
	}
}

func WithColumnName(fieldName, colName string) ModelOpt {
	return func(model *Model) error {
		if colName == "" {
			return errors.New("orm: custom column name can't be empty")
		}
		f, ok := model.Fields[fieldName]
		if !ok {
			internal.ErrFieldNotExist(fieldName)
		}
		f.ColName = colName
		return nil
	}
}
