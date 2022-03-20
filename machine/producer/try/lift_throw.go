package try

import "github.com/shunkeen/strym/machine/producer"

type liftThrow[T any] struct {
	producer.Throw
	_nil T
}

func LiftThrow[T any](m producer.Throw) producer.Try[T] {
	return &liftThrow[T]{Throw: m}
}

func (m *liftThrow[T]) Yield() (T, error) {
	return m._nil, m.Throw.Yield()
}
