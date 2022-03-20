package lazy

import "github.com/shunkeen/strym/machine/communicator"

func ChainHM[S, T any](pd Producer[S], cs Consumer[S, T]) Hermit[T] {
	m := communicator.Chain(pd.Machine, cs.Machine)
	return NewHermit(m)
}

func ChainPD[S, T any](pd Producer[S], ps Prosumer[S, T]) Producer[T] {
	m := communicator.Chain(pd.Machine, ps.Machine)
	return NewProducer(m)
}

func ChainCS[S, T, U any](ps Prosumer[S, T], cs Consumer[T, U]) Consumer[S, U] {
	m := communicator.Chain(ps.Machine, cs.Machine)
	return NewConsumer(m)
}

func Chain2[S, T1, U any](ps1 Prosumer[S, T1], ps2 Prosumer[T1, U]) Prosumer[S, U] {
	m1 := ps1.Machine
	m2 := ps2.Machine
	m := communicator.Chain(m1, m2)
	return NewProsumer(m)
}

func Chain3[S, T1, T2, U any](ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, U]) Prosumer[S, U] {
	m1 := Chain2(ps1, ps2).Machine
	m2 := ps3.Machine
	m := communicator.Chain(m1, m2)
	return NewProsumer(m)
}

func Chain4[S, T1, T2, T3, U any](ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, U]) Prosumer[S, U] {
	m1 := Chain2(ps1, ps2).Machine
	m2 := Chain2(ps3, ps4).Machine
	m := communicator.Chain(m1, m2)
	return NewProsumer(m)
}

func Chain5[S, T1, T2, T3, T4, U any](ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, U]) Prosumer[S, U] {
	m1 := Chain2(ps1, ps2).Machine
	m2 := Chain3(ps3, ps4, ps5).Machine
	m := communicator.Chain(m1, m2)
	return NewProsumer(m)
}

func Chain6[S, T1, T2, T3, T4, T5, U any](ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, U]) Prosumer[S, U] {
	m1 := Chain3(ps1, ps2, ps3).Machine
	m2 := Chain3(ps4, ps5, ps6).Machine
	m := communicator.Chain(m1, m2)
	return NewProsumer(m)
}

func Chain7[S, T1, T2, T3, T4, T5, T6, U any](ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, T6], ps7 Prosumer[T6, U]) Prosumer[S, U] {
	m1 := Chain4(ps1, ps2, ps3, ps4).Machine
	m2 := Chain3(ps5, ps6, ps7).Machine
	m := communicator.Chain(m1, m2)
	return NewProsumer(m)
}

func Chain8[S, T1, T2, T3, T4, T5, T6, T7, U any](ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, T6], ps7 Prosumer[T6, T7], ps8 Prosumer[T7, U]) Prosumer[S, U] {
	m1 := Chain4(ps1, ps2, ps3, ps4).Machine
	m2 := Chain4(ps5, ps6, ps7, ps8).Machine
	m := communicator.Chain(m1, m2)
	return NewProsumer(m)
}

func Chain9[S, T1, T2, T3, T4, T5, T6, T7, T8, U any](ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, T6], ps7 Prosumer[T6, T7], ps8 Prosumer[T7, T8], ps9 Prosumer[T8, U]) Prosumer[S, U] {
	m1 := Chain4(ps1, ps2, ps3, ps4).Machine
	m2 := Chain5(ps5, ps6, ps7, ps8, ps9).Machine
	m := communicator.Chain(m1, m2)
	return NewProsumer(m)
}

func Chain10[S, T1, T2, T3, T4, T5, T6, T7, T8, T9, U any](ps1 Prosumer[S, T1], ps2 Prosumer[T1, T2], ps3 Prosumer[T2, T3], ps4 Prosumer[T3, T4], ps5 Prosumer[T4, T5], ps6 Prosumer[T5, T6], ps7 Prosumer[T6, T7], ps8 Prosumer[T7, T8], ps9 Prosumer[T8, T9], ps10 Prosumer[T9, U]) Prosumer[S, U] {
	m1 := Chain5(ps1, ps2, ps3, ps4, ps5).Machine
	m2 := Chain5(ps6, ps7, ps8, ps9, ps10).Machine
	m := communicator.Chain(m1, m2)
	return NewProsumer(m)
}
