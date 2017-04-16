package input

import "testing"

func TestDelimitedWordWithoutStartingDelimiter(t *testing.T) {
	parser := CreateInput("someString")
	parsed, delimitedWord := parser.DelimitedWord('<', '>')
	if parsed {
		t.Errorf("DelimitedWord should not parse, but parsed %s", delimitedWord)
	}
}
