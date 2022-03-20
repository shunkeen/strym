package tuple

type Tuple0 func()

func New0() Tuple0 {
	return func() {
		return
	}
}

type Tuple1[T0 any] func() T0

func New1[T0 any](x0 T0) Tuple1[T0] {
	return func() T0 {
		return x0
	}
}

type Tuple2[T0, T1 any] func() (T0, T1)

func New2[T0, T1 any](x0 T0, x1 T1) Tuple2[T0, T1] {
	return func() (T0, T1) {
		return x0, x1
	}
}

type Tuple3[T0, T1, T2 any] func() (T0, T1, T2)

func New3[T0, T1, T2 any](x0 T0, x1 T1, x2 T2) Tuple3[T0, T1, T2] {
	return func() (T0, T1, T2) {
		return x0, x1, x2
	}
}
