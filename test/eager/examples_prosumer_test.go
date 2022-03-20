package eager_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/shunkeen/strym/eager"
	"github.com/stretchr/testify/assert"
)

func TestExampleMap(t *testing.T) {
	xs, _ := eager.Run3(
		eager.FromSlice([]int{1, 2, 3}),
		eager.Map(func(x int) int { return x + 1 }),
		eager.ToSlice[int](),
	)
	// []int{2, 3, 4}
	assert.Equal(t, []int{2, 3, 4}, xs)
}

func TestExampleTryMap(t *testing.T) {
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

func TestExampleFilter(t *testing.T) {
	xs, _ := eager.Run3(
		eager.FromSlice([]int{1, 2, 3}),
		eager.Filter(func(x int) bool { return (x % 2) != 0 }),
		eager.ToSlice[int](),
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

	tuple2, _ := eager.Run3(
		eager.FromSlice([]int{1, 2, 3}),
		eager.TryFilter(f),
		eager.Redirect(
			eager.ToSlice[int](),
			eager.ToSlice[error](),
		),
	)
	xs, es := tuple2()
	// []int{1, 3}, []error{*errors.errorString {s: "even"}}
	assert.Equal(t, []int{1, 3}, xs)
	assert.Equal(t, []error{errors.New("even")}, es)
}

func TestExampleIgnoreErr(t *testing.T) {
	// import "strconv"
	xs, _ := eager.Run4(
		eager.FromSlice([]string{"1", "a", "2"}),
		eager.TryMap(strconv.Atoi),
		eager.IgnoreErr[int](),
		eager.ToSlice[int](),
	)
	// []int{1, 2}
	assert.Equal(t, []int{1, 2}, xs)
}

func TestExampleReverse(t *testing.T) {
	xs, _ := eager.Run3(
		eager.FromSlice([]int{2, 5, 7}),
		eager.Reverse[int](),
		eager.ToSlice[int](),
	)
	// []int{7, 5, 2}
	assert.Equal(t, []int{7, 5, 2}, xs)
}

func TestExampleTake(t *testing.T) {
	xs, _ := eager.Run3(
		eager.FromSlice([]int{1, 2, 3, 4, 5}),
		eager.Take[int](3),
		eager.ToSlice[int](),
	)
	// []int{1, 2, 3}
	assert.Equal(t, []int{1, 2, 3}, xs)
}

func TestExampleDrop(t *testing.T) {
	xs, _ := eager.Run3(
		eager.FromSlice([]int{1, 2, 3, 4, 5}),
		eager.Drop[int](3),
		eager.ToSlice[int](),
	)
	// []int{4, 5}
	assert.Equal(t, []int{4, 5}, xs)
}

func TestExampleTakeWhile(t *testing.T) {
	xs, _ := eager.Run3(
		eager.FromSlice([]int{1, 2, 3, 4, 1, 2, 3, 4}),
		eager.TakeWhile(func(x int) bool { return x < 3 }),
		eager.ToSlice[int](),
	)
	// []int{1, 2}
	assert.Equal(t, []int{1, 2}, xs)
}

func TestExampleDropWhile(t *testing.T) {
	xs, _ := eager.Run3(
		eager.FromSlice([]int{1, 2, 3, 4, 5, 1, 2, 3}),
		eager.DropWhile(func(x int) bool { return x < 3 }),
		eager.ToSlice[int](),
	)
	// []int{3, 4, 5, 1, 2, 3}
	assert.Equal(t, []int{3, 4, 5, 1, 2, 3}, xs)
}

func TestExampleBreakIfErr(t *testing.T) {
	// import "strconv"
	xs, _ := eager.Run4(
		eager.FromSlice([]string{"1", "a", "2"}),
		eager.TryMap(strconv.Atoi),
		eager.BreakIfErr[int](),
		eager.ToSlice[int](),
	)
	// []int{1}
	assert.Equal(t, []int{1}, xs)
}

func TestExampleConcat(t *testing.T) {
	xs, _ := eager.Run4(
		eager.FromSlice([][]int{{1, 2, 3}, {4, 5}, {6}, {}}),
		eager.Map(eager.FromSlice[int]),
		eager.Concat[int](),
		eager.ToSlice[int](),
	)
	// []int{1, 2, 3, 4, 5, 6}
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, xs)
}

func TestExampleConcatMap(t *testing.T) {
	xs, _ := eager.Run3(
		eager.Range(4),
		eager.ConcatMap(eager.Range),
		eager.ToSlice[int](),
	)
	// []int{0, 0, 1, 0, 1, 2}
	assert.Equal(t, []int{0, 0, 1, 0, 1, 2}, xs)
}

func TestExampleScan(t *testing.T) {
	xs, _ := eager.Run3(
		eager.FromSlice([]int{1, 2, 3, 4}),
		eager.Scan(0, func(x, y int) int { return x + y }),
		eager.ToSlice[int](),
	)
	// []int{0, 1, 3, 6, 10}
	assert.Equal(t, []int{0, 1, 3, 6, 10}, xs)
}

func TestExampleScan1(t *testing.T) {
	xs, _ := eager.Run3(
		eager.FromSlice([]int{1, 2, 3, 4}),
		eager.Scan1(func(x, y int) int { return x + y }),
		eager.ToSlice[int](),
	)
	// []int{1, 3, 6, 10}
	assert.Equal(t, []int{1, 3, 6, 10}, xs)
}
