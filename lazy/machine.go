package lazy

import "github.com/shunkeen/strym/machine"

type Void = machine.Void

type Hermit[T any] struct {
	machine.Machine[Void, Void, T]
}

func NewHermit[T any](m machine.Machine[Void, Void, T]) Hermit[T] {
	return Hermit[T]{Machine: m}
}

type Producer[T any] struct {
	machine.Machine[Void, T, Void]
}

func NewProducer[T any](m machine.Machine[Void, T, Void]) Producer[T] {
	return Producer[T]{Machine: m}
}

type Prosumer[S, T any] struct {
	machine.Machine[S, T, Void]
}

func NewProsumer[S, T any](m machine.Machine[S, T, Void]) Prosumer[S, T] {
	return Prosumer[S, T]{Machine: m}
}

type Consumer[S, T any] struct {
	machine.Machine[S, Void, T]
}

func NewConsumer[S, T any](m machine.Machine[S, Void, T]) Consumer[S, T] {
	return Consumer[S, T]{Machine: m}
}
