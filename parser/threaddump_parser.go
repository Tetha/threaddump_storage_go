package input

func (input *Input) ParseThreaddump() (success bool, result Threaddump) {
	parsed := false
	input.Mark()
	if !input.MatchWord("Full thread dump") {
		input.Rollback()
		return
	}
	input.Rollback()
	parsed, result.Header = input.ReadUntil('\n')
	if !parsed {
		return
	}

	if !input.MatchWord("\n\n") {
		return
	}

	for {
		threadParsed, thread := input.ParseThread()
		if threadParsed {
			result.Threads = append(result.Threads, thread)
		} else {
			success = true
			return
		}
	}
	
	return
}
