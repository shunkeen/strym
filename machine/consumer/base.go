package consumer

import "github.com/shunkeen/strym/machine"

type Base struct {
	BaseGoTo machine.GoTo
}

func NewBase() Base {
	return Base{BaseGoTo: machine.GoToAwait}
}

func (m *Base) GoTo() machine.GoTo {
	return m.BaseGoTo
}

func (m *Base) DontWait() {
	m.BaseGoTo = machine.GoToReturn
}

func (m *Base) Yield() (machine.Void, error) {
	return machine.Void{}, nil
}

func (m *Base) Continue() {}

func (m *Base) Defer() {}
