package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/consumer"
)

type forEach[T any] struct {
	consumer.Base
	body func(T)
}

func ForEach[T any](f func(T)) consumer.Naive[T, machine.Void] {
	return &forEach[T]{
		body: f,
		Base: consumer.NewBase(),
	}
}

func (m *forEach[T]) Await(x T) {
	m.body(x)
}

func (m *forEach[T]) Return() machine.Void {
	return machine.Void{}
}
