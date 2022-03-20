package try

import (
	"errors"

	"github.com/shunkeen/strym/machine/producer"
	"github.com/shunkeen/strym/machine/producer/naive"
	"github.com/shunkeen/strym/machine/producer/throw"
	"golang.org/x/exp/constraints"
)

func RangeBy(start, step, stop int) producer.Try[int] {
	if stop == 0 {
		return rangeThrow[int]("RangeBy")
	}

	m := naive.RangeIntegerBy(start, step, stop)
	return LiftNaive(m)
}

func RangeIntegerBy[T constraints.Integer](start, step, stop T) producer.Try[T] {
	if stop == 0 {
		return rangeThrow[T]("RangeIntegerBy")
	}

	m := naive.RangeIntegerBy(start, step, stop)
	return LiftNaive(m)
}

func RangeFloatBy[T constraints.Float](start, step, stop T) producer.Try[T] {
	if stop == 0 {
		return rangeThrow[T]("RangeFloatBy")
	}

	m := naive.RangeFloatBy(start, step, stop)
	return LiftNaive(m)
}

func rangeThrow[T any](s string) producer.Try[T] {
	e := errors.New("strym." + s + ": zero step")
	m := throw.ThrowOnce(e)
	return LiftThrow[T](m)
}
