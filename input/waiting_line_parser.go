package input

func (input *Input) ParseWaitLine() (success bool, result StacktraceLine) {
	var parsed = false
	input.Mark()
	defer func() {
		if success {
			input.Commit()
		} else {
			input.Rollback()
		}
	}()

	result.Type = WaitingLine
	if !input.MatchWord("\t- waiting on ") {
		return
	}

	parsed, result.LockAddress = input.DelimitedWord('<', '>')
	if !parsed {
		return
	}

	parsed = input.MatchWord(" (a ")
	if !parsed {
		return
	}

	parsed, result.Class = input.readUntil(')')
	if !parsed {
		return
	}

	parsed = input.MatchWord(")\n")
	if !parsed {
		return
	}

	success = true
	return
}
