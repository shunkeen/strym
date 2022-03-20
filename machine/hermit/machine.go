package hermit

import "github.com/shunkeen/strym/machine"

type Naive[T any] interface {
	machine.Return[T]
	Default
}

type Throw interface {
	machine.Return[error]
	Default
}

type Opt[T any] interface {
	machine.ReturnOpt[T]
	Default
}

type Try[T any] interface {
	machine.ReturnTry[T]
	Default
}

type Default interface {
	machine.Await[machine.Void]
	machine.YieldTry[machine.Void]
	machine.Default
}
