package input

func (input *Input) ParseWaitLine() (bool, StacktraceLine) {
	var parseResult StacktraceLine
	var parsed = false
	input.Mark()
	if !input.MatchWord("\t- waiting on ") {
		input.Rollback()
		return false, parseResult
	}

	parsed, parseResult.LockAddress = input.DelimitedWord('<', '>')
	if !parsed {
		input.Rollback()
		return false, parseResult
	}

	parsed = input.MatchWord(" (a ")
	if !parsed {
		input.Rollback()
		return false, parseResult
	}

	parsed, parseResult.LockClass = input.ReadUntil(')')
	if !parsed {
		input.Rollback()
		return false, parseResult
	}

	parsed = input.MatchWord(")\n")
	if !parsed {
		input.Rollback()
		return false, parseResult
	}

	return true, parseResult
}
