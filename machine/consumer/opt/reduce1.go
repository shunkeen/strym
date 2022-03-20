package opt

import "github.com/shunkeen/strym/machine/consumer"

type reduce1[T any] struct {
	consumer.Base
	ok          bool
	accumulator T
	reducer     func(T, T) T
}

func Reduce1[T any](f func(T, T) T) consumer.Opt[T, T] {
	return &reduce1[T]{
		reducer: f,
		Base:    consumer.NewBase(),
	}
}

func (m *reduce1[T]) Await(x T) {
	if !m.ok {
		m.ok = true
		m.accumulator = x
		return
	}

	m.accumulator = m.reducer(m.accumulator, x)
}

func (m *reduce1[T]) Return() (T, bool) {
	return m.accumulator, m.ok
}
