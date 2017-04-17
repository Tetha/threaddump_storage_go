package input

func (input *Input) ParseLockedLine() (success bool, result StacktraceLine) {
	var parsed = false
	input.Mark()

	if !input.MatchWord("\t- locked ") {
		input.Rollback()
		return
	}

	result.Type = LockedLine

	parsed, result.LockAddress = input.DelimitedWord('<', '>')
	if !parsed {
		input.Rollback()
		return
	}

	if !input.MatchWord(" (a ") {
		input.Rollback()
		return
	}

	parsed, result.Class = input.ReadUntil(')')
	if !parsed {
		input.Rollback()
		return
	}

	if !input.MatchWord(")\n") {
		input.Rollback()
		return
	}

	input.Commit()
	success = true
	return
}
