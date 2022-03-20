package communicator

import "github.com/shunkeen/strym/machine"

type zipWith3[S, T1, T2, T3, U any] struct {
	with func(T1, T2, T3) U
	_1   zipWithMachine[S, T1]
	_2   zipWithMachine[S, T2]
	_3   zipWithMachine[S, T3]
	_nil U
}

func ZipWith3[S, T1, T2, T3, U any](m1 machine.Machine[S, T1, machine.Void], m2 machine.Machine[S, T2, machine.Void], m3 machine.Machine[S, T3, machine.Void], f func(T1, T2, T3) U) machine.Machine[S, U, machine.Void] {
	return &zipWith3[S, T1, T2, T3, U]{
		with: f,
		_1: zipWithMachine[S, T1]{
			Machine: m1,
			goTo:    m1.GoTo(),
		},
		_2: zipWithMachine[S, T2]{
			Machine: m2,
			goTo:    m2.GoTo(),
		},
		_3: zipWithMachine[S, T3]{
			Machine: m3,
			goTo:    m3.GoTo(),
		},
	}
}

func (m *zipWith3[S, T1, T2, T3, U]) GoTo() machine.GoTo {
	if m._1.err != nil || m._2.err != nil || m._3.err != nil || (m._1.ready && m._2.ready && m._3.ready) {
		return machine.GoToYield
	}

	if m._1.goTo == machine.GoToReturn && m._2.goTo == machine.GoToReturn && m._3.goTo == machine.GoToReturn {
		return machine.GoToReturn
	}

	return machine.GoToContinue
}

func (m *zipWith3[S, T1, T2, T3, U]) Yield() (U, error) {
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

	if m._3.err != nil {
		err := m._3.err
		m._3.err = nil
		return m._nil, err
	}

	m._1.ready = false
	m._2.ready = false
	m._3.ready = false
	return m.with(m._1.val, m._2.val, m._3.val), nil
}

func (m *zipWith3[S, T1, T2, T3, U]) Continue() {
	if m._1.goTo == machine.GoToContinue || m._2.goTo == machine.GoToContinue || m._3.goTo == machine.GoToContinue {
		if m._1.goTo == machine.GoToContinue {
			m._1.Continue()
			m._1.goTo = m._1.GoTo()
		}

		if m._2.goTo == machine.GoToContinue {
			m._2.Continue()
			m._2.goTo = m._2.GoTo()
		}

		if m._3.goTo == machine.GoToContinue {
			m._3.Continue()
			m._3.goTo = m._3.GoTo()
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

	if !m._3.ready && m._3.goTo == machine.GoToYield {
		m._3.ready = true
		x3, err := m._3.Yield()
		m._3.goTo = m._3.GoTo()
		if err != nil {
			m._3.err = err
		} else {
			m._3.val = x3
		}
	}
}

func (m *zipWith3[S, T1, T2, T3, U]) Defer() {
	defer m._1.Defer()
	defer m._2.Defer()
	defer m._3.Defer()
}

func (m *zipWith3[S, T1, T2, T3, U]) Await(x S, err error) {}

func (m *zipWith3[S, T1, T2, T3, U]) DontWait() {}

func (m *zipWith3[S, T1, T2, T3, U]) Return() (machine.Void, error) {
	return machine.Void{}, nil
}
