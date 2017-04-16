package input

import "testing"

func TestParseLockLine(t *testing.T) {
	parser := CreateInput("\t- locked <0x00000000e0e97cb0> (a io.netty.channel.nio.SelectedSelectionKeySet)\n$")
	parsed, line := parser.ParseLockedLine()

	if !parsed {
		t.Error("ParseLockedLine should succeed on valid inputs")
	}

	if line.Type != LockedLine {
		t.Errorf("Expected result lock line to be LockedLine (%d), but got %d", LockedLine, line.Type)
	}

	if line.LockAddress != "0x00000000e0e97cb0" {
		t.Errorf("Expected ParseLockedLine to extract LockAddress: <0x00000000e0e97cb0>, but got <%s>", line.LockAddress)
	}

	if line.LockClass != "io.netty.channel.nio.SelectedSelectionKeySet" {
		t.Errorf("Expeted ParseLockedLine to extract LockClass: io.netty.channel.nio.SelectedSelectionKeySet, but got <%s>", line.LockClass)
	}

	if parser.Current() != '$' {
		t.Errorf("Expected input to be placed on the next line but it got stuck on <%q>", parser.Current())
	}
}
