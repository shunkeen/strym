package communicator

import "github.com/shunkeen/strym/machine"

type concatMap[S, T, U any] struct {
	outerGoTo   machine.GoTo
	innerGoTo   machine.GoTo
	outer       machine.Machine[S, T, machine.Void]
	inner       machine.Machine[machine.Void, U, machine.Void]
	_nil        U
	err         error
	transformer func(T) machine.Machine[machine.Void, U, machine.Void]
	deferStack  []func()
}

func ConcatMap[S, T, U any](m machine.Machine[S, T, machine.Void], f func(T) machine.Machine[machine.Void, U, machine.Void]) machine.Machine[S, U, machine.Void] {
	return &concatMap[S, T, U]{
		outerGoTo:   m.GoTo(),
		outer:       m,
		transformer: f,
	}
}

func (m *concatMap[S, T, U]) GoTo() machine.GoTo {
	if m.err != nil || (m.inner != nil && m.innerGoTo == machine.GoToYield) {
		return machine.GoToYield
	}

	if m.inner == nil && m.outerGoTo == machine.GoToAwait {
		return machine.GoToAwait
	}

	if m.inner == nil && m.outerGoTo == machine.GoToReturn {
		return machine.GoToReturn
	}

	return machine.GoToContinue
}

func (m *concatMap[S, T, U]) Await(x S, err error) {
	m.outer.Await(x, err)
	m.outerGoTo = m.outer.GoTo()
}

func (m *concatMap[S, T, U]) DontWait() {
	m.outer.DontWait()
	m.outerGoTo = m.outer.GoTo()
}

func (m *concatMap[S, T, U]) Yield() (U, error) {
	if m.err != nil {
		err := m.err
		m.err = nil
		return m._nil, err
	}

	x, err := m.inner.Yield()
	m.innerGoTo = m.inner.GoTo()
	return x, err
}

func (m *concatMap[S, T, U]) Continue() {
	if m.inner != nil && m.innerGoTo == machine.GoToReturn {
		m.inner = nil
		return
	}

	if m.inner != nil && m.innerGoTo == machine.GoToContinue {
		m.inner.Continue()
		m.innerGoTo = m.inner.GoTo()
		return
	}

	if m.outerGoTo == machine.GoToContinue {
		m.outer.Continue()
		m.outerGoTo = m.outer.GoTo()
		return
	}

	x, err := m.outer.Yield()
	m.outerGoTo = m.outer.GoTo()
	if err != nil {
		m.err = err
		return
	}

	m.inner = m.transformer(x)
	m.deferStack = append(m.deferStack)
	m.innerGoTo = m.inner.GoTo()
}

func (m *concatMap[S, T, U]) Defer() {
	for _, f := range m.deferStack {
		if f != nil {
			defer f()
		}
	}
}

func (m *concatMap[S, T, U]) Return() (machine.Void, error) {
	return machine.Void{}, nil
}
