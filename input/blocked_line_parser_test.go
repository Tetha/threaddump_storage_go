package input

import "testing"

var parseBlockedLineTests = []lineParserTest{
	{
		"\t- parking to wait for <0x000000065be92a68> (a java.util.concurrent.locks.AbstractQueuedSynchronizer$ConditionObject)\n$",
		true,
		StacktraceLine{
			Type:        BlockedLine,
			LockAddress: "0x000000065be92a68",
			Class:       "java.util.concurrent.locks.AbstractQueuedSynchronizer$ConditionObject",
		},
		'$',
	},
	{
		"\t- waiting to lock <0x00000000e0c5a010> (a org.apache.logging.log4j.core.appender.FileManager)\n$",
		true,
		StacktraceLine{
			Type:        BlockedLine,
			LockAddress: "0x00000000e0c5a010",
			Class:       "org.apache.logging.log4j.core.appender.FileManager",
		},
		'$',
	},
}

func TestParseBlockedLine(t *testing.T) {
	for idx, tt := range parseBlockedLineTests {
		runParserTestcase(t, idx, tt, (*Input).parseBlockedLine)
	}
}
