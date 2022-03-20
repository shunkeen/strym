package throw

import (
	"github.com/shunkeen/strym/machine/producer"
	"github.com/shunkeen/strym/machine/producer/naive"
)

func ThrowOnce(err error) producer.Throw {
	return naive.Once(err)
}
