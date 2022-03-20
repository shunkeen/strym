package naive

import (
	"github.com/shunkeen/strym/machine/consumer"
)

type count[T any] struct {
	consumer.Base
	count int
}

func Count[T any]() consumer.Naive[T, int] {
	m := consumer.NewBase()
	return &count[T]{Base: m}
}

func (m *count[T]) Await(_ T) {
	m.count++
}

func (m *count[T]) Return() int {
	return m.count
}
