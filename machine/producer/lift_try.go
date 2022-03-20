package producer

import "github.com/shunkeen/strym/machine"

type liftTry[T any] struct {
	Try[T]
	_nil T
	err  error
}

func LiftTry[T any](m Try[T]) machine.Machine[machine.Void, T, machine.Void] {
	return &liftTry[T]{Try: m}
}

func (m *liftTry[T]) GoTo() machine.GoTo {
	if m.err != nil {
		return machine.GoToYield
	}

	return m.Try.GoTo()
}

func (m *liftTry[T]) Await(_ machine.Void, err error) {
	m.err = err
}

func (m *liftTry[T]) Yield() (T, error) {
	if m.err == nil {
		return m.Try.Yield()
	}

	err := m.err
	m.err = nil
	return m._nil, err
}
