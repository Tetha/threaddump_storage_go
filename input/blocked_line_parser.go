package input

func (input *Input) parseBlockedLine() (success bool, result StacktraceLine) {
	var parsed = false
	success = false

	input.Mark()
	defer func() {
		if success {
			input.Commit()
		} else {
			input.Rollback()
		}
	}()

	result.Type = BlockedLine

	if !(input.MatchWord("\t- parking to wait for ") || input.MatchWord("\t- waiting to lock ")) {
		return
	}

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
