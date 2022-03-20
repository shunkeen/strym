package dual

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/prosumer"
)

type dualId[T any] struct {
	prosumer.Base
	val T
	err error
}

func DualId[T any]() prosumer.Dual[T, T] {
	return &dualId[T]{Base: prosumer.NewBase()}
}

func (m *dualId[T]) Await(x T, err error) {
	m.BaseGoTo = machine.GoToYield
	m.val, m.err = x, err
}

func (m *dualId[T]) Yield() (T, error) {
	m.BaseGoTo = machine.GoToAwait
	return m.val, m.err
}
