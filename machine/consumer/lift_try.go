package consumer

import "github.com/shunkeen/strym/machine"

type liftTry[S, T any] struct {
	Try[S, T]
	_nil T
	err  error
}

func LiftTry[S, T any](m Try[S, T]) machine.Machine[S, machine.Void, T] {
	return &liftTry[S, T]{Try: m}
}

func (m *liftTry[S, T]) GoTo() machine.GoTo {
	if m.err != nil {
		return machine.GoToReturn
	}

	return m.Try.GoTo()
}

func (m *liftTry[S, T]) Await(x S, err error) {
	if err != nil {
		m.err = err
		return
	}

	m.Try.Await(x)
}

func (m *liftTry[S, T]) Return() (T, error) {
	if m.err == nil {
		return m.Try.Return()
	}

	err := m.err
	m.err = nil
	return m._nil, err
}
