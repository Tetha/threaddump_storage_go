package input

import "testing"

var parseBlockedLineTests = []struct {
	input string

	shouldParse     bool
	output          StacktraceLine
	expectedCurrent byte
}{
	{
		"\t- parking to wait for <0x000000065be92a68> (a java.util.concurrent.locks.AbstractQueuedSynchronizer$ConditionObject)\n$",
		true,
		StacktraceLine{
			Type:        BlockedLine,
			LockAddress: "0x000000065be92a68",
			Class:       "java.util.concurrent.locks.AbstractQueuedSynchronizer$ConditionObject",
		},
		'$',
	},
	{
		"\t- waiting to lock <0x00000000e0c5a010> (a org.apache.logging.log4j.core.appender.FileManager)\n$",
		true,
		StacktraceLine{
			Type:        BlockedLine,
			LockAddress: "0x00000000e0c5a010",
			Class:       "org.apache.logging.log4j.core.appender.FileManager",
		},
		'$',
	},
}

func TestParseBlockedLine(t *testing.T) {
	for idx, tt := range parseBlockedLineTests {
		parser := CreateInput(tt.input)
		parsed, line := parser.parseBlockedLine()
		if parsed != tt.shouldParse {
			t.Errorf("%d: expected parsed to be <%v>, was <%v>", idx, tt.shouldParse, parsed)
		}

		if tt.shouldParse && line != tt.output {
			t.Errorf("%d: output incorrect. Expected: <%v>, got: <%v>", idx, tt.output, line)
		}

		if parser.Current() != tt.expectedCurrent {
			t.Errorf("%d: input not advanced correctly: expected: <%v>, was <%v>", idx, tt.expectedCurrent, parser.Current())
		}
	}
}
