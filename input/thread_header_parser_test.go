package input

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var threadHeaderTests = []struct {
	input string

	shouldParse bool
	header      ThreadHeader
	current     byte
}{
	{
		"\"databaseQueryThread\" #242 prio=5 os_prio=0 tid=0x00007f645801f000 nid=0x4800 waiting on condition [0x00007f63df77a000]\n$",

		true,
		ThreadHeader{
			Name:             "databaseQueryThread",
			Id:               "242",
			IsDaemon:         false,
			Prio:             "5",
			OsPrio:           "0",
			Tid:              "0x00007f645801f000",
			Nid:              "0x4800",
			ThreadState:      "waiting on condition",
			ConditionAddress: "0x00007f63df77a000",
		},
		'$',
	},
	{
		"\"ping-JollyDolphin-repeating-task-watchdog\" #201 daemon prio=5 os_prio=0 tid=0x00007f6450040000 nid=0x4449 waiting on condition [0x00007f63e1798000]\n$",

		true,
		ThreadHeader{
			Name:             "ping-JollyDolphin-repeating-task-watchdog",
			Id:               "201",
			IsDaemon:         true,
			Prio:             "5",
			OsPrio:           "0",
			Tid:              "0x00007f6450040000",
			Nid:              "0x4449",
			ThreadState:      "waiting on condition",
			ConditionAddress: "0x00007f63e1798000",
		},
		'$',
	},
	{
		"\"VM Periodic Task Thread\" os_prio=0 tid=0x00007f6489288000 nid=0x41f8 waiting on condition\n$",

		true,
		ThreadHeader{
			Name:        "VM Periodic Task Thread",
			OsPrio:      "0",
			Tid:         "0x00007f6489288000",
			Nid:         "0x41f8",
			ThreadState: "waiting on condition",
		},
		'$',
	},
	{
		"\"elasticsearch[Samuel \"Starr\" Saxon][generic][T#446]\" #2315667 daemon prio=5 os_prio=0 tid=0x00007efa98010800 nid=0x7865 waiting on condition [0x00007efacf1b2000]\n$",

		true,
		ThreadHeader{
			Name:             "elasticsearch[Samuel \"Starr\" Saxon][generic][T#446]",
			IsDaemon:         true,
			Id:               "2315667",
			Prio:             "5",
			OsPrio:           "0",
			Tid:              "0x00007efa98010800",
			Nid:              "0x7865",
			ThreadState:      "waiting on condition",
			ConditionAddress: "0x00007efacf1b2000",
		},
		'$',
	},
}

func TestParseThreadHeader(t *testing.T) {
	for idx, tt := range threadHeaderTests {
		parser := CreateInput(tt.input)
		parsed, header := parser.ParseThreadHeader()

		assert := assert.New(t)
		desc := fmt.Sprintf("%d", idx)
		assert.Equal(tt.shouldParse, parsed, desc)
		if tt.shouldParse {
			assert.Equal(tt.header, header, desc)
		}
		assert.Equal(tt.current, parser.Current(), desc)
	}
}
