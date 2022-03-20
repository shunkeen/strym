package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/prosumer"
)

type drop[T any] struct {
	prosumer.Base
	count  int
	buffer T
}

func Drop[T any](n int) prosumer.Naive[T, T] {
	return &drop[T]{
		count: n,
		Base:  prosumer.NewBase(),
	}
}

func (m *drop[T]) Await(x T) {
	if m.count > 0 {
		m.count--
		return
	}

	m.BaseGoTo = machine.GoToYield
	m.buffer = x
}

func (m *drop[T]) Yield() T {
	m.BaseGoTo = machine.GoToAwait
	return m.buffer
}
