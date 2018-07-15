package input

import "testing"

func TestParseParkedLine(t *testing.T) {
	parser := CreateInput("\t- parking to wait for  <0x000000071fb23b18> (a java.util.concurrent.SynchronousQueue$TransferStack)\n$")
	parsed, line := parser.ParseParkedLine()

	if !parsed {
		t.Error("ParseLockedLine should succeed on valid inputs")
	}

	if line.Type != ParkedLine {
		t.Errorf("Expected result lock line to be ParkedLine (%d), but got %d", ParkedLine, line.Type)
	}

	if line.LockAddress != "0x000000071fb23b18" {
		t.Errorf("Expected ParseLockedLine to extract LockAddress: <0x00000000e0e97cb0>, but got <%s>", line.LockAddress)
	}

	if line.Class != "java.util.concurrent.SynchronousQueue$TransferStack" {
		t.Errorf("Expeted ParseLockedLine to extract Class: java.util.concurrent.SynchronousQueue$TransferStack, but got <%s>", line.Class)
	}

	if parser.Current() != '$' {
		t.Errorf("Expected input to be placed on the next line but it got stuck on <%q>", parser.Current())
	}
}
