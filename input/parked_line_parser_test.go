package input

import "testing"

var parkedLineTests = []lineParserTest{
	{
		"\t- parking to wait for  <0x000000071fb23b18> (a java.util.concurrent.SynchronousQueue$TransferStack)\n$",
		true,
		StacktraceLine{
			Type:        ParkedLine,
			LockAddress: "0x000000071fb23b18",
			Class:       "java.util.concurrent.SynchronousQueue$TransferStack",
		},
		'$',
	},
}

func TestParseParkedLine(t *testing.T) {
	for idx, tt := range parkedLineTests {
		runParserTestcase(t, idx, tt, (*Input).parseParkedLine)
	}
}
