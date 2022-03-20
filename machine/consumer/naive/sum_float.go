package naive

import (
	"github.com/shunkeen/strym/datatype/monoid"
	"github.com/shunkeen/strym/machine/consumer"
	"golang.org/x/exp/constraints"
)

type sumFloat[T constraints.Float] struct {
	consumer.Base
	sum monoid.Monoid[T]
}

func SumFloat[T constraints.Float]() consumer.Naive[T, T] {
	return &sumFloat[T]{
		Base: consumer.NewBase(),
		sum:  monoid.SumFloat[T](),
	}
}

func (m *sumFloat[T]) Await(x T) {
	m.sum.Append(x)
}

func (m *sumFloat[T]) Return() T {
	return m.sum.Get()
}
