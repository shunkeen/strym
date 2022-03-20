package lazy

import (
	"errors"

	"github.com/shunkeen/strym/datatype/tuple"
	"github.com/shunkeen/strym/machine/communicator"
	"github.com/shunkeen/strym/machine/consumer"
	"github.com/shunkeen/strym/machine/consumer/naive"
	"github.com/shunkeen/strym/machine/consumer/opt"
	"github.com/shunkeen/strym/machine/consumer/try"
	"golang.org/x/exp/constraints"
)

func ToSlice[T any]() Consumer[T, []T] {
	m := naive.ToSlice[T]()
	return naiveCS(m)
}

func RedirectWith[R, S, T, U any](cs1 Consumer[R, S], cs2 Consumer[error, T], f func(S, T) U) Consumer[R, U] {
	m1 := cs1.Machine
	m2 := cs2.Machine
	m := communicator.RedirectWith(m1, m2, f)
	return NewConsumer(m)
}

func Redirect[R, S, T any](cs1 Consumer[R, S], cs2 Consumer[error, T]) Consumer[R, tuple.Tuple2[S, T]] {
	return RedirectWith(cs1, cs2, tuple.New2[S, T])
}

func First[T any]() Consumer[T, T] {
	m := opt.First[T]()
	err := errors.New("strym.First: empty stream")
	return optCS(m, err)
}

func Last[T any]() Consumer[T, T] {
	m := opt.Last[T]()
	err := errors.New("strym.Last: empty stream")
	return optCS(m, err)
}

func Nth[T any](n int) Consumer[T, T] {
	m := opt.Nth[T](n)
	err := errors.New("strym.Nth: not found")
	return optCS(m, err)
}

func Includes[T constraints.Ordered](x T) Consumer[T, bool] {
	m := naive.Includes(x)
	return naiveCS(m)
}

func Max[T constraints.Ordered]() Consumer[T, T] {
	m := opt.Max[T]()
	err := errors.New("strym.Max: empty stream")
	return optCS(m, err)
}

func Min[T constraints.Ordered]() Consumer[T, T] {
	m := opt.Min[T]()
	err := errors.New("strym.Min: empty stream")
	return optCS(m, err)
}

func Sum() Consumer[int, int] {
	m := naive.SumInteger[int]()
	return naiveCS(m)
}

func SumInteger[T constraints.Integer]() Consumer[T, T] {
	m := naive.SumInteger[T]()
	return naiveCS(m)
}

func SumFloat[T constraints.Float]() Consumer[T, T] {
	m := naive.SumFloat[T]()
	return naiveCS(m)
}

func Product() Consumer[int, int] {
	m := naive.Product[int]()
	return naiveCS(m)
}

func ProductInteger[T constraints.Integer]() Consumer[T, T] {
	m := naive.Product[T]()
	return naiveCS(m)
}

func ProductFloat[T constraints.Float]() Consumer[T, T] {
	m := naive.Product[T]()
	return naiveCS(m)
}

func IsEmpty[T any]() Consumer[T, bool] {
	m := naive.IsEmpty[T]()
	return naiveCS(m)
}

func Count[T any]() Consumer[T, int] {
	m := naive.Count[T]()
	return naiveCS(m)
}

func And() Consumer[bool, bool] {
	m := naive.And()
	return naiveCS(m)
}

func Or() Consumer[bool, bool] {
	m := naive.Or()
	return naiveCS(m)
}

func All[T any](p func(T) bool) Consumer[T, bool] {
	m := naive.All(p)
	return naiveCS(m)
}

func Any[T any](p func(T) bool) Consumer[T, bool] {
	m := naive.Any(p)
	return naiveCS(m)
}

func ForEach[T any](f func(T)) Consumer[T, Void] {
	m := naive.ForEach(f)
	return naiveCS(m)
}

func SparkWith[S, T1, T2, U any](cs1 Consumer[S, T1], cs2 Consumer[S, T2], f func(T1, T2) U) Consumer[S, U] {
	m1 := cs1.Machine
	m2 := cs2.Machine
	m := communicator.SparkWith(m1, m2, f)
	return NewConsumer(m)
}

func Spark[S, T1, T2 any](cs1 Consumer[S, T1], cs2 Consumer[S, T2]) Consumer[S, tuple.Tuple2[T1, T2]] {
	return SparkWith(cs1, cs2, tuple.New2[T1, T2])
}

func Reduce[S, T any](x T, f func(T, S) T) Consumer[S, T] {
	m := naive.Reduce(x, f)
	return naiveCS(m)
}

func Reduce1[T any](f func(T, T) T) Consumer[T, T] {
	m := opt.Reduce1(f)
	err := errors.New("strym.Reduce1: empty stream")
	return optCS(m, err)
}

func naiveCS[S, T any](m consumer.Naive[S, T]) Consumer[S, T] {
	m2 := try.LiftNaive(m)
	return tryCS(m2)
}

func optCS[S, T any](m consumer.Opt[S, T], err error) Consumer[S, T] {
	m2 := try.LiftOpt(m, err)
	return tryCS(m2)
}

func tryCS[S, T any](m consumer.Try[S, T]) Consumer[S, T] {
	m2 := consumer.LiftTry(m)
	return NewConsumer(m2)
}
