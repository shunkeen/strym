package try

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/prosumer"
)

type tryFilter[T any] struct {
	prosumer.Base
	buffer      T
	_nil        T
	transformer func(T) error
}

func TryFilter[T any](f func(T) error) prosumer.Try[T, T] {
	return &tryFilter[T]{
		transformer: f,
		Base:        prosumer.NewBase(),
	}
}

func (m *tryFilter[T]) Await(x T) {
	m.BaseGoTo = machine.GoToYield
	m.buffer = x
}

func (m *tryFilter[T]) Yield() (T, error) {
	m.BaseGoTo = machine.GoToAwait

	err := m.transformer(m.buffer)
	if err != nil {
		return m._nil, err
	}

	return m.buffer, nil
}
