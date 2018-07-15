package input

func (input *Input) ParseParkedLine() (success bool, result StacktraceLine) {
	var parsed = false
	input.Mark()
	defer func() {
		if success {
			input.Commit()
		} else {
			input.Rollback()
		}
	}()

	if !input.MatchWord("\t- parking to wait for  ") {
		return
	}

	result.Type = ParkedLine

	parsed, result.LockAddress = input.DelimitedWord('<', '>')
	if !parsed {
		return
	}

	if !input.MatchWord(" (a ") {
		return
	}

	parsed, result.Class = input.ReadUntil(')')
	if !parsed {
		return
	}

	if !input.MatchWord(")\n") {
		return
	}

	success = true
	return
}
