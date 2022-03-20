package naive

import (
	"github.com/shunkeen/strym/datatype/monoid"
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/producer"
	"golang.org/x/exp/constraints"
)

type rangeFloat[T constraints.Float] struct {
	producer.Base
	up      bool
	step    T
	stop    T
	current monoid.Monoid[T]
}

func RangeFloat[T constraints.Float](stop T) producer.Naive[T] {
	return RangeFloatBy(0, stop, 1)
}

func RangeFloatTo[T constraints.Float](start, stop T) producer.Naive[T] {
	return RangeFloatBy(start, stop, 1)
}

func RangeFloatBy[T constraints.Float](start, stop, step T) producer.Naive[T] {
	m := monoid.SumFloat[T]().Append(start)
	return &rangeFloat[T]{
		up:      0 < step,
		current: m,
		step:    step,
		stop:    stop,
	}
}

func (m *rangeFloat[T]) GoTo() machine.GoTo {
	x := m.current.Get()
	if (m.up && x < m.stop) || (!m.up && x > m.stop) {
		return machine.GoToYield
	}

	return machine.GoToReturn
}

func (m *rangeFloat[T]) Yield() T {
	x := m.current.Get()
	m.current.Append(m.step)
	return x
}
