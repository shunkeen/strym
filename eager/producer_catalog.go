package eager

import (
	"github.com/shunkeen/strym/datatype/tuple"
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/communicator"
	"github.com/shunkeen/strym/machine/producer"
	"github.com/shunkeen/strym/machine/producer/naive"
	"github.com/shunkeen/strym/machine/producer/throw"
	"github.com/shunkeen/strym/machine/producer/try"
	"golang.org/x/exp/constraints"
)

func FromSlice[T any](xs []T) Producer[T] {
	m := naive.FromSlice(xs)
	return naivePD(m)
}

// func Iterate[T any](x T, f func(T) T) Producer[T] {
// 	m := naive.Iterate(x, f)
// 	return naivePD(m)
// }

// func Repeat[T any](x T) Producer[T] {
// 	m := naive.Repeat(x)
// 	return naivePD(m)
// }

// func CycleSlice[T any](xs []T) Producer[T] {
// 	m := naive.CycleSlice(xs)
// 	return naivePD(m)
// }

func Replicate[T any](n int, x T) Producer[T] {
	m := naive.Replicate(n, x)
	return naivePD(m)
}

func Flatten[T any](pd Producer[Producer[T]]) Producer[T] {
	f := func(x Producer[T]) Producer[T] { return x }
	return FlatMap(pd, f)
}

func FlatMap[S, T any](pd Producer[S], f func(S) Producer[T]) Producer[T] {
	g := func(x S) machine.Machine[struct{}, T, struct{}] {
		inner := f(x)
		return producerMachine(inner)
	}

	m := communicator.ConcatMap(producerMachine(pd), g)
	return NewProducer(m)
}

func Range(stop int) Producer[int] {
	m := naive.RangeInteger(stop)
	return naivePD(m)
}

func RangeTo(start, stop int) Producer[int] {
	m := naive.RangeIntegerTo(start, stop)
	return naivePD(m)
}

func RangeBy(start, stop, step int) Producer[int] {
	m := try.RangeBy(start, stop, step)
	return tryPD(m)
}

func RangeInteger[T constraints.Integer](stop T) Producer[T] {
	m := naive.RangeInteger(stop)
	return naivePD(m)
}

func RangeIntegerTo[T constraints.Integer](start, stop T) Producer[T] {
	m := naive.RangeIntegerTo(start, stop)
	return naivePD(m)
}

func RangeIntegerBy[T constraints.Integer](start, stop, step T) Producer[T] {
	m := try.RangeIntegerBy(start, stop, step)
	return tryPD(m)
}

func RangeFloat[T constraints.Float](stop T) Producer[T] {
	m := naive.RangeFloat(stop)
	return naivePD(m)
}

func RangeFloatTo[T constraints.Float](start, stop T) Producer[T] {
	m := naive.RangeFloatTo(start, stop)
	return naivePD(m)
}

func RangeFloatBy[T constraints.Float](start, stop, step T) Producer[T] {
	m := try.RangeFloatBy(start, stop, step)
	return tryPD(m)
}

func ZipWith[T1, T2, T3 any](pd1 Producer[T1], pd2 Producer[T2], f func(T1, T2) T3) Producer[T3] {
	m1 := producerMachine(pd1)
	m2 := producerMachine(pd2)
	m := communicator.ZipWith(m1, m2, f)
	return NewProducer(m)
}

func ZipWith3[T1, T2, T3, U any](pd1 Producer[T1], pd2 Producer[T2], pd3 Producer[T3], f func(T1, T2, T3) U) Producer[U] {
	m1 := producerMachine(pd1)
	m2 := producerMachine(pd2)
	m3 := producerMachine(pd3)
	m := communicator.ZipWith3(m1, m2, m3, f)
	return NewProducer(m)
}

func Zip[T1, T2 any](pd1 Producer[T1], pd2 Producer[T2]) Producer[tuple.Tuple2[T1, T2]] {
	return ZipWith(pd1, pd2, tuple.New2[T1, T2])
}

func Zip3[T1, T2, T3 any](pd1 Producer[T1], pd2 Producer[T2], pd3 Producer[T3]) Producer[tuple.Tuple3[T1, T2, T3]] {
	return ZipWith3(pd1, pd2, pd3, tuple.New3[T1, T2, T3])
}

func Once[T any](x T) Producer[T] {
	m := naive.Once(x)
	return naivePD(m)
}

func ThrowOnce[T any](err error) Producer[T] {
	m := throw.ThrowOnce(err)
	return throwPD[T](m)
}

func naivePD[T any](m producer.Naive[T]) Producer[T] {
	m2 := try.LiftNaive(m)
	return tryPD(m2)
}

func throwPD[T any](m producer.Throw) Producer[T] {
	m2 := try.LiftThrow[T](m)
	return tryPD(m2)
}

func tryPD[T any](m producer.Try[T]) Producer[T] {
	m2 := producer.LiftTry(m)
	return NewProducer(m2)
}
