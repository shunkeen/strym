package communicator

import "github.com/shunkeen/strym/machine"

type sparkWith[S, T1, T2, U any] struct {
	with func(T1, T2) U
	_1   sparkWithMachine[S, T1]
	_2   sparkWithMachine[S, T2]
	_nil U
}

type sparkWithMachine[S, T any] struct {
	goTo  machine.GoTo
	ready bool
	val   T
	err   error
	machine.Machine[S, machine.Void, T]
}

func SparkWith[S, T1, T2, U any](m1 machine.Machine[S, machine.Void, T1], m2 machine.Machine[S, machine.Void, T2], f func(T1, T2) U) machine.Machine[S, machine.Void, U] {
	return &sparkWith[S, T1, T2, U]{
		with: f,
		_1: sparkWithMachine[S, T1]{
			Machine: m1,
			goTo:    m1.GoTo(),
		},
		_2: sparkWithMachine[S, T2]{
			Machine: m2,
			goTo:    m2.GoTo(),
		},
	}
}

func (m *sparkWith[S, T1, T2, U]) GoTo() machine.GoTo {
	if m._1.err != nil || m._2.err != nil || (m._1.ready && m._2.ready) {
		return machine.GoToReturn
	}

	if m._1.goTo == machine.GoToAwait && m._2.goTo == machine.GoToAwait {
		return machine.GoToAwait
	}

	return machine.GoToContinue
}

func (m *sparkWith[S, T1, T2, U]) Await(x S, err error) {
	m._1.Await(x, err)
	m._2.Await(x, err)
	m._1.goTo = m._1.GoTo()
	m._2.goTo = m._2.GoTo()
}

func (m *sparkWith[S, T1, T2, U]) DontWait() {
	m._1.DontWait()
	m._2.DontWait()
	m._1.goTo = m._1.GoTo()
	m._2.goTo = m._2.GoTo()
}

func (m *sparkWith[S, T1, T2, U]) Return() (U, error) {
	if m._1.err != nil {
		err := m._1.err
		m._1.err = nil
		return m._nil, err
	}

	if m._2.err != nil {
		err := m._2.err
		m._2.err = nil
		return m._nil, err
	}

	return m.with(m._1.val, m._2.val), nil
}

func (m *sparkWith[S, T1, T2, U]) Continue() {
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

func (m *sparkWith[S, T1, T2, U]) Yield() (machine.Void, error) {
	return machine.Void{}, nil
}

func (m *sparkWith[S, T1, T2, U]) Defer() {
	defer m._1.Defer()
	defer m._2.Defer()
}
