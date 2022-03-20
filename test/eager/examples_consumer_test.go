package eager_test

import (
	"strconv"
	"testing"

	"github.com/shunkeen/strym/eager"
	"github.com/stretchr/testify/assert"
)

func TestExampleToSlice(t *testing.T) {
	xs, _ := eager.Run2(
		eager.FromSlice([]int{0, 1, 2, 3, 4}),
		eager.ToSlice[int](),
	)
	// []int{0, 1, 2, 3, 4}
	assert.Equal(t, []int{0, 1, 2, 3, 4}, xs)
}

func TestExampleRedirect(t *testing.T) {
	// import "strconv"
	// import "github.com/shunkeen/eager/datatype/tuple"
	tuple2, _ := eager.Run3(
		eager.FromSlice([]string{"1", "a", "2"}),
		eager.TryMap(strconv.Atoi),
		eager.Redirect(
			eager.ToSlice[int](),
			eager.ToSlice[error](),
		),
	)
	xs, es := tuple2()
	// []int{1, 2}, []error{*errors.errorString {s: "invalid syntax"}}
	assert.Equal(t, []int{1, 2}, xs)
	assert.Equal(t, 1, len(es))
	assert.EqualError(t, es[0], "strconv.Atoi: parsing \"a\": invalid syntax")
}

func TestExampleFirst(t *testing.T) {
	x, _ := eager.Run2(
		eager.FromSlice([]int{1, 2, 3}),
		eager.First[int](),
	)
	// 1
	assert.Equal(t, 1, x)
}

func TestExampleLast(t *testing.T) {
	x, _ := eager.Run2(
		eager.FromSlice([]int{1, 2, 3}),
		eager.Last[int](),
	)
	// 1
	assert.Equal(t, 3, x)
}

func TestExampleNth(t *testing.T) {
	x, _ := eager.Run2(
		eager.FromSlice([]int{1, 2, 3}),
		eager.Nth[int](1),
	)
	// 1
	assert.Equal(t, 2, x)
}

func TestExampleIncludes(t *testing.T) {
	x, _ := eager.Run2(
		eager.FromSlice([]int{1, 2, 3, 4, 5}),
		eager.Includes(3),
	)
	// true
	assert.Equal(t, true, x)
}

func TestExampleMax(t *testing.T) {
	x, _ := eager.Run2(
		eager.FromSlice([]int{3, 4, 2, 0, 1}),
		eager.Max[int](),
	)
	// 4
	assert.Equal(t, 4, x)
}

func TestExampleMin(t *testing.T) {
	x, _ := eager.Run2(
		eager.FromSlice([]int{3, 4, 2, 0, 1}),
		eager.Min[int](),
	)
	// 0
	assert.Equal(t, 0, x)
}

func TestExampleSum(t *testing.T) {
	x, _ := eager.Run2(
		eager.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		eager.Sum(),
	)
	// int(55)
	assert.Equal(t, int(55), x)

	y, _ := eager.Run3(
		eager.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		eager.Map(func(x int) int64 { return int64(x) }),
		eager.SumInteger[int64](),
	)
	// int64(55)
	assert.Equal(t, int64(55), y)

	z, _ := eager.Run2(
		eager.FromSlice([]float64{4.1, 2.0, 1.7}),
		eager.SumFloat[float64](),
	)
	// float64(7.8)
	assert.Equal(t, float64(7.8), z)
}

func TestExampleProduct(t *testing.T) {
	x, _ := eager.Run2(
		eager.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		eager.Product(),
	)
	// int(3628800)
	assert.Equal(t, int(3628800), x)

	y, _ := eager.Run3(
		eager.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		eager.Map(func(x int) int64 { return int64(x) }),
		eager.ProductInteger[int64](),
	)
	// int64(3628800)
	assert.Equal(t, int64(3628800), y)

	z, _ := eager.Run2(
		eager.FromSlice([]float64{4.1, 2.0, 1.7}),
		eager.ProductFloat[float64](),
	)
	// float64(13.939999999999998)
	assert.Equal(t, float64(13.939999999999998), z)
}

func TestExampleIsEmpty(t *testing.T) {
	x, _ := eager.Run2(eager.FromSlice([]int{}), eager.IsEmpty[int]())
	// true
	assert.Equal(t, true, x)

	y, _ := eager.Run2(eager.FromSlice([]int{1}), eager.IsEmpty[int]())
	// false
	assert.Equal(t, false, y)
}

func TestExampleCount(t *testing.T) {
	x, _ := eager.Run2(
		eager.FromSlice([]string{"a", "b", "c"}),
		eager.Count[string](),
	)
	// 3
	assert.Equal(t, 3, x)
}

func TestExampleAnd(t *testing.T) {
	x, _ := eager.Run2(
		eager.FromSlice([]bool{true, false, true}),
		eager.And(),
	)
	// false
	assert.Equal(t, false, x)
}

func TestExampleOr(t *testing.T) {
	x, _ := eager.Run2(
		eager.FromSlice([]bool{true, false, true}),
		eager.Or(),
	)
	// true
	assert.Equal(t, true, x)
}

func TestExampleAll(t *testing.T) {
	x, _ := eager.Run2(
		eager.FromSlice([]int{3, 1, 4}),
		eager.All(func(x int) bool { return x > 2 }),
	)
	// false
	assert.Equal(t, false, x)
}

func TestExampleAny(t *testing.T) {
	x, _ := eager.Run2(
		eager.FromSlice([]int{3, 1, 4}),
		eager.Any(func(x int) bool { return x > 2 }),
	)
	// true
	assert.Equal(t, true, x)
}

func TestExampleForEach(t *testing.T) {
	var acc string
	eager.Run2(
		eager.FromSlice([]string{"a", "b", "c"}),
		eager.ForEach(func(x string) { acc += x }),
	)
	// "abc"
	assert.Equal(t, "abc", acc)
}

func TestExampleReduce(t *testing.T) {
	x, _ := eager.Run2(
		eager.FromSlice([]string{"a", "b", "c", "d"}),
		eager.Reduce("foo", func(x, y string) string { return x + y }),
	)
	// "fooabcd"
	assert.Equal(t, "fooabcd", x)
}

func TestExampleReduce1(t *testing.T) {
	x, _ := eager.Run2(
		eager.FromSlice([]int{1, 2, 3, 4}),
		eager.Reduce1(func(x, y int) int { return x - y }),
	)
	// -8
	assert.Equal(t, -8, x)
}

func TestExampleSparkWith(t *testing.T) {
	x, _ := eager.Run2(
		eager.FromSlice([]int{1, 2, 3, 4}),
		eager.SparkWith(
			eager.Sum(),
			eager.Count[int](),
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
	tuple, _ := eager.Run2(
		eager.FromSlice([]int{0, 1, 2, 3, 4}),
		eager.Spark(
			eager.Sum(),
			eager.ChainCS(
				eager.Drop[int](2),
				eager.Product(),
			),
		),
	)
	x, y := tuple()
	// 10, 24
	assert.Equal(t, 10, x)
	assert.Equal(t, 24, y)
}
