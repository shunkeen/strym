package try

import "github.com/shunkeen/strym/machine/consumer"

type liftNaive[S, T any] struct {
	consumer.Naive[S, T]
}

func LiftNaive[S, T any](m consumer.Naive[S, T]) consumer.Try[S, T] {
	return &liftNaive[S, T]{Naive: m}
}

func (m *liftNaive[S, T]) Return() (T, error) {
	return m.Naive.Return(), nil
}
