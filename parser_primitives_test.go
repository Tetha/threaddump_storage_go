package input

import "testing"

func TestMatchString(t *testing.T) {
	parser := CreateInput("someString")
	matched := parser.MatchWord("some")
	if !matched {
		t.Error("MatchWord failed")
	}

	if parser.Current() != 'S' {
		t.Errorf("MatchWord should advance input to <S> but was %q", parser.Current())
	}
}
