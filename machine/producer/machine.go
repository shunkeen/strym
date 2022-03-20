package producer

import "github.com/shunkeen/strym/machine"

type Naive[T any] interface {
	machine.Yield[T]
	Default
}

type Throw interface {
	machine.Yield[error]
	Default
}

type Opt[T any] interface {
	machine.YieldOpt[T]
	Default
}

type Try[T any] interface {
	machine.YieldTry[T]
	Default
}

type Default interface {
	machine.Await[machine.Void]
	machine.ReturnTry[machine.Void]
	machine.Default
}
