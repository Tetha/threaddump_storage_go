package input

import "testing"

var matchStringTests = map[string]struct {
	input           string
	matchedWord     string
	expectedMatch   bool
	expectedCurrent byte
}{
	"happy case":      {"someString", "some", true, 'S'},
	"failure in word": {"someString", "somd", false, 's'},
	"end of string":   {"someString", "someStringAndThenSome", false, 's'},
	"uncovered line":  {"poof", "foop", false, 'p'},
}

func TestMatchString(t *testing.T) {
	for name, tt := range matchStringTests {
		parser := CreateInput(tt.input)
		matched := parser.MatchWord(tt.matchedWord)

		if matched != tt.expectedMatch {
			t.Errorf("%s: Expected MatchWord to return <%v>, got: <%v>", name, tt.expectedMatch, matched)
		}
		if parser.Current() != tt.expectedCurrent {
			t.Errorf("%s: MatchWord should advance input to <%v> but was %q", name, tt.expectedCurrent, parser.Current())
		}
	}
}
