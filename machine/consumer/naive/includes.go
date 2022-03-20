package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/consumer"
	"golang.org/x/exp/constraints"
)

type includes[T constraints.Ordered] struct {
	consumer.Base
	includes bool
	element  T
}

func Includes[T constraints.Ordered](x T) consumer.Naive[T, bool] {
	return &includes[T]{
		element: x,
		Base:    consumer.NewBase(),
	}
}

func (m *includes[T]) Await(x T) {
	if x == m.element {
		m.BaseGoTo = machine.GoToReturn
		m.includes = true
	}
}

func (m *includes[T]) Return() bool {
	return m.includes
}
