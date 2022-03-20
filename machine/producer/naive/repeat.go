package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/producer"
)

type repeat[T any] struct {
	producer.Base
	valule T
}

func Repeat[T any](x T) producer.Naive[T] {
	return &repeat[T]{valule: x}
}

func (m *repeat[T]) GoTo() machine.GoTo {
	return machine.GoToYield
}

func (m *repeat[T]) Yield() T {
	return m.valule
}
