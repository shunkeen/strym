package opt

import (
	"github.com/shunkeen/strym/machine/consumer"
)

type last[T any] struct {
	consumer.Base
	ok     bool
	buffer T
}

func Last[T any]() consumer.Opt[T, T] {
	m := consumer.NewBase()
	return &last[T]{Base: m}
}

func (m *last[T]) Await(x T) {
	m.ok = true
	m.buffer = x
}

func (m *last[T]) Return() (T, bool) {
	return m.buffer, m.ok
}
