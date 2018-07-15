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

func TestMatchStringFailsInWord(t *testing.T) {
	parser := CreateInput("someString")
	matched := parser.MatchWord("somd")
	if matched {
		t.Error("MatchWord didn't fail")
	}

	if parser.Current() != 's' {
		t.Errorf("MatchWord didn't roll back properly, and got stuck on %q", parser.Current())
	}
}

func TestMatchStringFailsAtEndOfString(t *testing.T) {
	parser := CreateInput("someString")
	matched := parser.MatchWord("someStringAndThenSome")
	if matched {
		t.Error("MatchWord didn't fail")
	}

	if parser.Current() != 's' {
		t.Errorf("MatchWord didn't roll back properly, and got stuck on %q", parser.Current())
	}
}
