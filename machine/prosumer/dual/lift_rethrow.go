package dual

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/prosumer"
)

type liftRethrow[T any] struct {
	prosumer.Rethrow
	hasBuffer bool
	buffer    T
	_nil      T
}

func LiftRethrow[T any](m prosumer.Rethrow) prosumer.Dual[T, T] {
	return &liftRethrow[T]{Rethrow: m}
}

func (m *liftRethrow[T]) GoTo() machine.GoTo {
	if m.hasBuffer {
		return machine.GoToYield
	}

	return m.Rethrow.GoTo()
}

func (m *liftRethrow[T]) Await(x T, err error) {
	if err == nil {
		m.hasBuffer = true
		m.buffer = x
		return
	}

	m.Rethrow.Await(err)
}

func (m *liftRethrow[T]) Yield() (T, error) {
	if m.hasBuffer {
		m.hasBuffer = false
		return m.buffer, nil
	}

	return m._nil, m.Rethrow.Yield()
}
