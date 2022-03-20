package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/producer"
)

type cycleSlice[T any] struct {
	producer.Base
	index int
	slice []T
}

func CycleSlice[T any](xs []T) producer.Naive[T] {
	return &cycleSlice[T]{slice: xs}
}

func (m *cycleSlice[T]) GoTo() machine.GoTo {
	return machine.GoToYield
}

func (m *cycleSlice[T]) Yield() T {
	i := m.index
	m.index = (m.index + 1) % len(m.slice)
	return m.slice[i]
}
