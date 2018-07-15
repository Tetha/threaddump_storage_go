package input

import "testing"

func TestReadUntilHappyCase(t *testing.T) {
	parser := CreateInput("foo(bar)")
	parsed, word := parser.ReadUntil('(')
	if !parsed {
		t.Error("ReadUntil should have succeeded")
	}
	if word != "foo" {
		t.Errorf("ReadUntil should have returned the word <foo> but returned <%s>", word)
	}
	if parser.Current() != '(' {
		t.Errorf("ReadUntil should place the input on the character to read until, but placed it on %q", parser.Current())
	}
}

func TestReadUntilWithMissingStopChar(t *testing.T) {
	parser := CreateInput("someString")
	parsed, _ := parser.ReadUntil('(')
	if parsed {
		t.Error("ReadUntil should not parse without the end character")
	}
	if parser.Current() != 's' {
		t.Errorf("ReadUntil should not advance the input, but advanced to %q", parser.Current())
	}
}
