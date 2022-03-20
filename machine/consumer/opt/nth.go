package opt

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/consumer"
)

type nth[T any] struct {
	consumer.Base
	count  int
	ok     bool
	buffer T
}

func Nth[T any](n int) consumer.Opt[T, T] {
	return &nth[T]{
		count: n,
		Base:  consumer.NewBase(),
	}
}

func (m *nth[T]) Await(x T) {
	if m.count < 0 {
		m.BaseGoTo = machine.GoToReturn
		return
	}

	if m.count > 0 {
		m.count--
		return
	}

	m.BaseGoTo = machine.GoToReturn
	m.ok = true
	m.buffer = x
}

func (m *nth[T]) Return() (T, bool) {
	return m.buffer, m.ok
}
