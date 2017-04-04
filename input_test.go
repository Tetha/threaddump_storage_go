package input

import "testing"

func TestCreateInput(t *testing.T) {
    result := CreateInput("someString")
    if(result.content != "someString") {
        t.Errorf("Expected content to be <someString>, but was %s", result.content)
    }

    if(result.position != 0) {
        t.Errorf("Expected position to be 0, but was %d", result.position)
    }
}

func TestCurrent(t *testing.T) {
    result := CreateInput("someString")
    if result.Current() != "s"[0] {
        t.Errorf("Expected current char to be <s>, but was %s", result.Current())
    }
}

func TestAdvance(t *testing.T) {
    subject := CreateInput("someString")
    subject.Advance()
    subject.Advance()
    if subject.Current() != "m"[0] {
        t.Errorf("Expect current to be <m>, but was %s", subject.Current())
    }
}
