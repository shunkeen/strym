package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/prosumer"
)

type take[T any] struct {
	prosumer.Base
	count  int
	buffer T
}

func Take[T any](n int) prosumer.Naive[T, T] {
	return &take[T]{
		count: n,
		Base:  prosumer.NewBase(),
	}
}

func (m *take[T]) GoTo() machine.GoTo {
	if m.count <= 0 {
		return machine.GoToReturn
	}

	return m.BaseGoTo
}

func (m *take[T]) Await(x T) {
	m.BaseGoTo = machine.GoToYield
	m.buffer = x
}

func (m *take[T]) Yield() T {
	m.BaseGoTo = machine.GoToAwait
	m.count--
	return m.buffer
}
