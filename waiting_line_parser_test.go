package input

import "testing"

func TestParseWaitLine(t *testing.T) {
	parser := CreateInput("\t- waiting on <0x0000000729ffa7f8> (a java.lang.ref.ReferenceQueue$Lock)\n$")
	parsed, line := parser.ParseWaitLine()
	if !parsed {
		t.Error("ParseWaitLine didn't succeed on valid input")
	}

	if line.Type != WaitingLine {
		t.Errorf("ParseWaitline should have returned a WaitingLine (%d), but created a %d", WaitingLine, line.Type)
	}

	if line.LockAddress != "0x0000000729ffa7f8" {
		t.Errorf("Lock address wasn't extracted properly, got: <%s>", line.LockAddress)
	}

	if line.LockClass != "java.lang.ref.ReferenceQueue$Lock" {
		t.Errorf("Lock class wasn't extracted properly, got: <%s>", line.LockClass)
	}
}
