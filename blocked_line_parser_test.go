package input

import "testing"

func TestParseBlockedLineValidInput(t *testing.T) {
	parser := CreateInput("\t- parking to wait for <0x000000065be92a68> (a java.util.concurrent.locks.AbstractQueuedSynchronizer$ConditionObject)\n$")
	parsed, line := parser.ParseBlockedLine()
	if !parsed {
		t.Error("ParseBlockedLine didn't succeed on valid input")
	}

	if line.Type != BlockedLine {
		t.Errorf("Expected type to be BlockedLine (%d), but was %d", BlockedLine, line.Type)
	}

	if line.LockAddress != "0x000000065be92a68" {
		t.Errorf("Expected extracted lock address to be <0x000000065be92a68>, got <%s>", line.LockAddress)
	}

	if line.LockClass != "java.util.concurrent.locks.AbstractQueuedSynchronizer$ConditionObject" {
		t.Errorf("Expected extracted lock class to be <java.util.concurrent.locks.AbstractQueuedSynchronizer$ConditionObject>, but got <%s>", line.LockClass)
	}

	if parser.Current() != '$' {
		t.Errorf("Expected ParseBlockedLine to advance the cursor to the next line, but it got stuck on <%q>", parser.Current())
	}
}
