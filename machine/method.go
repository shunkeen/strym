package machine

type Void = struct{}

type Await[T any] interface {
	Await(T)
}

type AwaitTry[T any] interface {
	Await(T, error)
}

type Yield[T any] interface {
	Yield() T
}

type YieldOpt[T any] interface {
	Yield() (T, bool)
}

type YieldTry[T any] interface {
	Yield() (T, error)
}

type Return[T any] interface {
	Return() T
}

type ReturnOpt[T any] interface {
	Return() (T, bool)
}

type ReturnTry[T any] interface {
	Return() (T, error)
}

type Default interface {
	GoTo() GoTo
	DontWait()
	Continue()
	Defer()
}
