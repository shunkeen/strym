package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/prosumer"
)

type takeWhile[T any] struct {
	prosumer.Base
	buffer    T
	predicate func(T) bool
}

func TakeWhile[T any](p func(T) bool) prosumer.Naive[T, T] {
	return &takeWhile[T]{
		predicate: p,
		Base:      prosumer.NewBase(),
	}
}

func (m *takeWhile[T]) Await(x T) {
	if !m.predicate(x) {
		m.BaseGoTo = machine.GoToReturn
		return
	}

	m.BaseGoTo = machine.GoToYield
	m.buffer = x
}

func (m *takeWhile[T]) Yield() T {
	m.BaseGoTo = machine.GoToAwait
	return m.buffer
}
