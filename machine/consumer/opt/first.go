package opt

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/consumer"
)

type first[T any] struct {
	consumer.Base
	ok     bool
	buffer T
}

func First[T any]() consumer.Opt[T, T] {
	m := consumer.NewBase()
	return &first[T]{Base: m}
}

func (m *first[T]) Await(x T) {
	m.BaseGoTo = machine.GoToReturn
	m.ok = true
	m.buffer = x
}

func (m *first[T]) Return() (T, bool) {
	return m.buffer, m.ok
}
