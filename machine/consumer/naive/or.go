package naive

import (
	"github.com/shunkeen/strym/machine"
	"github.com/shunkeen/strym/machine/consumer"
)

type or struct {
	consumer.Base
	or bool
}

func Or() consumer.Naive[bool, bool] {
	m := consumer.NewBase()
	return &or{Base: m}
}

func (m *or) Await(x bool) {
	if x {
		m.BaseGoTo = machine.GoToReturn
		m.or = true
	}
}

func (m *or) Return() bool {
	return m.or
}
