package communicator

import "github.com/shunkeen/strym/machine"

type chain[S, T, U, V any] struct {
	goTo1    machine.GoTo
	goTo2    machine.GoTo
	machine1 machine.Machine[S, T, machine.Void]
	machine2 machine.Machine[T, U, V]
}

func Chain[S, T, U, V any](m1 machine.Machine[S, T, machine.Void], m2 machine.Machine[T, U, V]) machine.Machine[S, U, V] {
	return &chain[S, T, U, V]{
		goTo1:    m1.GoTo(),
		goTo2:    m2.GoTo(),
		machine1: m1,
		machine2: m2,
	}
}

func (m *chain[S, T, U, V]) GoTo() machine.GoTo {
	if m.goTo2 == machine.GoToReturn || m.goTo2 == machine.GoToYield {
		return m.goTo2
	}

	if m.goTo1 == machine.GoToAwait && m.goTo2 == machine.GoToAwait {
		return machine.GoToAwait
	}

	return machine.GoToContinue
}

func (m *chain[S, T, U, V]) Await(x S, err error) {
	m.machine1.Await(x, err)
	m.goTo1 = m.machine1.GoTo()
	return
}

func (m *chain[S, T, U, V]) Yield() (U, error) {
	x, err := m.machine2.Yield()
	m.goTo2 = m.machine2.GoTo()
	return x, err
}

func (m *chain[S, T, U, V]) DontWait() {
	m.machine1.DontWait()
	m.goTo1 = m.machine1.GoTo()
}

func (m *chain[S, T, U, V]) Return() (V, error) {
	x, err := m.machine2.Return()
	m.goTo2 = m.machine2.GoTo()
	return x, err
}

func (m *chain[S, T, U, V]) Continue() {
	if m.goTo2 == machine.GoToContinue {
		m.machine2.Continue()
		m.goTo2 = m.machine2.GoTo()
		return
	}

	if m.goTo1 == machine.GoToReturn {
		m.machine2.DontWait()
		m.goTo2 = m.machine2.GoTo()
		return
	}

	if m.goTo1 == machine.GoToContinue {
		m.machine1.Continue()
		m.goTo1 = m.machine1.GoTo()
		return
	}

	m.machine2.Await(m.machine1.Yield())
	m.goTo1 = m.machine1.GoTo()
	m.goTo2 = m.machine2.GoTo()
}

func (m *chain[S, T, U, V]) Defer() {
	defer m.machine1.Defer()
	defer m.machine2.Defer()
}
