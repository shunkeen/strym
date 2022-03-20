package hermit

import "github.com/shunkeen/strym/machine"

type liftTry[T any] struct {
	Try[T]
	_nil T
	err  error
}

func LiftTry[T any](m Try[T]) machine.Machine[machine.Void, machine.Void, T] {
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

func (m *liftTry[T]) Return() (T, error) {
	if m.err == nil {
		return m.Try.Return()
	}

	err := m.err
	m.err = nil
	return m._nil, err
}
