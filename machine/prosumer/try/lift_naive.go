package try

import "github.com/shunkeen/strym/machine/prosumer"

type liftNaive[S, T any] struct {
	prosumer.Naive[S, T]
}

func LiftNaive[S, T any](m prosumer.Naive[S, T]) prosumer.Try[S, T] {
	return &liftNaive[S, T]{Naive: m}
}

func (m *liftNaive[S, T]) Yield() (T, error) {
	return m.Naive.Yield(), nil
}
