package input

import "strings"

func (input *Input) parseThreadHeader() (success bool, header ThreadHeader) {
	parsed := false
	word := ""
	input.Mark()
	defer func() {
		if success {
			input.Mark()
		} else {
			input.Commit()
		}
	}()

	if input.Current() != '"' {
		return
	}

	tries := 0
	input.Advance() // starting "
	start := input.Position()
	for tries < input.Length()+10 {
		tries++
		parsed, _ = input.readUntil('"')
		if !parsed {
			return
		}
		input.Advance() // final "

		// lookahead - always roll back
		input.Mark()
		if input.MatchWord(" #") || input.MatchWord(" prio=") || input.MatchWord(" os_prio=") {
			input.Rollback()
			header.Name = input.Slice(start, input.Position()-1)
			break
		}
	}

	// jvm internal threads don't have thread ids
	if input.MatchWord(" #") {
		parsed, header.Id = input.readUntil(' ')
		if !parsed {
			return
		}
	}
	input.Advance() // skip space

	if input.MatchWord("daemon ") {
		header.IsDaemon = true
	}

	// Prio
	// Prio is optional, e.g. internal threads don't have it
	input.Mark()
	parsed, word = input.readUntil(' ')
	if !parsed {
		return
	}
	input.Advance()

	chunks := strings.Split(word, "=")
	if chunks[0] == "prio" {
		input.Commit()
		header.Prio = chunks[1]
	} else {
		// rollback so the os_prio starts right after ID/daemon
		input.Rollback()
	}

	// Os-Prio
	parsed, word = input.readUntil(' ')
	if !parsed {
		return
	}
	input.Advance()

	chunks = strings.Split(word, "=")
	if chunks[0] != "os_prio" {
		return
	}
	header.OsPrio = chunks[1]

	// Tid
	parsed, word = input.readUntil(' ')
	if !parsed {
		return
	}
	input.Advance()

	chunks = strings.Split(word, "=")
	if chunks[0] != "tid" {
		return
	}
	header.Tid = chunks[1]

	// Nid
	parsed, word = input.readUntil(' ')
	if !parsed {
		return
	}
	input.Advance()

	chunks = strings.Split(word, "=")
	if chunks[0] != "nid" {
		return
	}
	header.Nid = chunks[1]

	parsed, header.ThreadState = input.readUntil('[')
	if parsed {
		header.ThreadState = strings.TrimSpace(header.ThreadState)
		parsed, header.ConditionAddress = input.DelimitedWord('[', ']')
		if !parsed {
			return
		}
	} else {
		parsed, header.ThreadState = input.readUntil('\n')
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
