package input

import "testing"

var positionLineTests = []lineParserTest{
	{
		"\tat sun.nio.ch.EPollArrayWrapper.poll(EPollArrayWrapper.java:269)\n$",
		true,
		StacktraceLine{
			Type:       PositionLine,
			Class:      "sun.nio.ch.EPollArrayWrapper",
			Method:     "poll",
			SourceFile: "EPollArrayWrapper.java",
			SourceLine: 269,
		},
		'$',
	},
	{
		"\tat sun.misc.Unsafe.park(Native Method)\n$",
		true,
		StacktraceLine{
			Type:       PositionLine,
			Class:      "sun.misc.Unsafe",
			Method:     "park",
			SourceFile: "Native Method",
			SourceLine: -1,
		},
		'$',
	},
}

func TestParseThreadPosition(t *testing.T) {
	for idx, tt := range positionLineTests {
		runParserTestcase(t, idx, tt, (*Input).parseThreadPosition)
	}
}
