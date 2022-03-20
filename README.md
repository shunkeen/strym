strym
=====
​
Strym is a lazy stream library for Go.
​
* Go 1.18+ Generics
* Lazy Stream
* with Error Stream

​
## Install
```
go get github.com/shunkeen/strym
```

​
## Usage
```go
package main
​
import (
    "fmt"
    "strconv"
​
    "github.com/shunkeen/strym/lazy"
)
​
func main() {
	xs := []string{"2", "a", "3", "b"}

	y, _ := lazy.Run4(
		lazy.FromSlice(xs),
		lazy.TryMap(strconv.Atoi),
		lazy.IgnoreErr[int](),
		lazy.Sum,
	)
​
    fmt.Println(y) // 5
}
```
​

## Spec

* Producer
    * FromSlice
    * Iterate
    * Repeat
    * CycleSlice
    * Replicate
    * Flatten
    * FlatMap
    * Range
    * RangeTo
    * RangeBy
    * ZipWith
    * ZipWith3
    * Zip
    * Zip3
    * Once
    * ThrowOnce
* Prosumer
    * Map
    * TryMap
    * Filter
    * TryFilter
    * IgnoreErr
    * Reverse
    * Take
    * Drop
    * TakeWhile
    * DropWhile
    * BreakIfErr
    * Concat
    * ConcatMap
    * Scan
    * Scan1
* Consumer
    * ToSlice
    * Redirect
    * First
    * Last
    * Nth
    * Includes
    * Max
    * Min
    * Sum
    * Product
    * IsEmpty
    * Count
    * And
    * Or
    * All
    * Any
    * ForEach
    * Reduce
    * Reduce1
    * SparkWith
    * Spark
* Other
    * Run1 -- Run12
    * Chain2 -- Chain10
    * ChainHM
    * ChainPD
    * ChainCS


## Examples
​
### Producer

#### FromSlice
```go
// import "github.com/shunkeen/strym/lazy"
xs, _ := lazy.Run2(
    lazy.FromSlice([]int{0, 1, 2, 3, 4}),
    lazy.ToSlice[int](),
)
// []int{0, 1, 2, 3, 4}
```

#### Iterate
```go
xs, _ := lazy.Run3(
    lazy.Iterate(true, func(x bool) bool { return !x }),
    lazy.Take[bool](5),
    lazy.ToSlice[bool](),
)
// []bool{true, false, true, false, true}
```

#### Repeat
```go
xs, _ := lazy.Run3(
    lazy.Repeat(17),
    lazy.Take[int](5),
    lazy.ToSlice[int](),
)
// []int{17, 17, 17, 17, 17}
```

#### Repeat
```go
xs, _ := lazy.Run3(
    lazy.CycleSlice([]int{2, 5, 7}),
    lazy.Take[int](5),
    lazy.ToSlice[int](),
)
// []int{2, 5, 7, 2, 5}
```

#### Replicate
```go
xs, _ := lazy.Run2(
    lazy.Replicate(4, true),
    lazy.ToSlice[bool](),
)
// []bool{true, true, true, true}
```

#### Flatten
```go
xs, _ := lazy.Run2(
    lazy.Flatten(lazy.ChainPD(
        lazy.FromSlice([]int{0, 1, 2, 3}),
        lazy.Map(lazy.Range),
    )),
    lazy.ToSlice[int](),
)
// []int{0, 0, 1, 0, 1, 2}
```

#### FlatMap
```go
xs, _ := lazy.Run2(
    lazy.FlatMap(
        lazy.FromSlice([]int{0, 1, 2, 3}),
        lazy.Range,
    ),
    lazy.ToSlice[int](),
)
// []int{0, 0, 1, 0, 1, 2}
```

#### Range
```go
xs, _ := lazy.Run2(
    lazy.Range(5),
    lazy.ToSlice[int](),
)
// []int{0, 1, 2, 3, 4}

ys, _ := lazy.Run2(
    lazy.RangeTo(5, 10),
    lazy.ToSlice[int](),
)
// []int{5, 6, 7, 8, 9}

zs, _ := lazy.Run2(
    lazy.RangeBy(0, 10, 2),
    lazy.ToSlice[int](),
)
// []int{0, 2, 4, 6, 8}
```

