package input

import "strings"

func (input *Input) ParseThreadState() (success bool, state string, clarification string) {
	parsed := false

	if !input.MatchWord("   java.lang.Thread.State: ") {
		return
	}

	parsed, state = input.ReadUntil('(')
	if parsed {
		state = strings.TrimSpace(state)
		parsed, clarification = input.DelimitedWord('(', ')')
		if !parsed {
			return
		}
	} else {
		parsed, state = input.ReadUntil('\n')
		if !parsed {
			return
		}
	}

	if !input.MatchWord("\n") {
		return
	}

	success = true
	return
}
