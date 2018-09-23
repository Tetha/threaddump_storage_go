package input

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var threadStateTests = map[string]struct {
	input string

	shouldParse   bool
	state         string
	clarification string
	current       byte
}{
	"no clarification": {
		"   java.lang.Thread.State: RUNNABLE\n$",
		true,
		"RUNNABLE",
		"",
		'$',
	},
	"clarification": {
		"   java.lang.Thread.State: WAITING (parking)\n$",
		true,
		"WAITING",
		"parking",
		'$',
	},
}

func TestParseThreadState(t *testing.T) {
	for description, tt := range threadStateTests {
		parser := CreateInput(tt.input)
		parsed, state, clarification := parser.parseThreadState()

		assert := assert.New(t)
		assert.Equal(tt.shouldParse, parsed, description)
		assert.Equal(tt.state, state, description)
		assert.Equal(tt.clarification, clarification, description)
		assert.Equal(tt.current, parser.Current(), description)
	}
}
