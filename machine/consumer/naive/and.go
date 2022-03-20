package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/consumer"
)

type and struct {
	consumer.Base
	and bool
}

func And() consumer.Naive[bool, bool] {
	return &and{
		and:  true,
		Base: consumer.NewBase(),
	}
}

func (m *and) Await(x bool) {
	if !x {
		m.BaseGoTo = machine.GoToReturn
		m.and = false
	}
}

func (m *and) Return() bool {
	return m.and
}
