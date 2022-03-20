package eager

import (
	"github.com/shunkeen/strym/machine"
)

type Hermit[T any] func() (T, error)

func NewHermit[T any](m machine.Machine[Void, Void, T]) Hermit[T] {
	defer m.Defer()

	for {
		switch m.GoTo() {
		case machine.GoToContinue:
			m.Continue()

		case machine.GoToReturn:
			x, err := m.Return()
			return func() (T, error) { return x, err }

		default:
			panic("eager.Hermit: undefined state")
		}
	}
}
