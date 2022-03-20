package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/producer"
)

type fromSlice[T any] struct {
	producer.Base
	index int
	slice []T
}

func FromSlice[T any](xs []T) producer.Naive[T] {
	return &fromSlice[T]{slice: xs}
}

func (m *fromSlice[T]) GoTo() machine.GoTo {
	if m.index < len(m.slice) {
		return machine.GoToYield
	}

	return machine.GoToReturn
}

func (m *fromSlice[T]) Yield() T {
	i := m.index
	m.index++
	return m.slice[i]
}
