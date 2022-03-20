package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/producer"
)

type iterate[T any] struct {
	producer.Base
	current T
	next    func(T) T
}

func Iterate[T any](x T, f func(T) T) producer.Naive[T] {
	return &iterate[T]{
		current: x,
		next:    f,
	}
}

func (m *iterate[T]) GoTo() machine.GoTo {
	return machine.GoToYield
}

func (m *iterate[T]) Yield() T {
	x := m.current
	m.current = m.next(m.current)
	return x
}
