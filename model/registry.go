package model

import (
	"errors"
	"reflect"
	"sync"
	"unicode"
)

// 我们支持的全部标签上的 key 都放在这里
// 方便用户查找，和我们后期维护
const (
	TagKeyColumn = "column"
)

type Registry struct {
	models sync.Map //store all registered models info
}

func (r *Registry) Get(val any) (*Model, error) {
	typ := reflect.TypeOf(val)
	m, ok := r.models.Load(typ)
	if ok {
		return m.(*Model), nil
	}

	m, err := r.ParseModel(val)
	if err != nil {
		return nil, err
	}

	r.models.Store(typ, m)
	return m.(*Model), nil
}

func (r *Registry) Register(val any, opts ...ModelOpt) (*Model, error) {
	m, err := r.ParseModel(val)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		err = opt(m)
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (r *Registry) ParseModel(obj any) (*Model, error) {
	typ := reflect.TypeOf(obj)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("orm: only support struct or pointer of a struct")
	}
	if typ.NumField() == 0 {
		return nil, errors.New("orm: struct has no field")
	}
	fields := make(map[string]Field, typ.NumField())
	columns := make(map[string]Field, typ.NumField())

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		colName := f.Tag.Get(TagKeyColumn)
		if colName == "" {
			colName = underScoreName(f.Name)
		}
		field := Field{
			FieldName: f.Name,
			ColName:   colName,
			Typ:       f.Type,
			Offset:    f.Offset,
		}

		fields[f.Name] = field
		columns[colName] = field
	}

	return &Model{
		TableName: underScoreName(typ.Name()),
		Fields:    fields,
		Columns:   columns,
	}, nil
}

// underScoreName return underscore of a name
// example: UserName -> user_name
func underScoreName(s string) string {
	var buf []byte
	for i, v := range s {
		if unicode.IsUpper(v) {
			if i != 0 {
				buf = append(buf, '_')
			}
			buf = append(buf, byte(unicode.ToLower(v)))
		} else {
			buf = append(buf, byte(v))
		}

	}
	return string(buf)
}

type Register interface {
	Get(val any) (*Model, error)
	Register(val any, opt ...ModelOpt) (*Model, error)
}
