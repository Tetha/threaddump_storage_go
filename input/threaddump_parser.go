package input

func (input *Input) ParseThreaddump() (parseFailure string, result Threaddump) {
	parseFailure = "Unknown error"

	input.Mark()
	if !input.MatchWord("Full thread dump") {
		parseFailure = "Could not match header <Full Thread dump>"
		input.Rollback()
		return
	}
	input.Rollback()
	if parsed, header := input.readUntil('\n'); parsed {
		result.Header = header
	} else {
		parseFailure = "Could not find header terminating newline"
		return
	}

	if !input.MatchWord("\n\n") {
		parseFailure = "Missing double newline after header"
		return
	}

	for {
		threadParsed, thread := input.ParseThread()
		if threadParsed {
			result.Threads = append(result.Threads, thread)
		} else {
			parseFailure = ""
			return
		}
	}

	return
}
