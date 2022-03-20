package machine

type Machine[S, T, U any] interface {
	AwaitTry[S]
	YieldTry[T]
	ReturnTry[U]
	Default
}
