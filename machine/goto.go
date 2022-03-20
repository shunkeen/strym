package machine

type GoTo int

const (
	GoToAwait GoTo = iota
	GoToYield
	GoToReturn
	GoToContinue
)

func (g GoTo) String() string {
	return []string{
		"go to await",
		"go to yield",
		"go to return",
		"go to continue",
	}[g]
}
