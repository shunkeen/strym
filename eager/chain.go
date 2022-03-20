package eager

func ChainHM[S, U any](pd Producer[S], cs Consumer[S, U]) Hermit[U] {
	x, err := cs(pd())
	return func() (U, error) { return x, err }
}

func ChainPD[S, T any](pd Producer[S], ps Prosumer[S, T]) Producer[T] {
	xs, es := ps(pd())
	return func() ([]T, []error) { return xs, es }
}

func ChainCS[T, U, V any](ps Prosumer[T, U], cs Consumer[U, V]) Consumer[T, V] {
	return func(xs []T, es []error) (V, error) {
		return cs(ps(xs, es))
	}
}

func Chain2[S, T1, U any](ps1 Prosumer[S, T1], ps2 Prosumer[T1, U]) Prosumer[S, U] {
	return func(xs []S, es []error) ([]U, []error) {
		return ps2(ps1(xs, es))
	}
}

func Chain3[S, T1, T2, U any](ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, U]) Prosumer[S, U] {
	return func(xs []S, es []error) ([]U, []error) {
		return ps3(ps2(ps1(xs, es)))
	}
}

func Chain4[S, T1, T2, T3, U any](ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, U]) Prosumer[S, U] {
	return func(xs []S, es []error) ([]U, []error) {
		return ps4(ps3(ps2(ps1(xs, es))))
	}
}

func Chain5[S, T1, T2, T3, T4, U any](ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, U]) Prosumer[S, U] {
	return func(xs []S, es []error) ([]U, []error) {
		return ps5(ps4(ps3(ps2(ps1(xs, es)))))
	}
}

func Chain6[S, T1, T2, T3, T4, T5, U any](ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, U]) Prosumer[S, U] {
	return func(xs []S, es []error) ([]U, []error) {
		return ps6(ps5(ps4(ps3(ps2(ps1(xs, es))))))
	}
}

func Chain7[S, T1, T2, T3, T4, T5, T6, U any](ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, T6], ps7 Prosumer[T6, U]) Prosumer[S, U] {
	return func(xs []S, es []error) ([]U, []error) {
		return ps7(ps6(ps5(ps4(ps3(ps2(ps1(xs, es)))))))
	}
}

func Chain8[S, T1, T2, T3, T4, T5, T6, T7, U any](ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, T6], ps7 Prosumer[T6, T7], ps8 Prosumer[T7, U]) Prosumer[S, U] {
	return func(xs []S, es []error) ([]U, []error) {
		return ps8(ps7(ps6(ps5(ps4(ps3(ps2(ps1(xs, es))))))))
	}
}

func Chain9[S, T1, T2, T3, T4, T5, T6, T7, T8, U any](ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, T6], ps7 Prosumer[T6, T7], ps8 Prosumer[T7, T8], ps9 Prosumer[T8, U]) Prosumer[S, U] {
	return func(xs []S, es []error) ([]U, []error) {
		return ps9(ps8(ps7(ps6(ps5(ps4(ps3(ps2(ps1(xs, es)))))))))
	}
}

func Chain10[S, T1, T2, T3, T4, T5, T6, T7, T8, T9, U any](ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, T6], ps7 Prosumer[T6, T7], ps8 Prosumer[T7, T8], ps9 Prosumer[T8, T9], ps10 Prosumer[T9, U]) Prosumer[S, U] {
	return func(xs []S, es []error) ([]U, []error) {
		return ps10(ps9(ps8(ps7(ps6(ps5(ps4(ps3(ps2(ps1(xs, es))))))))))
	}
}
