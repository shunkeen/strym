package test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/shunkeen/strym/lazy"
	"github.com/stretchr/testify/assert"
)

func TestExampleUsage(t *testing.T) {
	xs := []string{"2", "a", "3", "b"}

	y, _ := lazy.Run4(
		lazy.FromSlice(xs),
		lazy.TryMap(strconv.Atoi),
		lazy.IgnoreErr[int](),
		lazy.Sum(),
	)

	fmt.Println(y)
	assert.Equal(t, 5, y)
}
