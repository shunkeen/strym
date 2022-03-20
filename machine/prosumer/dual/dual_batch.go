package dual

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/prosumer"
)

type dualBatch[S, T any] struct {
	prosumer.Base
	doneInput bool
	doneBatch bool
	index     int
	_nil      T
	vals1     []S
	errs1     []error
	vals2     []T
	errs2     []error
	run       func([]S, []error) ([]T, []error)
}

func DualBatch[S, T any](f func([]S, []error) ([]T, []error)) prosumer.Dual[S, T] {
	return &dualBatch[S, T]{
		run:  f,
		Base: prosumer.NewBase(),
	}
}

func (m *dualBatch[S, T]) GoTo() machine.GoTo {
	if !m.doneInput {
		return machine.GoToAwait
	}

	if m.doneBatch || (m.index >= len(m.vals2) && m.index < len(m.errs2) && m.errs2[m.index] != nil) {
		return machine.GoToContinue
	}

	if m.index < len(m.vals2) || m.index < len(m.errs2) {
		return machine.GoToYield
	}

	if m.index < len(m.errs2) {
		return machine.GoToYield
	}

	return machine.GoToReturn
}

func (m *dualBatch[S, T]) Await(x S, err error) {
	m.vals1 = append(m.vals1, x)
	m.errs1 = append(m.errs1, err)
}

func (m *dualBatch[S, T]) DontWait() {
	m.doneInput = true
}

func (m *dualBatch[S, T]) Continue() {
	if !m.doneBatch {
		m.doneBatch = true
		m.vals2, m.errs2 = m.run(m.vals1, m.errs1)
		return
	}

	m.index++
}

func (m *dualBatch[S, T]) Yield() (T, error) {
	i := m.index
	m.index++

	if i < len(m.vals2) && i < len(m.errs2) {
		if m.errs2[i] != nil {
			return m._nil, m.errs2[i]
		}

		return m.vals2[i], nil
	}

	if i >= len(m.vals2) {
		return m._nil, m.errs2[i]
	}

	return m.vals2[i], nil
}
