package consumer

import "github.com/shunkeen/strym/machine"

type Naive[S, T any] interface {
	machine.Await[S]
	machine.Return[T]
	Default
}

type Throw[S any] interface {
	machine.Await[S]
	machine.Return[error]
	Default
}

type Opt[S, T any] interface {
	machine.Await[S]
	machine.ReturnOpt[T]
	Default
}

type Try[S, T any] interface {
	machine.Await[S]
	machine.ReturnTry[T]
	Default
}

type Catch[T any] interface {
	machine.Await[error]
	machine.Return[T]
	Default
}

type Rethrow interface {
	machine.Await[error]
	machine.Return[error]
	Default
}

type NearRetry[S, T any] interface {
	machine.Await[error]
	machine.ReturnOpt[T]
	Default
}

type Retry[S, T any] interface {
	machine.Await[error]
	machine.ReturnTry[T]
	Default
}

type Cleanup[S, T any] interface {
	machine.AwaitTry[S]
	machine.Return[T]
	Default
}

type Cleandown[T any] interface {
	machine.AwaitTry[T]
	machine.Return[error]
	Default
}

type NearDual[S, T any] interface {
	machine.AwaitTry[S]
	machine.ReturnOpt[T]
	Default
}

type Dual[S, T any] interface {
	machine.AwaitTry[S]
	machine.ReturnTry[T]
	Default
}

type Default interface {
	machine.YieldTry[machine.Void]
	machine.Default
}
