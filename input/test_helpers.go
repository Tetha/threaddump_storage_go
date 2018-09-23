package input

import "testing"

type lineParserTest struct {
	input string

	shouldParse     bool
	output          StacktraceLine
	expectedCurrent byte
}

func runParserTestcase(t *testing.T, idx int, tt lineParserTest, parseMethod func(*Input) (bool, StacktraceLine)) {
	parser := CreateInput(tt.input)
	parsed, line := parseMethod(&parser)
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
