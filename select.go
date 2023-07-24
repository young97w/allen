package allen

import (
	"reflect"
	"strings"
)

type Selector[T any] struct {
	table string
	sb    strings.Builder
}

func NewSelector[T any]() *Selector[T] {
	return &Selector[T]{}
}

func (s *Selector[T]) From(table string) *Selector[T] {
	s.table = table
	return s
}

func (s *Selector[T]) Builder() (*Query, error) {
	if s.table == "" {
		var t T
		s.table = reflect.TypeOf(t).Name()
	}
	s.sb.WriteString("SELECT * FROM ")
	s.sb.WriteString(s.table)
	s.sb.WriteByte(';')
	return &Query{SQL: s.sb.String()}, nil
}
