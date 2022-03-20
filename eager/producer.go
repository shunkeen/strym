package eager

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/producer"
	"github.com/shunkeen/strym/machine/producer/try"
)

type Producer[T any] func() ([]T, []error)

func NewProducer[T any](m machine.Machine[Void, T, Void]) Producer[T] {
	defer m.Defer()

	var xs []T
	var es []error

	for {
		switch m.GoTo() {
		case machine.GoToContinue:
			m.Continue()

		case machine.GoToReturn:
			return func() ([]T, []error) { return xs, es }

		case machine.GoToYield:
			x, err := m.Yield()
			xs = append(xs, x)
			es = append(es, err)

		default:
			panic("eager.Producer: undefined state")
		}
	}
}

func producerMachine[T any](pd Producer[T]) machine.Machine[Void, T, Void] {
	m := try.FromTrySlice(pd())
	return producer.LiftTry(m)
}
