package input

import "strings"

func (input *Input) ParseThreadState() (success bool, state string, clarification string) {
	parsed := false
	input.Mark()
	defer func() {
		if success {
			input.Mark()
		} else {
			input.Commit()
		}
	}()

	if !input.MatchWord("   java.lang.Thread.State: ") {
		return
	}

	// need to be careful because this might read too far
	// if that's the case we need to re-start reading the
	// state from the current position
	input.Mark()
	parsed, state = input.readUntil('(')
	if parsed && !strings.Contains(state, "\n") { // stay on the same line
		input.Mark()
		state = strings.TrimSpace(state)
		parsed, clarification = input.DelimitedWord('(', ')')
		if !parsed {
			return
		}
	} else {
		input.Rollback()
		parsed, state = input.readUntil('\n')
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
