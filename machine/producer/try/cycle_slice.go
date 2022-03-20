package try

import (
	"errors"

	"github.com/shunkeen/strym/machine/producer"
	"github.com/shunkeen/strym/machine/producer/naive"
	"github.com/shunkeen/strym/machine/producer/throw"
)

func CycleSlice[T any](xs []T) producer.Try[T] {
	if len(xs) == 0 {
		e := errors.New("strym.CycleSlice: empty slice")
		m := throw.ThrowOnce(e)
		return LiftThrow[T](m)
	}

	m := naive.CycleSlice(xs)
	return LiftNaive(m)
}
