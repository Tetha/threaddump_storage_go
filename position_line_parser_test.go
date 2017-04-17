package input;

import "testing"

func TestJavaCodePosition(t *testing.T) {
	parser := CreateInput("\tat sun.nio.ch.EPollArrayWrapper.poll(EPollArrayWrapper.java:269)\n$")
	parsed, line := parser.ParseThreadPosition()

	if !parsed {
		t.Error("ParseThreadPosition should succeed with valid inputs")
	}

	if line.Type != PositionLine {
		t.Errorf("Expected result to be PositionLine (%d), but was %d", PositionLine, line.Type)
	}

	if line.Class != "sun.nio.ch.EPollArrayWrapper" {
		t.Errorf("Expected ParseThreadPosition to extract class <sun.nio.ch.EPollArrayWrapper>, but got <%s>", line.Class)
	}

	if line.Method != "poll" {
		t.Errorf("Expected ParseThreadPosition to extract Method <poll> but got <%s>", line.Method)
	}

	if line.SourceFile != "EPollArrayWrapper.java" {
		t.Errorf("Expected ParseThreadPosition to extract SourceFile <EPollArrayWrapper.java>, but got <%s>", line.SourceFile)
	}

	if line.SourceLine != 269 {
		t.Errorf("Expected ParseThreadPosition to extract SourceLine <269> but got <%d>", line.SourceLine)
	}

	if parser.Current() != '$' {
		t.Errorf("Expected ParseThreadPosition to consume the entire line but it got stuck on <%q>", parser.Current())
	}
}

func TestNativePositionLine(t *testing.T) {
	parser := CreateInput("\tat sun.misc.Unsafe.park(Native Method)\n$")
	parsed, line := parser.ParseThreadPosition()
	if !parsed {
		t.Errorf("ParseThreadPosition should succeed with valid positions")
	}

	if line.Type != PositionLine {
		t.Errorf("Expected result to be PositionLine (%d), but was %d", PositionLine, line.Type)
	}

	if line.Class != "sun.misc.Unsafe" {
		t.Errorf("Expected ParseThreadPosition to extract class <sun.misc.Unsave>, but got <%s>", line.Class)
	}

	if line.Method != "park" {
		t.Errorf("Expected ParseThreadPosition to extract Method <park> but got <%s>", line.Method)
	}


	if line.SourceFile != "Native Method" {
		t.Errorf("Expected ParseThreadPosition to extract SourceFile <Native Method>, but got <%s>", line.SourceFile)
	}

	if line.SourceLine != -1 {
		t.Errorf("Expected ParseThreadPosition to extract SourceLine <-1> but got <%d>", line.SourceLine)
	}

	if parser.Current() != '$' {
		t.Errorf("Expected ParseThreadPosition to consume the entire line but it got stuck on <%q>", parser.Current())
	}
}
