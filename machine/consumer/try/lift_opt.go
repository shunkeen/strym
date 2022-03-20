package try

import "github.com/shunkeen/strym/machine/consumer"

type liftOpt[S, T any] struct {
	consumer.Opt[S, T]
	err  error
	_nil T
}

func LiftOpt[S, T any](m consumer.Opt[S, T], err error) consumer.Try[S, T] {
	return &liftOpt[S, T]{
		err: err,
		Opt: m,
	}
}

func (m *liftOpt[S, T]) Return() (T, error) {
	x, ok := m.Opt.Return()
	if !ok {
		return m._nil, m.err
	}

	return x, nil
}
