package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/consumer"
)

type naiveAny[T any] struct {
	consumer.Base
	any       bool
	predicate func(T) bool
}

func Any[T any](p func(T) bool) consumer.Naive[T, bool] {
	return &naiveAny[T]{
		predicate: p,
		Base:      consumer.NewBase(),
	}
}

func (m *naiveAny[T]) Await(x T) {
	if m.predicate(x) {
		m.BaseGoTo = machine.GoToReturn
		m.any = true
	}
}

func (m *naiveAny[T]) Return() bool {
	return m.any
}
