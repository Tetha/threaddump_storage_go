package input

const (
	Uninitialized = iota
	WaitingLine = iota
	BlockedLine = iota
)

type StacktraceLine struct {
	Type int
	LockAddress string
	LockClass string
}
