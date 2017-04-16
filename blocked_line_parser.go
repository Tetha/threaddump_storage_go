package input

func (input *Input) ParseBlockedLine() (bool, StacktraceLine) {
	var result StacktraceLine
	var parsed = false

	input.Mark()
	if !input.MatchWord("\t- parking to wait for ") {
		input.Rollback()
		return false, result
	}

	parsed, result.LockAddress = input.DelimitedWord('<', '>')
	if !parsed {
		input.Rollback()
		return false, result
	}

	if !input.MatchWord(" (a ") {
		input.Rollback()
		return false, result
	}

	parsed, result.LockClass = input.ReadUntil(')')
	if !parsed {
		input.Rollback()
		return false, result
	}

	if !input.MatchWord(")\n") {
		input.Rollback()
		return false, result
	}

	input.Commit()
	result.Type = BlockedLine
	return true, result
}
