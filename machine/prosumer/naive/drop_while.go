package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/prosumer"
)

type dropWhile[T any] struct {
	prosumer.Base
	passthrough bool
	buffer      T
	predicate   func(T) bool
}

func DropWhile[T any](p func(T) bool) prosumer.Naive[T, T] {
	return &dropWhile[T]{
		predicate: p,
		Base:      prosumer.NewBase(),
	}
}

func (m *dropWhile[T]) Await(x T) {
	if m.passthrough || !m.predicate(x) {
		m.passthrough = true
		m.BaseGoTo = machine.GoToYield
		m.buffer = x
	}
}

func (m *dropWhile[T]) Yield() T {
	m.BaseGoTo = machine.GoToAwait
	return m.buffer
}
