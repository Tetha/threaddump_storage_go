package input

func (input *Input) ParseBlockedLine() (success bool, result StacktraceLine) {
	var parsed = false
	success = false

	result.Type = BlockedLine
	input.Mark()
	if !(input.MatchWord("\t- parking to wait for ") || input.MatchWord("\t- waiting to lock ")) {
		input.Rollback()
		return
	}

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
