package eager_test

import (
	"errors"
	"testing"

	"github.com/shunkeen/strym/datatype/tuple"
	"github.com/shunkeen/strym/eager"
	"github.com/stretchr/testify/assert"
)

func TestExampleFromSlice(t *testing.T) {
	xs, _ := eager.Run2(
		eager.FromSlice([]int{0, 1, 2, 3, 4}),
		eager.ToSlice[int](),
	)
	// []int{0, 1, 2, 3, 4}
	assert.Equal(t, []int{0, 1, 2, 3, 4}, xs)
}

// func TestExampleIterate(t *testing.T) {
// 	xs, _ := eager.Run3(
// 		eager.Iterate(true, func(x bool) bool { return !x }),
// 		eager.Take[bool](5),
// 		eager.ToSlice[bool](),
// 	)
// 	// []bool{true, false, true, false, true}
// 	assert.Equal(t, []bool{true, false, true, false, true}, xs)
// }

// func TestExampleRepeat(t *testing.T) {
// 	xs, _ := eager.Run3(
// 		eager.Repeat(17),
// 		eager.Take[int](5),
// 		eager.ToSlice[int](),
// 	)
// 	// []int{17, 17, 17, 17, 17}
// 	assert.Equal(t, []int{17, 17, 17, 17, 17}, xs)
// }

// func TestExampleCycleSlice(t *testing.T) {
// 	xs, _ := eager.Run3(
// 		eager.CycleSlice([]int{2, 5, 7}),
// 		eager.Take[int](5),
// 		eager.ToSlice[int](),
// 	)
// 	// []int{2, 5, 7, 2, 5}
// 	assert.Equal(t, []int{2, 5, 7, 2, 5}, xs)
// }

func TestExampleReplicate(t *testing.T) {
	xs, _ := eager.Run2(
		eager.Replicate(4, true),
		eager.ToSlice[bool](),
	)
	// []bool{true, true, true, true}
	assert.Equal(t, []bool{true, true, true, true}, xs)
}

func TestExampleFlatten(t *testing.T) {
	xs, _ := eager.Run2(
		eager.Flatten(eager.ChainPD(
			eager.FromSlice([]int{0, 1, 2, 3}),
			eager.Map(eager.Range),
		)),
		eager.ToSlice[int](),
	)
	// []int{0, 0, 1, 0, 1, 2}
	assert.Equal(t, []int{0, 0, 1, 0, 1, 2}, xs)
}

func TestExampleFlatMap(t *testing.T) {
	xs, _ := eager.Run2(
		eager.FlatMap(
			eager.FromSlice([]int{0, 1, 2, 3}),
			eager.Range,
		),
		eager.ToSlice[int](),
	)
	// []int{0, 0, 1, 0, 1, 2}
	assert.Equal(t, []int{0, 0, 1, 0, 1, 2}, xs)
}

func TestExampleRange(t *testing.T) {
	xs, _ := eager.Run2(
		eager.Range(5),
		eager.ToSlice[int](),
	)
	// []int{0, 1, 2, 3, 4}
	assert.Equal(t, []int{0, 1, 2, 3, 4}, xs)

	ys, _ := eager.Run2(
		eager.RangeTo(5, 10),
		eager.ToSlice[int](),
	)
	// []int{5, 6, 7, 8, 9}
	assert.Equal(t, []int{5, 6, 7, 8, 9}, ys)

	zs, _ := eager.Run2(
		eager.RangeBy(0, 10, 2),
		eager.ToSlice[int](),
	)
	// []int{0, 2, 4, 6, 8}
	assert.Equal(t, []int{0, 2, 4, 6, 8}, zs)
}

func TestExampleOtherRange(t *testing.T) {
	xs, _ := eager.Run2(
		eager.RangeBy(8, 0, -2),
		eager.ToSlice[int](),
	)
	// []int{8, 6, 4, 2}
	assert.Equal(t, []int{8, 6, 4, 2}, xs)

	ys, _ := eager.Run2(
		eager.RangeInteger[int64](5),
		eager.ToSlice[int64](),
	)
	// []int64{0, 1, 2, 3, 4}
	assert.Equal(t, []int64{0, 1, 2, 3, 4}, ys)

	zs, _ := eager.Run2(
		eager.RangeFloatBy(1.0, 0.5, -0.1),
		eager.ToSlice[float64](),
	)
	// []float64{1.0, 0.9, 0.8, 0.7, 0.6}
	assert.Equal(t, []float64{1.0, 0.9, 0.8, 0.7, 0.6}, zs)
}

func TestExampleZipWith(t *testing.T) {
	xs, _ := eager.Run2(
		eager.ZipWith(
			eager.FromSlice([]int{1, 2, 3}),
			eager.FromSlice([]int{4, 5, 6}),
			func(x, y int) int { return x + y },
		),
		eager.ToSlice[int](),
	)
	// []int{5, 7, 9}
	assert.Equal(t, []int{5, 7, 9}, xs)

	ys, _ := eager.Run2(
		eager.ZipWith3(
			eager.FromSlice([]int{7, 8, 9}),
			eager.FromSlice([]int{4, 5, 6}),
			eager.FromSlice([]int{1, 2, 3}),
			func(x, y, z int) int { return x - y - z },
		),
		eager.ToSlice[int](),
	)
	// []int{2, 1, 0}
	assert.Equal(t, []int{2, 1, 0}, ys)
}

func TestExampleZip(t *testing.T) {
	xs, _ := eager.Run2(
		eager.Zip(
			eager.FromSlice([]int{1, 2}),
			eager.FromSlice([]string{"A", "B"}),
		),
		eager.ToSlice[tuple.Tuple2[int, string]](),
	)
	x1, x2 := xs[0]()
	// 1, "A"
	assert.Equal(t, 1, x1)
	assert.Equal(t, "A", x2)

	x1, x2 = xs[1]()
	// 2, "B"
	assert.Equal(t, 2, x1)
	assert.Equal(t, "B", x2)

	ys, _ := eager.Run2(
		eager.Zip3(
			eager.FromSlice([]int{1, 2}),
			eager.FromSlice([]string{"A", "B"}),
			eager.FromSlice([]bool{true, false}),
		),
		eager.ToSlice[tuple.Tuple3[int, string, bool]](),
	)
	x1, x2, x3 := ys[0]()
	// 1, "A", true
	assert.Equal(t, 1, x1)
	assert.Equal(t, "A", x2)
	assert.Equal(t, true, x3)

	x1, x2, x3 = ys[1]()
	// 2, "B", false
	assert.Equal(t, 2, x1)
	assert.Equal(t, "B", x2)
	assert.Equal(t, false, x3)
}

func TestExampleOnce(t *testing.T) {
	xs, _ := eager.Run2(
		eager.Once(999),
		eager.ToSlice[int](),
	)
	// []int{999}
	assert.Equal(t, []int{999}, xs)
}

func TestExampleThrowOnce(t *testing.T) {
	// import "errors"
	_, err := eager.Run2(
		eager.ThrowOnce[int](errors.New("once")),
		eager.ToSlice[int](),
	)
	// Error{"once"}
	assert.EqualError(t, err, "once")
}
