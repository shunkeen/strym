package eager

func Run1[V any](hm Hermit[V]) (V, error) {
	return hm()
}

func Run2[S, V any](pd Producer[S], cs Consumer[S, V]) (V, error) {
	return cs(pd())
}

func Run3[S, U, V any](pd Producer[S], ps1 Prosumer[S, U], cs Consumer[U, V]) (V, error) {
	return cs(ps1(pd()))
}

func Run4[S, T1, U, V any](pd Producer[S], ps1 Prosumer[S, T1], ps2 Prosumer[T1, U], cs Consumer[U, V]) (V, error) {
	return cs(ps2(ps1(pd())))
}

func Run5[S, T1, T2, U, V any](pd Producer[S], ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, U], cs Consumer[U, V]) (V, error) {
	return cs(ps3(ps2(ps1(pd()))))
}

func Run6[S, T1, T2, T3, U, V any](pd Producer[S], ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, U], cs Consumer[U, V]) (V, error) {
	return cs(ps4(ps3(ps2(ps1(pd())))))
}

func Run7[S, T1, T2, T3, T4, U, V any](pd Producer[S], ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, U], cs Consumer[U, V]) (V, error) {
	return cs(ps5(ps4(ps3(ps2(ps1(pd()))))))
}

func Run8[S, T1, T2, T3, T4, T5, U, V any](pd Producer[S], ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, U], cs Consumer[U, V]) (V, error) {
	return cs(ps6(ps5(ps4(ps3(ps2(ps1(pd())))))))
}

func Run9[S, T1, T2, T3, T4, T5, T6, U, V any](pd Producer[S], ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, T6], ps7 Prosumer[T6, U], cs Consumer[U, V]) (V, error) {
	return cs(ps7(ps6(ps5(ps4(ps3(ps2(ps1(pd()))))))))
}

func Run10[S, T1, T2, T3, T4, T5, T6, T7, U, V any](pd Producer[S], ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, T6], ps7 Prosumer[T6, T7], ps8 Prosumer[T7, U], cs Consumer[U, V]) (V, error) {
	return cs(ps8(ps7(ps6(ps5(ps4(ps3(ps2(ps1(pd())))))))))
}

func Run11[S, T1, T2, T3, T4, T5, T6, T7, T8, U, V any](pd Producer[S], ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, T6], ps7 Prosumer[T6, T7], ps8 Prosumer[T7, T8], ps9 Prosumer[T8, U], cs Consumer[U, V]) (V, error) {
	return cs(ps9(ps8(ps7(ps6(ps5(ps4(ps3(ps2(ps1(pd()))))))))))
}

func Run12[S, T1, T2, T3, T4, T5, T6, T7, T8, T9, U, V any](pd Producer[S], ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, T6], ps7 Prosumer[T6, T7], ps8 Prosumer[T7, T8], ps9 Prosumer[T8, T9], ps10 Prosumer[T9, U], cs Consumer[U, V]) (V, error) {
	return cs(ps10(ps9(ps8(ps7(ps6(ps5(ps4(ps3(ps2(ps1(pd())))))))))))
}
