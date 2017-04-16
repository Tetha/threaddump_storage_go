package input

const (
	Uninitialized = iota
	WaitingLine = iota
	BlockedLine = iota
	LockedLine = iota
)

type StacktraceLine struct {
	Type int
	LockAddress string
	LockClass string
}
