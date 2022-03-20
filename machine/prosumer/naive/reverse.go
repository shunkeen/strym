package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/prosumer"
)

type reverse[T any] struct {
	prosumer.Base
	index  int
	buffer []T
}

func Reverse[T any]() prosumer.Naive[T, T] {
	m := prosumer.NewBase()
	return &reverse[T]{Base: m}
}

func (m *reverse[T]) Await(x T) {
	m.buffer = append(m.buffer, x)
}

func (m *reverse[T]) DontWait() {
	if len(m.buffer) == 0 {
		m.BaseGoTo = machine.GoToReturn
	} else {
		m.BaseGoTo = machine.GoToYield
		m.index = len(m.buffer) - 1
	}
}

func (m *reverse[T]) Yield() T {
	i := m.index
	m.index--
	if m.index < 0 {
		m.BaseGoTo = machine.GoToReturn
	}
	return m.buffer[i]
}
