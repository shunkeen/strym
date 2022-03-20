package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/prosumer"
)

type naiveMap[S, T any] struct {
	prosumer.Base
	buffer      S
	transformer func(S) T
}

func Map[S, T any](f func(S) T) prosumer.Naive[S, T] {
	return &naiveMap[S, T]{
		transformer: f,
		Base:        prosumer.NewBase(),
	}
}

func (m *naiveMap[S, T]) Await(x S) {
	m.BaseGoTo = machine.GoToYield
	m.buffer = x
}

func (m *naiveMap[S, T]) Yield() T {
	m.BaseGoTo = machine.GoToAwait
	return m.transformer(m.buffer)
}
