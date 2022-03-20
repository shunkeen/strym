package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/prosumer"
)

type scan[S, T any] struct {
	prosumer.Base
	buffer T
	scaner func(T, S) T
}

func Scan[S, T any](x T, f func(T, S) T) prosumer.Naive[S, T] {
	m := prosumer.NewBase()
	m.BaseGoTo = machine.GoToYield

	return &scan[S, T]{
		buffer: x,
		scaner: f,
		Base:   m,
	}
}

func (m *scan[S, T]) Await(x S) {
	m.BaseGoTo = machine.GoToYield
	m.buffer = m.scaner(m.buffer, x)
}

func (m *scan[S, T]) Yield() T {
	m.BaseGoTo = machine.GoToAwait
	return m.buffer
}
