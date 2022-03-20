package try

import "github.com/shunkeen/strym/machine/producer"

type liftNaive[T any] struct {
	producer.Naive[T]
}

func LiftNaive[T any](m producer.Naive[T]) producer.Try[T] {
	return &liftNaive[T]{Naive: m}
}

func (m *liftNaive[T]) Yield() (T, error) {
	return m.Naive.Yield(), nil
}
