package input

func (input *Input) ParseBlockedLine() (bool, StacktraceLine) {
	var result StacktraceLine
	var parsed = false

	if !input.MatchWord("\t- parking to wait for ") {
		return false, result
	}

	parsed, result.LockAddress = input.DelimitedWord('<', '>')
	if !parsed {
		return false, result
	}

	if !input.MatchWord(" (a ") {
		return false, result
	}

	parsed, result.LockClass = input.ReadUntil(')')
	if !parsed {
		return false, result
	}

	if !input.MatchWord(")\n") {
		return false, result
	}

	result.Type = BlockedLine
	return true, result
}
