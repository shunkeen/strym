package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/producer"
)

type once[T any] struct {
	producer.Base
	goTo machine.GoTo
	val  T
}

func Once[T any](x T) producer.Naive[T] {
	return &once[T]{
		goTo: machine.GoToYield,
		val:  x,
	}
}

func (m *once[T]) GoTo() machine.GoTo {
	return m.goTo
}

func (m *once[T]) Yield() T {
	m.goTo = machine.GoToReturn
	return m.val
}
