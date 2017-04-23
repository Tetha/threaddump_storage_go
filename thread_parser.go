package input

func (input *Input) ParseThread() (success bool, result Thread) {
	parsed := false

	headerParsed, header := input.ParseThreadHeader()
	if !headerParsed {
		return
	}
	result.Name, result.Id, result.IsDaemon, result.Prio, result.OsPrio, result.Tid, result.Nid, result.ThreadState, result.ConditionAddress = header.Name, header.Id, header.IsDaemon, header.Prio, header.OsPrio, header.Tid, header.Nid, header.ThreadState, header.ConditionAddress

	parsed, result.JavaState, result.JavaStateDetail = input.ParseThreadState()
	if !parsed {
		return
	}
	return
}
