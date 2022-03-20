package naive

import (
	"github.com/shunkeen/strym/machine/consumer"
	"golang.org/x/exp/constraints"
)

type product[T constraints.Integer | constraints.Float] struct {
	consumer.Base
	product T
}

func Product[T constraints.Integer | constraints.Float]() consumer.Naive[T, T] {
	return &product[T]{
		product: 1,
		Base:    consumer.NewBase(),
	}
}

func (m *product[T]) Await(x T) {
	m.product *= x
}

func (m *product[T]) Return() T {
	return m.product
}
