package lazy

import "github.com/shunkeen/strym/machine"

func Run1[V any](hm Hermit[V]) (V, error) {
	defer hm.Defer()
	for {
		switch hm.GoTo() {
		case machine.GoToContinue:
			hm.Continue()

		case machine.GoToReturn:
			return hm.Return()

		default:
			panic("lazy.Hermit: undefined state")
		}
	}
}

func Run2[S, V any](pd Producer[S], cs Consumer[S, V]) (V, error) {
	hm := ChainHM(pd, cs)
	return Run1(hm)
}

func Run3[S, U, V any](pd Producer[S], ps1 Prosumer[S, U], cs Consumer[U, V]) (V, error) {
	pd2 := ChainPD(pd, ps1)
	return Run2(pd2, cs)
}

func Run4[S, T1, U, V any](pd Producer[S], ps1 Prosumer[S, T1], ps2 Prosumer[T1, U], cs Consumer[U, V]) (V, error) {
	return Run3(pd, Chain2(ps1, ps2), cs)
}

func Run5[S, T1, T2, U, V any](pd Producer[S], ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, U], cs Consumer[U, V]) (V, error) {
	return Run3(pd, Chain3(ps1, ps2, ps3), cs)
}

func Run6[S, T1, T2, T3, U, V any](pd Producer[S], ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, U], cs Consumer[U, V]) (V, error) {
	return Run3(pd, Chain4(ps1, ps2, ps3, ps4), cs)
}

func Run7[S, T1, T2, T3, T4, U, V any](pd Producer[S], ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, U], cs Consumer[U, V]) (V, error) {
	return Run3(pd, Chain5(ps1, ps2, ps3, ps4, ps5), cs)
}

func Run8[S, T1, T2, T3, T4, T5, U, V any](pd Producer[S], ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, U], cs Consumer[U, V]) (V, error) {
	return Run3(pd, Chain6(ps1, ps2, ps3, ps4, ps5, ps6), cs)
}

func Run9[S, T1, T2, T3, T4, T5, T6, U, V any](pd Producer[S], ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, T6], ps7 Prosumer[T6, U], cs Consumer[U, V]) (V, error) {
	return Run3(pd, Chain7(ps1, ps2, ps3, ps4, ps5, ps6, ps7), cs)
}

func Run10[S, T1, T2, T3, T4, T5, T6, T7, U, V any](pd Producer[S], ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, T6], ps7 Prosumer[T6, T7], ps8 Prosumer[T7, U], cs Consumer[U, V]) (V, error) {
	return Run3(pd, Chain8(ps1, ps2, ps3, ps4, ps5, ps6, ps7, ps8), cs)
}

func Run11[S, T1, T2, T3, T4, T5, T6, T7, T8, U, V any](pd Producer[S], ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, T6], ps7 Prosumer[T6, T7], ps8 Prosumer[T7, T8], ps9 Prosumer[T8, U], cs Consumer[U, V]) (V, error) {
	return Run3(pd, Chain9(ps1, ps2, ps3, ps4, ps5, ps6, ps7, ps8, ps9), cs)
}

func Run12[S, T1, T2, T3, T4, T5, T6, T7, T8, T9, U, V any](pd Producer[S], ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, T6], ps7 Prosumer[T6, T7], ps8 Prosumer[T7, T8], ps9 Prosumer[T8, T9], ps10 Prosumer[T9, U], cs Consumer[U, V]) (V, error) {
	return Run3(pd, Chain10(ps1, ps2, ps3, ps4, ps5, ps6, ps7, ps8, ps9, ps10), cs)
}
