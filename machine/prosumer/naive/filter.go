package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/prosumer"
)

type filter[T any] struct {
	prosumer.Base
	buffer    T
	predicate func(T) bool
}

func Filter[T any](p func(T) bool) prosumer.Naive[T, T] {
	return &filter[T]{
		predicate: p,
		Base:      prosumer.NewBase(),
	}
}

func (m *filter[T]) Await(x T) {
	if m.predicate(x) {
		m.BaseGoTo = machine.GoToYield
		m.buffer = x
	}
}

func (m *filter[T]) Yield() T {
	m.BaseGoTo = machine.GoToAwait
	return m.buffer
}
