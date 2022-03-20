package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/consumer"
)

type isEmpty[T any] struct {
	consumer.Base
	isEmpty bool
}

func IsEmpty[T any]() consumer.Naive[T, bool] {
	return &isEmpty[T]{
		isEmpty: true,
		Base:    consumer.NewBase(),
	}
}

func (m *isEmpty[T]) Await(x T) {
	m.BaseGoTo = machine.GoToReturn
	m.isEmpty = false
}

func (m *isEmpty[T]) Return() bool {
	return m.isEmpty
}
