package input

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseThreadStateRunning(t *testing.T) {
	parser := CreateInput("   java.lang.Thread.State: RUNNABLE\n$")
	parsed, state, clarification := parser.ParseThreadState()

	assert := assert.New(t)
	assert.True(parsed, "ParseThreadState should succeed on valid input")
	assert.Equal("RUNNABLE", state)
	assert.Equal("", clarification, "Runnable has no clarification")
	assert.Equal(byte('$'), parser.Current(), "ParseThreadState should consume the entire input")
}

func TestParseThreadStateWaiting(t *testing.T) {
	parser := CreateInput("   java.lang.Thread.State: WAITING (parking)\n$")
	parsed, state, clarification := parser.ParseThreadState()

	assert := assert.New(t)
	assert.True(parsed, "ParseThreadState should succeed on valid input")
	assert.Equal("WAITING", state)
	assert.Equal("parking", clarification, "Runnable has no clarification")
	assert.Equal(byte('$'), parser.Current(), "ParseThreadState should consume the entire input")
}
