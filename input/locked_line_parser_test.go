package input

import "testing"

var lockLineTests = []struct {
	input string

	shouldParse bool
	output      StacktraceLine
	current     byte
}{
	{
		"\t- locked <0x00000000e0e97cb0> (a io.netty.channel.nio.SelectedSelectionKeySet)\n$",
		true,
		StacktraceLine{
			Type:        LockedLine,
			LockAddress: "0x00000000e0e97cb0",
			Class:       "io.netty.channel.nio.SelectedSelectionKeySet",
		},
		'$',
	},
}

func TestParseLockLine(t *testing.T) {
	for idx, tt := range lockLineTests {
		parser := CreateInput(tt.input)
		parsed, line := parser.ParseLockedLine()

		if parsed != tt.shouldParse {
			t.Errorf("%d: Expected parsed to be <%v>, got <%v>", idx, tt.shouldParse, parsed)
		}

		if tt.shouldParse && line != tt.output {
			t.Errorf("%d: Expected output to be <%v>, got <%v>", idx, tt.output, line)
		}

		if parser.Current() != tt.current {
			t.Errorf("%d: Expected input to be advanced to <%v>, was: <%v>", idx, tt.current, parser.Current())
		}
	}
}
