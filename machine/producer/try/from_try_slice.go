package try

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/producer"
)

type fromTrySlice[T any] struct {
	producer.Base
	index int
	_nil  T
	vals  []T
	errs  []error
}

func FromTrySlice[T any](xs []T, es []error) producer.Try[T] {
	return &fromTrySlice[T]{
		vals: xs,
		errs: es,
	}
}

func (m *fromTrySlice[T]) GoTo() machine.GoTo {
	if m.index >= len(m.vals) && m.index < len(m.errs) && m.errs[m.index] != nil {
		return machine.GoToContinue
	}

	if m.index < len(m.vals) || m.index < len(m.errs) {
		return machine.GoToYield
	}

	if m.index < len(m.errs) {
		return machine.GoToYield
	}

	return machine.GoToReturn
}

func (m *fromTrySlice[T]) Yield() (T, error) {
	i := m.index
	m.index++

	if i < len(m.vals) && i < len(m.errs) {
		if m.errs[i] != nil {
			return m._nil, m.errs[i]
		}

		return m.vals[i], nil
	}

	if i >= len(m.vals) {
		return m._nil, m.errs[i]
	}

	return m.vals[i], nil
}

func (m *fromTrySlice[T]) Continue() {
	m.index++
}
