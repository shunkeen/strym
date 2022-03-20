package producer

import "github.com/shunkeen/strym/machine"

type Base struct{}

func (m *Base) Await(_ machine.Void) {}

func (m *Base) Return() (machine.Void, error) {
	return machine.Void{}, nil
}

func (m *Base) DontWait() {}

func (m *Base) Continue() {}

func (m *Base) Defer() {}
