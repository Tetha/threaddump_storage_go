package input

import "testing"

func TestDelimitedWordWithoutStartingDelimiter(t *testing.T) {
	parser := CreateInput("someString")
	parsed, delimitedWord := parser.DelimitedWord('<', '>')
	if parsed {
		t.Errorf("DelimitedWord should not parse, but parsed %s", delimitedWord)
	}
}

func TestDelimitedWordWithDelimitedWord(t *testing.T) {
	parser := CreateInput("<someString>$")
	parsed, delimitedWord := parser.DelimitedWord('<', '>')
	if !parsed {
		t.Error("DelimitedWord didn't parse properly")
	}
	if delimitedWord != "someString" {
		t.Errorf("DelimitedWord should have parsed <someString>, but parsed <%s>", delimitedWord)
	}
	if parser.Current() != '$' {
		t.Errorf("DelimitedWord should advance the input past the closing delimiter, but only advanced to %q", parser.Current())
	}
}