```go
xs, _ := lazy.Run2(
    lazy.RangeBy(8, 0, -2),
    lazy.ToSlice[int](),
)
// []int{8, 6, 4, 2}

ys, _ := lazy.Run2(
    lazy.RangeInteger[int64](5),
    lazy.ToSlice[int64](),
)
// []int64{0, 1, 2, 3, 4}

zs, _ := lazy.Run2(
    lazy.RangeFloatBy(1.0, 0.5, -0.1),
    lazy.ToSlice[float64](),
)
// []float64{1.0, 0.9, 0.8, 0.7, 0.6}
```

#### ZipWith
​
```go
xs, _ := lazy.Run2(
    lazy.ZipWith(
        lazy.FromSlice([]int{1, 2, 3}),
        lazy.FromSlice([]int{4, 5, 6}),
        func(x, y int) int { return x + y },
    ),
    lazy.ToSlice[int](),
)
// []int{5, 7, 9}

ys, _ := lazy.Run2(
    lazy.ZipWith3(
        lazy.FromSlice([]int{7, 8, 9}),
        lazy.FromSlice([]int{4, 5, 6}),
        lazy.FromSlice([]int{1, 2, 3}),
        func(x, y, z int) int { return x - y - z },
    ),
    lazy.ToSlice[int](),
)
// []int{2, 1, 0}
```


#### Zip

```go
// import "github.com/shunkeen/eager/datatype/tuple"
xs, _ := lazy.Run2(
    lazy.Zip(
        lazy.FromSlice([]int{1, 2}),
        lazy.FromSlice([]string{"A", "B"}),
    ),
    lazy.ToSlice[tuple.Tuple2[int, string]](),
)
x1, x2 := xs[0]()
// 1, "A"

x1, x2 = xs[1]()
// 2, "B"

ys, _ := lazy.Run2(
    lazy.Zip3(
        lazy.FromSlice([]int{1, 2}),
        lazy.FromSlice([]string{"A", "B"}),
        lazy.FromSlice([]bool{true, false}),
    ),
    lazy.ToSlice[tuple.Tuple3[int, string, bool]](),
)
x1, x2, x3 := ys[0]()
// 1, "A", true

x1, x2, x3 = ys[1]()
// 2, "B", false
```

#### Once

```go
xs, _ := lazy.Run2(
    lazy.Once(999),
    lazy.ToSlice[int](),
)
// []int{999}
```


#### OnceThrow

```go
// import "errors"
_, err := lazy.Run2(
    lazy.ThrowOnce[int](errors.New("once")),
    lazy.ToSlice[int](),
)
// *errors.errorString {s: "once"}
```

### Prosumer

#### Map
```go
xs, _ := lazy.Run3(
    lazy.FromSlice([]int{1, 2, 3}),
    lazy.Map(func(x int) int { return x + 1 }),
    lazy.ToSlice[int](),
)
// []int{2, 3, 4}
```

#### TryMap
```go
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
```

#### Filter
```go
xs, _ := lazy.Run3(
    lazy.FromSlice([]int{1, 2, 3}),
    lazy.Filter(func(x int) bool { return (x % 2) != 0 }),
    lazy.ToSlice[int](),
)
// []int{1, 3}
```

#### TryFilter
```go
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
```

#### IgnoreErr
```go
// import "strconv"
xs, _ := lazy.Run4(
    lazy.FromSlice([]string{"1", "a", "2"}),
    lazy.TryMap(strconv.Atoi),
    lazy.IgnoreErr[int](),
    lazy.ToSlice[int](),
)
// []int{1, 2}
```

#### Reverse
```go
xs, _ := lazy.Run3(
    lazy.FromSlice([]int{2, 5, 7}),
    lazy.Reverse[int](),
    lazy.ToSlice[int](),
)
// []int{7, 5, 2}
```

