package naive

import (
	"github.com/shunkeen/strym/machine/consumer"
	"golang.org/x/exp/constraints"
)

type sumInteger[T constraints.Integer] struct {
	consumer.Base
	sum T
}

func SumInteger[T constraints.Integer]() consumer.Naive[T, T] {
	m := consumer.NewBase()
	return &sumInteger[T]{Base: m}
}

func (m *sumInteger[T]) Await(x T) {
	m.sum += x
}

func (m *sumInteger[T]) Return() T {
	return m.sum
}
