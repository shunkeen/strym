package rethrow

import (
	"github.com/shunkeen/strym/machine/prosumer"
	"github.com/shunkeen/strym/machine/prosumer/naive"
)

func BreakIfErr() prosumer.Rethrow {
	f := func(_ error) bool { return false }
	return naive.TakeWhile(f)
}
