package input

import "testing"

var delimitedWordTests = map[string]struct {
	input string

	mustParse       bool
	expectedWord    string
	expectedCurrent byte
}{
	"no starting delimiter": {"someString$", false, "", 's'},
	"happy case":            {"<someString>$", true, "someString", '$'},
	"missing delimiter":     {"<someString", false, "", '<'},
}

func TestDelimitedWord(t *testing.T) {
	for name, tt := range delimitedWordTests {
		parser := CreateInput(tt.input)
		parsed, delimitedWord := parser.DelimitedWord('<', '>')
		if parsed != tt.mustParse {
			t.Errorf("%s: Expected return value to be <%v>, got <%v>", name, tt.mustParse, parsed)
		}

		if tt.mustParse && tt.expectedWord != delimitedWord {
			t.Errorf("%s: Expected delimitedWord to parse word <%s>, but got <%s>", name, tt.expectedWord, delimitedWord)
		}

		if parser.Current() != tt.expectedCurrent {
			t.Errorf("%s: Expected delimitedWord to leave the input on <%x>, but it was <%x>", name, tt.expectedCurrent, parser.Current())
		}
	}
}
