package lazy_test

import (
	"strconv"
	"testing"

	"github.com/shunkeen/strym/lazy"
	"github.com/stretchr/testify/assert"
)

func TestExampleToSlice(t *testing.T) {
	xs, _ := lazy.Run2(
		lazy.FromSlice([]int{0, 1, 2, 3, 4}),
		lazy.ToSlice[int](),
	)
	// []int{0, 1, 2, 3, 4}
	assert.Equal(t, []int{0, 1, 2, 3, 4}, xs)
}

func TestExampleRedirect(t *testing.T) {
	// import "strconv"
	// import "github.com/shunkeen/eager/datatype/tuple"
	tuple2, _ := lazy.Run3(
		lazy.FromSlice([]string{"1", "a", "2"}),
		lazy.TryMap(strconv.Atoi),
		lazy.Redirect(
			lazy.ToSlice[int](),
			lazy.ToSlice[error](),
		),
	)
	xs, es := tuple2()
	// []int{1, 2}, []error{*errors.errorString {s: "invalid syntax"}}
	assert.Equal(t, []int{1, 2}, xs)
	assert.Equal(t, 1, len(es))
	assert.EqualError(t, es[0], "strconv.Atoi: parsing \"a\": invalid syntax")
}

func TestExampleFirst(t *testing.T) {
	x, _ := lazy.Run2(
		lazy.FromSlice([]int{1, 2, 3}),
		lazy.First[int](),
	)
	// 1
	assert.Equal(t, 1, x)
}

func TestExampleLast(t *testing.T) {
	x, _ := lazy.Run2(
		lazy.FromSlice([]int{1, 2, 3}),
		lazy.Last[int](),
	)
	// 3
	assert.Equal(t, 3, x)
}

func TestExampleNth(t *testing.T) {
	x, _ := lazy.Run2(
		lazy.FromSlice([]int{1, 2, 3}),
		lazy.Nth[int](1),
	)
	// 2
	assert.Equal(t, 2, x)
}

func TestExampleIncludes(t *testing.T) {
	x, _ := lazy.Run2(
		lazy.FromSlice([]int{1, 2, 3, 4, 5}),
		lazy.Includes(3),
	)
	// true
	assert.Equal(t, true, x)
}

func TestExampleMax(t *testing.T) {
	x, _ := lazy.Run2(
		lazy.FromSlice([]int{3, 4, 2, 0, 1}),
		lazy.Max[int](),
	)
	// 4
	assert.Equal(t, 4, x)
}

func TestExampleMin(t *testing.T) {
	x, _ := lazy.Run2(
		lazy.FromSlice([]int{3, 4, 2, 0, 1}),
		lazy.Min[int](),
	)
	// 0
	assert.Equal(t, 0, x)
}

func TestExampleSum(t *testing.T) {
	x, _ := lazy.Run2(
		lazy.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		lazy.Sum(),
	)
	// int(55)
	assert.Equal(t, int(55), x)

	y, _ := lazy.Run3(
		lazy.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		lazy.Map(func(x int) int64 { return int64(x) }),
		lazy.SumInteger[int64](),
	)
	// int64(55)
	assert.Equal(t, int64(55), y)

	z, _ := lazy.Run2(
		lazy.FromSlice([]float64{4.1, 2.0, 1.7}),
		lazy.SumFloat[float64](),
	)
	// float64(7.8)
	assert.Equal(t, float64(7.8), z)
}

func TestExampleProduct(t *testing.T) {
	x, _ := lazy.Run2(
		lazy.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		lazy.Product(),
	)
	// int(3628800)
	assert.Equal(t, int(3628800), x)

	y, _ := lazy.Run3(
		lazy.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		lazy.Map(func(x int) int64 { return int64(x) }),
		lazy.ProductInteger[int64](),
	)
	// int64(3628800)
	assert.Equal(t, int64(3628800), y)

	z, _ := lazy.Run2(
		lazy.FromSlice([]float64{4.1, 2.0, 1.7}),
		lazy.ProductFloat[float64](),
	)
	// float64(13.939999999999998)
	assert.Equal(t, float64(13.939999999999998), z)
}

func TestExampleIsEmpty(t *testing.T) {
	x, _ := lazy.Run2(lazy.FromSlice([]int{}), lazy.IsEmpty[int]())
	// true
	assert.Equal(t, true, x)

	y, _ := lazy.Run2(lazy.FromSlice([]int{1}), lazy.IsEmpty[int]())
	// false
	assert.Equal(t, false, y)
}

func TestExampleCount(t *testing.T) {
	x, _ := lazy.Run2(
		lazy.FromSlice([]string{"a", "b", "c"}),
		lazy.Count[string](),
	)
	// 3
	assert.Equal(t, 3, x)
}

func TestExampleAnd(t *testing.T) {
	x, _ := lazy.Run2(
		lazy.FromSlice([]bool{true, false, true}),
		lazy.And(),
	)
	// false
	assert.Equal(t, false, x)
}

func TestExampleOr(t *testing.T) {
	x, _ := lazy.Run2(
		lazy.FromSlice([]bool{true, false, true}),
		lazy.Or(),
	)
	// true
	assert.Equal(t, true, x)
}

func TestExampleAll(t *testing.T) {
	x, _ := lazy.Run2(
		lazy.FromSlice([]int{3, 1, 4}),
		lazy.All(func(x int) bool { return x > 2 }),
	)
	// false
	assert.Equal(t, false, x)
}

func TestExampleAny(t *testing.T) {
	x, _ := lazy.Run2(
		lazy.FromSlice([]int{3, 1, 4}),
		lazy.Any(func(x int) bool { return x > 2 }),
	)
	// true
	assert.Equal(t, true, x)
}

func TestExampleForEach(t *testing.T) {
	var acc string
	lazy.Run2(
		lazy.FromSlice([]string{"a", "b", "c"}),
		lazy.ForEach(func(x string) { acc += x }),
	)
	// "abc"
	assert.Equal(t, "abc", acc)
}

func TestExampleReduce(t *testing.T) {
	x, _ := lazy.Run2(
		lazy.FromSlice([]string{"a", "b", "c", "d"}),
		lazy.Reduce("foo", func(x, y string) string { return x + y }),
	)
	// "fooabcd"
	assert.Equal(t, "fooabcd", x)
}

func TestExampleReduce1(t *testing.T) {
	x, _ := lazy.Run2(
		lazy.FromSlice([]int{1, 2, 3, 4}),
		lazy.Reduce1(func(x, y int) int { return x - y }),
	)
	// -8
	assert.Equal(t, -8, x)
}

func TestExampleSparkWith(t *testing.T) {
	x, _ := lazy.Run2(
		lazy.FromSlice([]int{1, 2, 3, 4}),
		lazy.SparkWith(
			lazy.Sum(),
			lazy.Count[int](),
			func(sum, count int) float64 {
				return float64(sum) / float64(count)
			},
		),
	)
	// 2.5
	assert.Equal(t, 2.5, x)
}

func TestExampleSpark(t *testing.T) {
	// import "github.com/shunkeen/eager/datatype/tuple"
	tuple, _ := lazy.Run2(
		lazy.FromSlice([]int{0, 1, 2, 3, 4}),
		lazy.Spark(
			lazy.Sum(),
			lazy.ChainCS(
				lazy.Drop[int](2),
				lazy.Product(),
			),
		),
	)
	x, y := tuple()
	// 10, 24
	assert.Equal(t, 10, x)
	assert.Equal(t, 24, y)
}
