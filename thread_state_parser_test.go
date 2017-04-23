package input

import(
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
}
