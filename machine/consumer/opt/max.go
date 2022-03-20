package opt

import (
	"github.com/shunkeen/strym/machine/consumer"
	"golang.org/x/exp/constraints"
)

type max[T constraints.Ordered] struct {
	consumer.Base
	ok     bool
	buffer T
}

func Max[T constraints.Ordered]() consumer.Opt[T, T] {
	m := consumer.NewBase()
	return &max[T]{Base: m}
}

func (m *max[T]) Await(x T) {
	if !m.ok || m.buffer < x {
		m.ok = true
		m.buffer = x
	}
}

func (m *max[T]) Return() (T, bool) {
	return m.buffer, m.ok
}
