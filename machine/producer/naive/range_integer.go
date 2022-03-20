package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/producer"
	"golang.org/x/exp/constraints"
)

type rangeInteger[T constraints.Integer] struct {
	producer.Base
	up      bool
	current T
	step    T
	stop    T
}

func RangeInteger[T constraints.Integer](stop T) producer.Naive[T] {
	return RangeIntegerBy(0, stop, 1)
}

func RangeIntegerTo[T constraints.Integer](start, stop T) producer.Naive[T] {
	return RangeIntegerBy(start, stop, 1)
}

func RangeIntegerBy[T constraints.Integer](start, stop, step T) producer.Naive[T] {
	return &rangeInteger[T]{
		up:      0 < step,
		current: start,
		step:    step,
		stop:    stop,
	}
}

func (m *rangeInteger[T]) GoTo() machine.GoTo {
	if (m.up && m.current < m.stop) || (!m.up && m.current > m.stop) {
		return machine.GoToYield
	}

	return machine.GoToReturn
}

func (m *rangeInteger[T]) Yield() T {
	x := m.current
	m.current += m.step
	return x
}
