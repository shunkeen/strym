package monoid

import "golang.org/x/exp/constraints"

type kahan[T constraints.Float] struct {
	Default[T]
	c T
}

func SumFloat[T constraints.Float]() Monoid[T] {
	return &kahan[T]{Default: Default[T]{}}
}

func (m *kahan[T]) Append(x T) Monoid[T] {
	y := x - m.c
	t := m.get + y
	m.c = (t - m.get) - y
	m.get = t
	return m
}
