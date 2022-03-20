package communicator

import "github.com/shunkeen/strym/machine"

type zipWith[S, T1, T2, U any] struct {
	with func(T1, T2) U
	_1   zipWithMachine[S, T1]
	_2   zipWithMachine[S, T2]
	_nil U
}

type zipWithMachine[S, T any] struct {
	goTo  machine.GoTo
	ready bool
	val   T
	err   error
	machine.Machine[S, T, machine.Void]
}

func ZipWith[S, T1, T2, U any](m1 machine.Machine[S, T1, machine.Void], m2 machine.Machine[S, T2, machine.Void], f func(T1, T2) U) machine.Machine[S, U, machine.Void] {
	return &zipWith[S, T1, T2, U]{
		with: f,
		_1: zipWithMachine[S, T1]{
			Machine: m1,
			goTo:    m1.GoTo(),
		},
		_2: zipWithMachine[S, T2]{
			Machine: m2,
			goTo:    m2.GoTo(),
		},
	}
}

func (m *zipWith[S, T1, T2, U]) GoTo() machine.GoTo {
	if m._1.err != nil || m._2.err != nil || (m._1.ready && m._2.ready) {
		return machine.GoToYield
	}

	if m._1.goTo == machine.GoToReturn && m._2.goTo == machine.GoToReturn {
		return machine.GoToReturn
	}

	return machine.GoToContinue
}

func (m *zipWith[S, T1, T2, U]) Yield() (U, error) {
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

	m._1.ready = false
	m._2.ready = false
	return m.with(m._1.val, m._2.val), nil
}

func (m *zipWith[S, T1, T2, U]) Continue() {
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

	if !m._1.ready && m._1.goTo == machine.GoToYield {
		m._1.ready = true
		x1, err := m._1.Yield()
		m._1.goTo = m._1.GoTo()
		if err != nil {
			m._1.err = err
		} else {
			m._1.val = x1
		}
	}

	if !m._2.ready && m._2.goTo == machine.GoToYield {
		m._2.ready = true
		x2, err := m._2.Yield()
		m._2.goTo = m._2.GoTo()
		if err != nil {
			m._2.err = err
		} else {
			m._2.val = x2
		}
	}
}

func (m *zipWith[S, T1, T2, U]) Defer() {
	defer m._1.Defer()
	defer m._2.Defer()
}

func (m *zipWith[S, T1, T2, U]) Await(x S, err error) {}

func (m *zipWith[S, T1, T2, U]) DontWait() {}

func (m *zipWith[S, T1, T2, U]) Return() (machine.Void, error) {
	return machine.Void{}, nil
}
