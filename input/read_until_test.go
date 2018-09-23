package input

import "testing"

var readUntilTests = map[string]struct {
	input    string
	stopChar byte

	shouldParse bool
	output      string
	current     byte
}{
	"happy case":    {"foo(bar)", '(', true, "foo", '('},
	"missing start": {"someString", '(', false, "", 's'},
}

func TestReadUntil(t *testing.T) {
	for name, tt := range readUntilTests {
		parser := CreateInput(tt.input)
		parsed, word := parser.ReadUntil(tt.stopChar)

		if parsed != tt.shouldParse {
			t.Errorf("%s: Expected parsed to be <%v>, was <%v>", name, tt.shouldParse, parsed)
		}

		if tt.shouldParse && word != tt.output {
			t.Errorf("%s: Expected parsed word to be <%v>, was <%v>", name, tt.output, word)
		}

		if parser.Current() != tt.current {
			t.Errorf("%s: Expected input to be on <%q>, but was <%q>", name, tt.current, parser.Current())
		}
	}
}
