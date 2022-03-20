package prosumer

import "github.com/shunkeen/strym/machine"

type liftTry[S, T any] struct {
	Try[S, T]
	_nil T
	err  error
}

func LiftTry[S, T any](m Try[S, T]) machine.Machine[S, T, machine.Void] {
	return &liftTry[S, T]{Try: m}
}

func (m *liftTry[S, T]) GoTo() machine.GoTo {
	if m.err != nil {
		return machine.GoToReturn
	}

	return m.Try.GoTo()
}

func (m *liftTry[S, T]) Await(x S, err error) {
	if m.err != nil {
		m.err = err
		return
	}

	m.Try.Await(x)
}

func (m *liftTry[S, T]) Yield() (T, error) {
	if m.err == nil {
		return m.Try.Yield()
	}

	err := m.err
	m.err = nil
	return m._nil, err
}
