package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/prosumer"
)

type scan1[T any] struct {
	prosumer.Base
	isFirst bool
	buffer  T
	scaner  func(T, T) T
}

func Scan1[T any](f func(T, T) T) prosumer.Naive[T, T] {
	return &scan1[T]{
		isFirst: true,
		scaner:  f,
		Base:    prosumer.NewBase(),
	}
}

func (m *scan1[T]) Await(x T) {
	m.BaseGoTo = machine.GoToYield

	if m.isFirst {
		m.isFirst = false
		m.buffer = x
		return
	}

	m.buffer = m.scaner(m.buffer, x)
}

func (m *scan1[T]) Yield() T {
	m.BaseGoTo = machine.GoToAwait
	return m.buffer
}
