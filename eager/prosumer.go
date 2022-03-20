package eager

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/prosumer/dual"
)

type Prosumer[S, T any] func([]S, []error) ([]T, []error)

func NewProsumer[S, T any](m machine.Machine[S, T, Void]) Prosumer[S, T] {
	return func(xs []S, es []error) ([]T, []error) {
		defer m.Defer()

		var i int
		var xs2 []T
		var es2 []error

		for {
			switch m.GoTo() {
			case machine.GoToContinue:
				m.Continue()

			case machine.GoToReturn:
				return xs2, es2

			case machine.GoToYield:
				x, err := m.Yield()
				xs2 = append(xs2, x)
				es2 = append(es2, err)

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
				panic("eager.Prosumer: undefined state")
			}
		}
	}
}

func prosumerMachine[S, T any](ps Prosumer[S, T]) machine.Machine[S, T, Void] {
	return dual.DualBatch(ps)
}
