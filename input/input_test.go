package input

import "testing"

func TestCreateInput(t *testing.T) {
	result := CreateInput("someString")
	if result.content != "someString" {
		t.Errorf("Expected content to be <someString>, but was %s", result.content)
	}

	if result.position != 0 {
		t.Errorf("Expected position to be 0, but was %d", result.position)
	}
}

func TestCurrent(t *testing.T) {
	result := CreateInput("someString")
	if result.Current() != 's' {
		t.Errorf("Expected current char to be <s>, but was %b", result.Current())
	}
}

func TestAdvance(t *testing.T) {
	subject := CreateInput("someString")
	subject.Advance()
	subject.Advance()
	if subject.Current() != 'm' {
		t.Errorf("Expect current to be <m>, but was %b", subject.Current())
	}
}

func TestMarkRollback(t *testing.T) {
	subject := CreateInput("someString")
	subject.Advance()
	subject.Advance()
	if subject.Current() != 'm' {
		t.Errorf("Current is broken")
	}

	subject.Mark()
	subject.Advance()
	subject.Advance()
	if subject.Current() != 'S' {
		t.Errorf("Curent is broken")
	}

	err := subject.Rollback()
	if subject.Current() != 'm' {
		t.Errorf("Expect Rollback to go back to <m> but got %q", subject.Current())
	}
	if err != nil {
		t.Error("Rollback issued an error in a valid situation")
	}
}

func TestInvalidRollback(t *testing.T) {
	subject := CreateInput("someString")
	err := subject.Rollback()
	if err == nil {
		t.Errorf("Rollback didn't fail properly without a previous mark: %s", err)
	}
}

func TestMarkCommit(t *testing.T) {
	subject := CreateInput("someString")
	subject.Advance()
	subject.Advance()
	if subject.Current() != 'm' {
		t.Errorf("Current is broken")
	}

	subject.Mark()
	subject.Advance()
	subject.Advance()
	if subject.Current() != 'S' {
		t.Errorf("Curent is broken")
	}

	err := subject.Commit()
	if subject.Current() != 'S' {
		t.Errorf("Expect Mark to stick to the position, but it rolled back to %b", subject.Current())
	}
	if err != nil {
		t.Error("Mark issued an error in a valid situation")
	}
}

func TestInvalidMark(t *testing.T) {
	subject := CreateInput("someString")
	err := subject.Commit()
	if err == nil {
		t.Errorf("Mark didn't fail properly without a previous mark: %s", err)
	}
}
