package input

func (input *Input) parseLockedLine() (success bool, result StacktraceLine) {
	var parsed = false
	input.Mark()
	defer func() {
		if success {
			input.Commit()
		} else {
			input.Rollback()
		}
	}()

	if !input.MatchWord("\t- locked ") {
		return
	}

	result.Type = LockedLine

	parsed, result.LockAddress = input.DelimitedWord('<', '>')
	if !parsed {
		return
	}

	if !input.MatchWord(" (a ") {
		return
	}

	parsed, result.Class = input.readUntil(')')
	if !parsed {
		return
	}

	if !input.MatchWord(")\n") {
		return
	}

	success = true
	return
}
