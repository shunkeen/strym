package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/producer"
)

type replicate[T any] struct {
	producer.Base
	count  int
	valule T
}

func Replicate[T any](n int, x T) producer.Naive[T] {
	return &replicate[T]{
		count:  n,
		valule: x,
	}
}

func (m *replicate[T]) GoTo() machine.GoTo {
	if m.count > 0 {
		return machine.GoToYield
	}

	return machine.GoToReturn
}

func (m *replicate[T]) Yield() T {
	m.count--
	return m.valule
}
