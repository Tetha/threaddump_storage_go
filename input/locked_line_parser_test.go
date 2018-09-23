package input

import "testing"

var lockLineTests = []lineParserTest{
	{
		"\t- locked <0x00000000e0e97cb0> (a io.netty.channel.nio.SelectedSelectionKeySet)\n$",
		true,
		StacktraceLine{
			Type:        LockedLine,
			LockAddress: "0x00000000e0e97cb0",
			Class:       "io.netty.channel.nio.SelectedSelectionKeySet",
		},
		'$',
	},
}

func TestParseLockLine(t *testing.T) {
	for idx, tt := range lockLineTests {
		runParserTestcase(t, idx, tt, (*Input).parseLockedLine)
	}
}
