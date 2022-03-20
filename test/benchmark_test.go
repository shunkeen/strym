package test

import (
	"testing"

	"github.com/shunkeen/strym/eager"
	"github.com/shunkeen/strym/lazy"
)

func BenchmarkProducer(b *testing.B) {
	stop := 1_000_000

	b.Run("for", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			ys := make([]int, 0, stop)
			for i := 0; i < stop; i++ {
				y := i - stop/2
				ys = append(ys, y)
			}
		}
	})

	b.Run("lazy", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _ = lazy.Run3(
				lazy.Range(stop),
				lazy.Map(func(x int) int { return x - stop/2 }),
				lazy.ToSlice[int](),
			)
		}
	})

	b.Run("eager", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _ = eager.Run3(
				eager.Range(stop),
				eager.Map(func(x int) int { return x - stop/2 }),
				eager.ToSlice[int](),
			)
		}
	})
}

func BenchmarkProsumer(b *testing.B) {
	stop := 1_000_000
	xs, _ := eager.Run2(
		eager.Range(stop),
		eager.ToSlice[int](),
	)

	b.Run("for", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			ys := make([]int, 0, stop)
			for _, i := range xs {
				y := i - stop/2
				ys = append(ys, y)
			}
		}
	})

	b.Run("lazy", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _ = lazy.Run3(
				lazy.FromSlice(xs),
				lazy.Map(func(x int) int { return x - stop/2 }),
				lazy.ToSlice[int](),
			)
		}
	})

	b.Run("eager", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _ = eager.Run3(
				eager.FromSlice(xs),
				eager.Map(func(x int) int { return x - stop/2 }),
				eager.ToSlice[int](),
			)
		}
	})
}

func BenchmarkConsumer(b *testing.B) {
	stop := 1_000_000
	xs, _ := eager.Run2(
		eager.Range(stop),
		eager.ToSlice[int](),
	)

	b.Run("for", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			x := 0
			for _, i := range xs {
				x += i - stop/2
			}
		}
	})

	b.Run("lazy", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _ = lazy.Run3(
				lazy.FromSlice(xs),
				lazy.Map(func(x int) int { return x - stop/2 }),
				lazy.Sum(),
			)
		}
	})

	b.Run("eager", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _ = eager.Run3(
				eager.FromSlice(xs),
				eager.Map(func(x int) int { return x - stop/2 }),
				eager.Sum(),
			)
		}
	})
}

func BenchmarkHermit(b *testing.B) {
	stop := 1_000_000

	b.Run("for", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			x := 0
			for i := 0; i < stop; i++ {
				x += i - stop/2
			}
		}
	})

	b.Run("lazy", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _ = lazy.Run3(
				lazy.Range(stop),
				lazy.Map(func(x int) int { return x - stop/2 }),
				lazy.Sum(),
			)
		}
	})

	b.Run("eager", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, _ = eager.Run3(
				eager.Range(stop),
				eager.Map(func(x int) int { return x - stop/2 }),
				eager.Sum(),
			)
		}
	})
}
