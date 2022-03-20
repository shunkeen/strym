package lazy_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/shunkeen/strym/lazy"
	"github.com/stretchr/testify/assert"
)

func TestExampleMap(t *testing.T) {
	xs, _ := lazy.Run3(
		lazy.FromSlice([]int{1, 2, 3}),
		lazy.Map(func(x int) int { return x + 1 }),
		lazy.ToSlice[int](),
	)
	// []int{2, 3, 4}
	assert.Equal(t, []int{2, 3, 4}, xs)
}

func TestExampleTryMap(t *testing.T) {
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

func TestExampleFilter(t *testing.T) {
	xs, _ := lazy.Run3(
		lazy.FromSlice([]int{1, 2, 3}),
		lazy.Filter(func(x int) bool { return (x % 2) != 0 }),
		lazy.ToSlice[int](),
	)
	// []int{1, 3}
	assert.Equal(t, []int{1, 3}, xs)
}

func TestExampleTryFilter(t *testing.T) {
	// import "errors"
	// import "github.com/shunkeen/eager/datatype/tuple"
	f := func(x int) error {
		if (x % 2) == 0 {
			return errors.New("even")
		}
		return nil
	}

	tuple2, _ := lazy.Run3(
		lazy.FromSlice([]int{1, 2, 3}),
		lazy.TryFilter(f),
		lazy.Redirect(
			lazy.ToSlice[int](),
			lazy.ToSlice[error](),
		),
	)
	xs, es := tuple2()
	// []int{1, 3}, []error{*errors.errorString {s: "even"}}
	assert.Equal(t, []int{1, 3}, xs)
	assert.Equal(t, []error{errors.New("even")}, es)
}

func TestExampleIgnoreErr(t *testing.T) {
	// import "strconv"
	xs, _ := lazy.Run4(
		lazy.FromSlice([]string{"1", "a", "2"}),
		lazy.TryMap(strconv.Atoi),
		lazy.IgnoreErr[int](),
		lazy.ToSlice[int](),
	)
	// []int{1, 2}
	assert.Equal(t, []int{1, 2}, xs)
}

func TestExampleReverse(t *testing.T) {
	xs, _ := lazy.Run3(
		lazy.FromSlice([]int{2, 5, 7}),
		lazy.Reverse[int](),
		lazy.ToSlice[int](),
	)
	// []int{7, 5, 2}
	assert.Equal(t, []int{7, 5, 2}, xs)
}

func TestExampleTake(t *testing.T) {
	xs, _ := lazy.Run3(
		lazy.FromSlice([]int{1, 2, 3, 4, 5}),
		lazy.Take[int](3),
		lazy.ToSlice[int](),
	)
	// []int{1, 2, 3}
	assert.Equal(t, []int{1, 2, 3}, xs)
}

func TestExampleDrop(t *testing.T) {
	xs, _ := lazy.Run3(
		lazy.FromSlice([]int{1, 2, 3, 4, 5}),
		lazy.Drop[int](3),
		lazy.ToSlice[int](),
	)
	// []int{4, 5}
	assert.Equal(t, []int{4, 5}, xs)
}

func TestExampleTakeWhile(t *testing.T) {
	xs, _ := lazy.Run3(
		lazy.FromSlice([]int{1, 2, 3, 4, 1, 2, 3, 4}),
		lazy.TakeWhile(func(x int) bool { return x < 3 }),
		lazy.ToSlice[int](),
	)
	// []int{1, 2}
	assert.Equal(t, []int{1, 2}, xs)
}

func TestExampleDropWhile(t *testing.T) {
	xs, _ := lazy.Run3(
		lazy.FromSlice([]int{1, 2, 3, 4, 5, 1, 2, 3}),
		lazy.DropWhile(func(x int) bool { return x < 3 }),
		lazy.ToSlice[int](),
	)
	// []int{3, 4, 5, 1, 2, 3}
	assert.Equal(t, []int{3, 4, 5, 1, 2, 3}, xs)
}

func TestExampleBreakIfErr(t *testing.T) {
	// import "strconv"
	xs, _ := lazy.Run4(
		lazy.FromSlice([]string{"1", "a", "2"}),
		lazy.TryMap(strconv.Atoi),
		lazy.BreakIfErr[int](),
		lazy.ToSlice[int](),
	)
	// []int{1}
	assert.Equal(t, []int{1}, xs)
}

func TestExampleConcat(t *testing.T) {
	xs, _ := lazy.Run4(
		lazy.FromSlice([][]int{{1, 2, 3}, {4, 5}, {6}, {}}),
		lazy.Map(lazy.FromSlice[int]),
		lazy.Concat[int](),
		lazy.ToSlice[int](),
	)
	// []int{1, 2, 3, 4, 5, 6}
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, xs)
}

func TestExampleConcatMap(t *testing.T) {
	xs, _ := lazy.Run3(
		lazy.Range(4),
		lazy.ConcatMap(lazy.Range),
		lazy.ToSlice[int](),
	)
	// []int{0, 0, 1, 0, 1, 2}
	assert.Equal(t, []int{0, 0, 1, 0, 1, 2}, xs)
}

func TestExampleScan(t *testing.T) {
	xs, _ := lazy.Run3(
		lazy.FromSlice([]int{1, 2, 3, 4}),
		lazy.Scan(0, func(x, y int) int { return x + y }),
		lazy.ToSlice[int](),
	)
	// []int{0, 1, 3, 6, 10}
	assert.Equal(t, []int{0, 1, 3, 6, 10}, xs)
}

func TestExampleScan1(t *testing.T) {
	xs, _ := lazy.Run3(
		lazy.FromSlice([]int{1, 2, 3, 4}),
		lazy.Scan1(func(x, y int) int { return x + y }),
		lazy.ToSlice[int](),
	)
	// []int{1, 3, 6, 10}
	assert.Equal(t, []int{1, 3, 6, 10}, xs)
}
