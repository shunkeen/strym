package naive

import "github.com/shunkeen/strym/machine/consumer"

type toSlice[T any] struct {
	consumer.Base
	slice []T
}

func ToSlice[T any]() consumer.Naive[T, []T] {
	m := consumer.NewBase()
	return &toSlice[T]{Base: m}
}

func (m *toSlice[T]) Await(x T) {
	m.slice = append(m.slice, x)
}

func (m *toSlice[T]) Return() []T {
	return m.slice
}
