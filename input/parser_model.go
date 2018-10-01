package input

const (
	Uninitialized = iota
	WaitingLine
	BlockedLine
	LockedLine
	PositionLine
	ParkedLine
)

type StacktraceLine struct {
	Type        int
	LockAddress string
	Class       string
	Method      string
	SourceFile  string
	SourceLine  int
}

type ThreadHeader struct {
	Name             string
	ID               string
	IsDaemon         bool
	Prio             string
	OsPrio           string
	Tid              string
	Nid              string
	ThreadState      string
	ConditionAddress string
}

type Thread struct {
	Name             string
	ID               string
	IsDaemon         bool
	Prio             string
	OsPrio           string
	Tid              string
	Nid              string
	ThreadState      string
	ConditionAddress string
	JavaState        string
	JavaStateDetail  string
	Stacktrace       []StacktraceLine
}

type Threaddump struct {
	Header  string
	Threads []Thread
}
