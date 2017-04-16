package input

const (
	WaitingLine = iota
)

type StacktraceLine struct {
	Type int
	LockAddress string
	LockClass string
}
