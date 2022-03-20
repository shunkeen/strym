package eager

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/communicator"
	"github.com/shunkeen/strym/machine/prosumer"
	"github.com/shunkeen/strym/machine/prosumer/dual"
	"github.com/shunkeen/strym/machine/prosumer/naive"
	"github.com/shunkeen/strym/machine/prosumer/rethrow"
	"github.com/shunkeen/strym/machine/prosumer/try"
)

func Map[S, T any](f func(S) T) Prosumer[S, T] {
	m := naive.Map(f)
	return naivePS(m)
}

func TryMap[S, T any](f func(S) (T, error)) Prosumer[S, T] {
	m := try.TryMap(f)
	return tryPS(m)
}

func Filter[T any](p func(T) bool) Prosumer[T, T] {
	m := naive.Filter(p)
	return naivePS(m)
}

func TryFilter[T any](p func(T) error) Prosumer[T, T] {
	m := try.TryFilter(p)
	return tryPS(m)
}

func IgnoreErr[T any]() Prosumer[T, T] {
	m := rethrow.IgnoreErr()
	return rethrowPS[T](m)
}

func Reverse[T any]() Prosumer[T, T] {
	m := naive.Reverse[T]()
	return naivePS(m)
}

func Take[T any](n int) Prosumer[T, T] {
	m := naive.Take[T](n)
	return naivePS(m)
}

func Drop[T any](n int) Prosumer[T, T] {
	m := naive.Drop[T](n)
	return naivePS(m)
}

func TakeWhile[T any](p func(T) bool) Prosumer[T, T] {
	m := naive.TakeWhile(p)
	return naivePS(m)
}

func DropWhile[T any](p func(T) bool) Prosumer[T, T] {
	m := naive.DropWhile(p)
	return naivePS(m)
}

func BreakIfErr[T any]() Prosumer[T, T] {
	m := rethrow.BreakIfErr()
	return rethrowPS[T](m)
}

func Concat[T any]() Prosumer[Producer[T], T] {
	f := func(inner Producer[T]) machine.Machine[struct{}, T, struct{}] {
		return producerMachine(inner)
	}

	id := dual.DualId[Producer[T]]()
	m := communicator.ConcatMap[Producer[T], Producer[T]](id, f)
	return NewProsumer(m)
}

func ConcatMap[S, T any](f func(S) Producer[T]) Prosumer[S, T] {
	g := func(x S) machine.Machine[struct{}, T, struct{}] {
		inner := f(x)
		return producerMachine(inner)
	}

	id := dual.DualId[S]()
	m := communicator.ConcatMap[S, S](id, g)
	return NewProsumer(m)
}

func Scan[S, T any](x T, f func(T, S) T) Prosumer[S, T] {
	m := naive.Scan(x, f)
	return naivePS(m)
}

func Scan1[T any](f func(T, T) T) Prosumer[T, T] {
	m := naive.Scan1(f)
	return naivePS(m)
}

func naivePS[S, T any](m prosumer.Naive[S, T]) Prosumer[S, T] {
	m2 := try.LiftNaive(m)
	return tryPS(m2)
}

func rethrowPS[T any](m prosumer.Rethrow) Prosumer[T, T] {
	m2 := dual.LiftRethrow[T](m)
	return NewProsumer[T, T](m2)
}

func tryPS[S, T any](m prosumer.Try[S, T]) Prosumer[S, T] {
	m2 := prosumer.LiftTry(m)
	return NewProsumer(m2)
}
