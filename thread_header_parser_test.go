package input

import "testing"

func TestParseThreadHeaderApplicationThread(t *testing.T) {
	parser := CreateInput("\"databaseQueryThread\" #242 prio=5 os_prio=0 tid=0x00007f645801f000 nid=0x4800 waiting on condition [0x00007f63df77a000]\n$")

	parsed, header := parser.ParseThreadHeader()

	if !parsed {
		t.Errorf("ParseThreadHeader must succeed for valid input")
	}

	if header.Name != "databaseQueryThread" {
		t.Errorf("ParseThreadHeader did not extract the correct thread name <databaseQueryThread>, but <%s>", header.Name)
	}

	if header.Id != "242" {
		t.Errorf("ParseThreadHeader did not extract the correct id <242>, but <%s>", header.Id)
	}

	if header.IsDaemon {
		t.Errorf("ParseThreadHeader must not flag the thread as a daemon")
	}

	if header.Prio != "5" {
		t.Errorf("ParseThreadHeader did not extract the right prio <5> but <%s>", header.Prio)
	}

	if header.OsPrio != "0" {
		t.Errorf("ParseThreadHeader did not extract the right os_prio <0> but <%s>", header.OsPrio)
	}

	if header.Tid != "0x00007f645801f000" {
		t.Errorf("ParseThreadHeader did not extract the right tid <0x00007f645801f000>, but <%s>", header.Tid)
	}

	if header.Nid != "0x4800" {
		t.Errorf("ParseThreadHeader did not extract the right Nid <0x4800>, but <%s>", header.Nid)
	}

	if header.ThreadState != "waiting on condition" {
		t.Errorf("ParseThreadHeader did not extract the right ThreadState <waiting on condition>, but <%s>", header.ThreadState)
	}

	if header.ConditionAddress != "0x00007f63df77a000" {
		t.Errorf("ParseThreadHeader did not extract the right ConditionAddress <0x00007f63df77a000>, but <%s>", header.ConditionAddress)
	}
	if parser.Current() != '$' {
		t.Errorf("Expected parser to move the input past the line")
	}
}

func TestParseThreadHeaderDaemonThread(t *testing.T) {
	parser := CreateInput("\"ping-JollyDolphin-repeating-task-watchdog\" #201 daemon prio=5 os_prio=0 tid=0x00007f6450040000 nid=0x4449 waiting on condition [0x00007f63e1798000]\n$")
	parsed, header := parser.ParseThreadHeader()

	if !parsed {
		t.Error("Expected thread header parse to suceed on valid input")
	}

	if !header.IsDaemon {
		t.Error("Expected ParseThreadHeader to mark DaemonThreads as daemons")
	}

	if header.Prio != "5" {
		t.Error("Expected ParseThreadHeader to parse the rest of the fields properly")
	}
}

func TestParseThreadHeaderJvmThreads(t *testing.T) {
	parser := CreateInput("\"VM Periodic Task Thread\" os_prio=0 tid=0x00007f6489288000 nid=0x41f8 waiting on condition\n$")

	parsed, header := parser.ParseThreadHeader()

	if !parsed {
		t.Error("ParseThreadHeader should succeed with valid inputs")
	}

	if header.Name != "VM Periodic Task Thread" {
		t.Errorf("Expected ParseThreadHeader to extract the name <VM Periodic Task Thread>, but got %s", header.Name)
	}

	if header.OsPrio != "0" {
		t.Errorf("Expected ParseThreadHeader to extract the os_prio <0>, but got %s", header.OsPrio)
	}

	if parser.Current() != '$' {
		t.Errorf("Expected ParseThreadHeader to consume the entire line but it got stuck on %q", parser.Current())
	}
}

func TestParseThreadHeaderThreadNameWithQuotes(t *testing.T) {
	parser := CreateInput("\"elasticsearch[Samuel \"Starr\" Saxon][generic][T#446]\" #2315667 daemon prio=5 os_prio=0 tid=0x00007efa98010800 nid=0x7865 waiting on condition [0x00007efacf1b2000]\n$")

	parsed, header := parser.ParseThreadHeader()

	if !parsed {
		t.Errorf("ParseThreadHeader has to succeeedon valid input")
	}

	if header.Name != "elasticsearch[Samuel \"Starr\" Saxon][generic][T#446]" {
		t.Errorf("Expected ParseThreadHeader to extract thread name <elasticsearch[Samuel \"Starr\" Saxon][generic][T#446]>, but got <%s>", header.Name)
	}
}
