package monoid

type Monoid[T any] interface {
	Append(T) Monoid[T]
	Get() T
}

type Default[T any] struct {
	get T
}

func (m *Default[T]) Get() T {
	return m.get
}
