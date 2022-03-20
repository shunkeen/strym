package eager

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/consumer/dual"
)

type Consumer[S, T any] func([]S, []error) (T, error)

func NewConsumer[S, T any](m machine.Machine[S, Void, T]) Consumer[S, T] {
	return func(xs []S, es []error) (T, error) {
		defer m.Defer()

		var i int

		for {
			switch m.GoTo() {
			case machine.GoToContinue:
				m.Continue()

			case machine.GoToReturn:
				return m.Return()

			case machine.GoToAwait:
				if i < len(xs) || i < len(es) {
					var x S
					var err error

					if i < len(xs) {
						x = xs[i]
					}

					if i < len(es) {
						err = es[i]
					}

					m.Await(x, err)
					i++
				} else {
					m.DontWait()
				}

			default:
				panic("eager.Consumer: undefined state")
			}
		}
	}
}

func consumerMachine[S, T any](cs Consumer[S, T]) machine.Machine[S, Void, T] {
	return dual.DualConsume(cs)
}
