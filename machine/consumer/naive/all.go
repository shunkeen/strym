package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/consumer"
)

type all[T any] struct {
	consumer.Base
	all       bool
	predicate func(T) bool
}

func All[T any](p func(T) bool) consumer.Naive[T, bool] {
	return &all[T]{
		all:       true,
		predicate: p,
		Base:      consumer.NewBase(),
	}
}

func (m *all[T]) Await(x T) {
	if !m.predicate(x) {
		m.BaseGoTo = machine.GoToReturn
		m.all = false
	}
}

func (m *all[T]) Return() bool {
	return m.all
}
