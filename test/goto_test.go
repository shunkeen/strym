package test

import (
	"testing"

	"github.com/shunkeen/strym/machine"
	"github.com/stretchr/testify/assert"
)

func TestGoTo_String(t *testing.T) {
	assert.Equal(t, "go to await", machine.GoToAwait.String())
	assert.Equal(t, "go to yield", machine.GoToYield.String())
	assert.Equal(t, "go to return", machine.GoToReturn.String())
	assert.Equal(t, "go to continue", machine.GoToContinue.String())
}
