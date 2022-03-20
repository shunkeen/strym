package try

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/prosumer"
)

type tryMap[S, T any] struct {
	prosumer.Base
	buffer      S
	transformer func(S) (T, error)
}

func TryMap[S, T any](f func(S) (T, error)) prosumer.Try[S, T] {
	return &tryMap[S, T]{
		transformer: f,
		Base:        prosumer.NewBase(),
	}
}

func (m *tryMap[S, T]) Await(x S) {
	m.BaseGoTo = machine.GoToYield
	m.buffer = x
}

func (m *tryMap[S, T]) Yield() (T, error) {
	m.BaseGoTo = machine.GoToAwait
	return m.transformer(m.buffer)
}
