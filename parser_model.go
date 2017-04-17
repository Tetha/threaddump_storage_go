package input

const (
	Uninitialized = iota
	WaitingLine = iota
	BlockedLine = iota
	LockedLine = iota
	PositionLine = iota
)

type StacktraceLine struct {
	Type int
	LockAddress string
	Class string
	Method string
	SourceFile string
	SourceLine int
}
