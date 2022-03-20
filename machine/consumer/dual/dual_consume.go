package dual

import "github.com/shunkeen/strym/machine/consumer"

type dualConsume[S, T any] struct {
	consumer.Base
	vals []S
	errs []error
	run  func([]S, []error) (T, error)
}

func DualConsume[S, T any](f func([]S, []error) (T, error)) consumer.Dual[S, T] {
	return &dualConsume[S, T]{
		Base: consumer.NewBase(),
		run:  f,
	}
}

func (m *dualConsume[S, T]) Await(x S, err error) {
	m.vals = append(m.vals, x)
	m.errs = append(m.errs, err)
}

func (m *dualConsume[S, T]) Return() (T, error) {
	return m.run(m.vals, m.errs)
}
