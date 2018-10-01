package input

func (input *Input) parseThread() (success bool, result Thread) {
	parsed := false

	headerParsed, header := input.parseThreadHeader()
	if !headerParsed {
		return
	}
	result.Name, result.ID, result.IsDaemon, result.Prio, result.OsPrio, result.Tid, result.Nid, result.ThreadState, result.ConditionAddress = header.Name, header.ID, header.IsDaemon, header.Prio, header.OsPrio, header.Tid, header.Nid, header.ThreadState, header.ConditionAddress

	parsed, result.JavaState, result.JavaStateDetail = input.parseThreadState()
	if !parsed {
		return
	}
	for {
		lineParsed, line := input.parseStacktraceLine()
		if lineParsed {
			result.Stacktrace = append(result.Stacktrace, line)
		} else {
			break
		}
	}

	if !input.MatchWord("\n") {
		return
	}
	success = true
	return
}

func (input *Input) parseStacktraceLine() (success bool, line StacktraceLine) {
	input.Mark()
	success, line = input.parseWaitLine()
	if success {
		input.Commit()
		return
	}
	input.Rollback()

	input.Mark()
	success, line = input.parseBlockedLine()
	if success {
		input.Commit()
		return
	}
	input.Rollback()

	input.Mark()
	success, line = input.parseLockedLine()
	if success {
		input.Commit()
		return
	}
	input.Rollback()

	input.Mark()
	success, line = input.parseThreadPosition()
	if success {
		input.Commit()
		return
	}
	input.Rollback()

	input.Mark()
	success, line = input.parseParkedLine()
	if success {
		input.Commit()
		return
	}
	input.Rollback()
	return
}
