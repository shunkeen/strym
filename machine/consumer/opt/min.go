package opt

import (
	"github.com/shunkeen/strym/machine/consumer"
	"golang.org/x/exp/constraints"
)

type min[T constraints.Ordered] struct {
	consumer.Base
	ok     bool
	buffer T
}

func Min[T constraints.Ordered]() consumer.Opt[T, T] {
	m := consumer.NewBase()
	return &min[T]{Base: m}
}

func (m *min[T]) Await(x T) {
	if !m.ok || m.buffer > x {
		m.ok = true
		m.buffer = x
	}
}

func (m *min[T]) Return() (T, bool) {
	return m.buffer, m.ok
}