#### Take
```go
xs, _ := lazy.Run3(
    lazy.FromSlice([]int{1, 2, 3, 4, 5}),
    lazy.Take[int](3),
    lazy.ToSlice[int](),
)
// []int{1, 2, 3}
```

#### Drop
```go
xs, _ := lazy.Run3(
    lazy.FromSlice([]int{1, 2, 3, 4, 5}),
    lazy.Drop[int](3),
    lazy.ToSlice[int](),
)
// []int{4, 5}
```

#### TakeWhile
```go
xs, _ := lazy.Run3(
    lazy.FromSlice([]int{1, 2, 3, 4, 1, 2, 3, 4}),
    lazy.TakeWhile(func(x int) bool { return x < 3 }),
    lazy.ToSlice[int](),
)
// []int{1, 2}
```

#### TakeWhile
```go
xs, _ := lazy.Run3(
    lazy.FromSlice([]int{1, 2, 3, 4, 5, 1, 2, 3}),
    lazy.DropWhile(func(x int) bool { return x < 3 }),
    lazy.ToSlice[int](),
)
// []int{3, 4, 5, 1, 2, 3}
```

### BreakIfErr
```go
// import "strconv"
xs, _ := lazy.Run4(
    lazy.FromSlice([]string{"1", "a", "2"}),
    lazy.TryMap(strconv.Atoi),
    lazy.BreakIfErr[int](),
    lazy.ToSlice[int](),
)
// []int{1}
```

### Concat
```go
xs, _ := lazy.Run4(
    lazy.FromSlice([][]int{{1, 2, 3}, {4, 5}, {6}, {}}),
    lazy.Map(lazy.FromSlice[int]),
    lazy.Concat[int](),
    lazy.ToSlice[int](),
)
// []int{1, 2, 3, 4, 5, 6}
```

### ConcatMap
```go
xs, _ := lazy.Run3(
    lazy.Range(4),
    lazy.ConcatMap(lazy.Range),
    lazy.ToSlice[int](),
)
// []int{0, 0, 1, 0, 1, 2}
```

### Scan
```go
xs, _ := lazy.Run3(
    lazy.FromSlice([]int{1, 2, 3, 4}),
    lazy.Scan(0, func(x, y int) int { return x + y }),
    lazy.ToSlice[int](),
)
// []int{0, 1, 3, 6, 10}
```

### Scan1
```go
xs, _ := lazy.Run3(
    lazy.FromSlice([]int{1, 2, 3, 4}),
    lazy.Scan1(func(x, y int) int { return x + y }),
    lazy.ToSlice[int](),
)
// []int{1, 3, 6, 10}
```

​
### Consumer

#### ToSlice
```go
xs, _ := lazy.Run2(
    lazy.FromSlice([]int{0, 1, 2, 3, 4}),
    lazy.ToSlice[int](),
)
// []int{0, 1, 2, 3, 4}
```


#### Redirect
```go
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
```

#### First
```go
x, _ := lazy.Run2(
    lazy.FromSlice([]int{1, 2, 3}),
    lazy.First[int](),
)
// 1
```

#### Last
```go
x, _ := lazy.Run2(
    lazy.FromSlice([]int{1, 2, 3}),
    lazy.Last[int](),
)
// 3
```

#### Nth
```go
x, _ := lazy.Run2(
    lazy.FromSlice([]int{1, 2, 3}),
    lazy.Nth[int](1),
)
// 2
```

#### Includes
```go
x, _ := lazy.Run2(
    lazy.FromSlice([]int{1, 2, 3, 4, 5}),
    lazy.Includes(3),
)
// true
```

#### Max
```go
x, _ := lazy.Run2(
    lazy.FromSlice([]int{3, 4, 2, 0, 1}),
    lazy.Max[int](),
)
// 4
```

#### Min
```go
x, _ := lazy.Run2(
    lazy.FromSlice([]int{3, 4, 2, 0, 1}),
    lazy.Min[int](),
)
// 0
```

