package communicator

import "github.com/shunkeen/strym/machine"

type redirectWith[R, S, T, U any] struct {
	with func(S, T) U
	_1   redirectWithMachine[R, S]
	_2   redirectWithMachine[error, T]
	_nil U
}

type redirectWithMachine[S, T any] struct {
	machine.Machine[S, machine.Void, T]
	goTo  machine.GoTo
	ready bool
	val   T
	err   error
}

func RedirectWith[R, S, T, U any](m1 machine.Machine[R, machine.Void, S], m2 machine.Machine[error, machine.Void, T], f func(S, T) U) machine.Machine[R, machine.Void, U] {
	return &redirectWith[R, S, T, U]{
		_1: redirectWithMachine[R, S]{
			Machine: m1,
			goTo:    m1.GoTo(),
		},
		_2: redirectWithMachine[error, T]{
			Machine: m2,
			goTo:    m2.GoTo(),
		},
		with: f,
	}
}

func (m *redirectWith[R, S, T, U]) GoTo() machine.GoTo {
	if m._1.err != nil || m._2.err != nil || (m._1.ready && m._2.ready) {
		return machine.GoToReturn
	}

	if m._1.goTo == machine.GoToAwait && m._2.goTo == machine.GoToAwait {
		return machine.GoToAwait
	}

	return machine.GoToContinue
}

func (m *redirectWith[R, S, T, U]) Continue() {
	if m._1.goTo == machine.GoToContinue || m._2.goTo == machine.GoToContinue {
		if m._1.goTo == machine.GoToContinue {
			m._1.Continue()
			m._1.goTo = m._1.GoTo()
		}

		if m._2.goTo == machine.GoToContinue {
			m._2.Continue()
			m._2.goTo = m._2.GoTo()
		}

		return
	}

	if m._1.goTo == machine.GoToReturn {
		x, err := m._1.Return()
		m._1.goTo = m._1.GoTo()
		if err != nil {
			m._1.err = err
		} else {
			m._1.ready = true
			m._1.val = x
		}
	}

	if m._2.goTo == machine.GoToReturn {
		y, err := m._2.Return()
		m._2.goTo = m._2.GoTo()
		if err != nil {
			m._2.err = err
		} else {
			m._2.ready = true
			m._2.val = y
		}
	}
}

func (m *redirectWith[R, S, T, U]) Await(x R, err error) {
	if err != nil {
		m._2.Await(err, nil)
		m._2.goTo = m._2.GoTo()
	} else {
		m._1.Await(x, nil)
		m._1.goTo = m._1.GoTo()
	}
}

func (m *redirectWith[R, S, T, U]) DontWait() {
	m._1.DontWait()
	m._2.DontWait()
	m._1.goTo = m._1.GoTo()
	m._2.goTo = m._2.GoTo()
}

func (m *redirectWith[R, S, T, U]) Return() (U, error) {
	if m._1.err != nil {
		return m._nil, m._1.err
	}

	if m._2.err != nil {
		return m._nil, m._2.err
	}

	return m.with(m._1.val, m._2.val), nil
}

func (m *redirectWith[R, S, T, U]) Yield() (machine.Void, error) {
	return machine.Void{}, nil
}

func (m *redirectWith[R, S, T, U]) Defer() {
	defer m._1.Defer()
	defer m._2.Defer()
}
