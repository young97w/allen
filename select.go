package allen

type Selector[T any] struct {
	table   TableReference
	cols    []Selectable
	where   []Predicate
	groupBy []Selectable
	having  []Predicate
	orderBy []Selectable
	offset  int
	limit   int
	builder
	core
}

func NewSelector[T any]() *Selector[T] {
	return &Selector[T]{}
}

func (s *Selector[T]) Select(cols ...Selectable) *Selector[T] {
	s.cols = cols
	return s
}

func (s *Selector[T]) From(table TableReference) *Selector[T] {
	s.table = table
	return s
}

func (s *Selector[T]) Where(ps ...Predicate) *Selector[T] {
	s.where = ps
	return s
}

func (s *Selector[T]) GroupBy(groupBy ...Selectable) *Selector[T] {
	s.groupBy = groupBy
	return s
}

func (s *Selector[T]) Having(having ...Predicate) *Selector[T] {
	s.having = having
	return s
}

func (s *Selector[T]) OrderBy(orderBy ...Selectable) *Selector[T] {
	s.orderBy = orderBy
	return s
}

func (s *Selector[T]) Limit(limit, offset int) *Selector[T] {
	s.limit, s.offset = limit, offset
	return s
}

func (s *Selector[T]) Build() (*Query, error) {
	m, err := s.r.Get(new(T))
	if err != nil {
		return nil, err
	}
	s.builder.model = m

	s.sb.WriteString("SELECT * FROM ")
	s.sb.WriteString(m.TableName)
	s.sb.WriteByte(';')
	return &Query{SQL: s.sb.String()}, nil
}

type Selectable interface {
	selectable()
}