#### Sum
```go
x, _ := lazy.Run2(
    lazy.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
    lazy.Sum(),
)
// int(55)

y, _ := lazy.Run3(
    lazy.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
    lazy.Map(func(x int) int64 { return int64(x) }),
    lazy.SumInteger[int64](),
)
// int64(55)

z, _ := lazy.Run2(
    lazy.FromSlice([]float64{4.1, 2.0, 1.7}),
    lazy.SumFloat[float64](),
)
// float64(7.8)
```

#### Product
```go
x, _ := lazy.Run2(
    lazy.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
    lazy.Product(),
)
// int(3628800)

y, _ := lazy.Run3(
    lazy.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
    lazy.Map(func(x int) int64 { return int64(x) }),
    lazy.ProductInteger[int64](),
)
// int64(3628800)

z, _ := lazy.Run2(
    lazy.FromSlice([]float64{4.1, 2.0, 1.7}),
    lazy.ProductFloat[float64](),
)
// float64(13.939999999999998)
```

#### IsEmpty
```go
x, _ := lazy.Run2(lazy.FromSlice([]int{}), lazy.IsEmpty[int]())
// true

y, _ := lazy.Run2(lazy.FromSlice([]int{1}), lazy.IsEmpty[int]())
// false
```

#### Count
```go
x, _ := lazy.Run2(
    lazy.FromSlice([]string{"a", "b", "c"}),
    lazy.Count[string](),
)
// 3
```

#### And
```go
x, _ := lazy.Run2(
    lazy.FromSlice([]bool{true, false, true}),
    lazy.And(),
)
// false
```

#### Or
```go
x, _ := lazy.Run2(
    lazy.FromSlice([]bool{true, false, true}),
    lazy.Or(),
)
// true
```

#### All
```go
x, _ := lazy.Run2(
    lazy.FromSlice([]int{3, 1, 4}),
    lazy.All(func(x int) bool { return x > 2 }),
)
// false
```

#### Any
```go
x, _ := lazy.Run2(
    lazy.FromSlice([]int{3, 1, 4}),
    lazy.Any(func(x int) bool { return x > 2 }),
)
// true
```

#### ForEach
```go
var acc string
lazy.Run2(
    lazy.FromSlice([]string{"a", "b", "c"}),
    lazy.ForEach(func(x string) { acc += x }),
)
// "abc"
```

#### Reduce
```go
x, _ := lazy.Run2(
    lazy.FromSlice([]string{"a", "b", "c", "d"}),
    lazy.Reduce("foo", func(x, y string) string { return x + y }),
)
// "fooabcd"
```

#### Reduce1
```go
x, _ := lazy.Run2(
    lazy.FromSlice([]int{1, 2, 3, 4}),
    lazy.Reduce1(func(x, y int) int { return x - y }),
)
// -8
```

#### SparkWith
```go
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
```

#### Spark
```go
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
```

### Other

#### Run1 -- Run12
```go
h := lazy.ChainPD(
    lazy.FromSlice([]int{0, 1, 2, 3, 4}),
    lazy.Sum,
)

x, _ := lazy.Run1(h)
// 10
```

```go
x, _ := lazy.Run2(
    lazy.FromSlice([]int{0, 1, 2, 3, 4}),
    lazy.Sum,
)
// 10
```

```go
x, err := lazy.RunN!(
    instanceOfProducer,
    instanceOfProsumer1,
    instanceOfProsumer2,
    ...,
    instanceOfProsumerM, // M = N - 2
    instanceOfConsumer,
)
```

#### Chain2 -- Chain10
```go
newProsumer := lazy.ChainN!(
    instanceOfProsumer1,
    instanceOfProsumer2,
    ...,
    instanceOfProsumerN,
)
```

#### ChainHM
```go
newHermit := lazy.ChainHM(
    instanceOfProducer,
    instanceOfConsumer,
)
```

#### ChainPD
```go
newProducer := lazy.ChainPD(
    instanceOfProducer,
    instanceOfProsumer,
)
```

#### ChainCS
```go
newConsumer := lazy.ChainCS(
    instanceOfProsumer,
    instanceOfConsumer,
)
```

## Contributors
* [shunkeen](https://github.com/shunkeen)


## License
Copyright © 2022 @shunkeen.
This project is MIT licensed.
