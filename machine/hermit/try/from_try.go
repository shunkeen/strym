package try

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/hermit"
)

type fromTry[T any] struct {
	hermit.Base
	val T
	err error
}

func FromTry[T any](x T, err error) hermit.Try[T] {
	return &fromTry[T]{
		val: x,
		err: err,
	}
}

func (m *fromTry[T]) GoTo() machine.GoTo {
	return machine.GoToReturn
}

func (m *fromTry[T]) Return() (T, error) {
	return m.val, m.err
}
