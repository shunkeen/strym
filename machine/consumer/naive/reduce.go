package naive

import "github.com/shunkeen/strym/machine/consumer"

type reduce[S, T any] struct {
	consumer.Base
	accumulator T
	reducer     func(T, S) T
}

func Reduce[S, T any](x T, f func(T, S) T) consumer.Naive[S, T] {
	return &reduce[S, T]{
		accumulator: x,
		reducer:     f,
		Base:        consumer.NewBase(),
	}
}

func (m *reduce[S, T]) Await(x S) {
	m.accumulator = m.reducer(m.accumulator, x)
}

func (m *reduce[S, T]) Return() T {
	return m.accumulator
}
