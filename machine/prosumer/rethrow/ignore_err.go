package rethrow

import (
	"github.com/shunkeen/strym/machine/prosumer"
	"github.com/shunkeen/strym/machine/prosumer/naive"
)

func IgnoreErr() prosumer.Rethrow {
	f := func(_ error) bool { return false }
	return naive.Filter(f)
}
