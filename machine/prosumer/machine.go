package prosumer

import "github.com/shunkeen/strym/machine"

type Naive[S, T any] interface {
	machine.Await[S]
	machine.Yield[T]
	Default
}

type Throw[S any] interface {
	machine.Await[S]
	machine.Yield[error]
	Default
}

type Opt[S, T any] interface {
	machine.Await[S]
	machine.YieldOpt[T]
	Default
}

type Try[S, T any] interface {
	machine.Await[S]
	machine.YieldTry[T]
	Default
}

type Catch[T any] interface {
	machine.Await[error]
	machine.Yield[T]
	Default
}

type Rethrow interface {
	machine.Await[error]
	machine.Yield[error]
	Default
}

type NearRetry[S, T any] interface {
	machine.Await[error]
	machine.YieldOpt[T]
	Default
}

type Retry[S, T any] interface {
	machine.Await[error]
	machine.YieldTry[T]
	Default
}

type Cleanup[S, T any] interface {
	machine.AwaitTry[S]
	machine.Yield[T]
	Default
}

type Cleandown[T any] interface {
	machine.AwaitTry[T]
	machine.Yield[error]
	Default
}

type NearDual[S, T any] interface {
	machine.AwaitTry[S]
	machine.YieldOpt[T]
	Default
}

type Dual[S, T any] interface {
	machine.AwaitTry[S]
	machine.YieldTry[T]
	Default
}

type Default interface {
	machine.ReturnTry[machine.Void]
	machine.Default
}
