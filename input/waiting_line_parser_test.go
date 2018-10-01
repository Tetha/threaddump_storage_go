package input

import "testing"

var waitLineTests = []lineParserTest{
	{
		"\t- waiting on <0x0000000729ffa7f8> (a java.lang.ref.ReferenceQueue$Lock)\n$",
		true,
		StacktraceLine{
			Type:        WaitingLine,
			LockAddress: "0x0000000729ffa7f8",
			Class:       "java.lang.ref.ReferenceQueue$Lock",
		},
		'$',
	},
}

func TestParseWaitLine(t *testing.T) {
	for idx, tt := range waitLineTests {
		runParserTestcase(t, idx, tt, (*Input).ParseWaitLine)
	}
}
