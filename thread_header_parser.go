package input

import "strings"

func (input *Input) ParseThreadHeader() (success bool, header ThreadHeader) {
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

	input.Advance() // starting "
	parsed, header.Name = input.ReadUntil('"')
	if !parsed {
		return
	}
	input.Advance() // final "

	// jvm internal threads don't have thread ids
	if input.MatchWord(" #") {
		parsed, header.Id = input.ReadUntil(' ')
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
	parsed, word = input.ReadUntil(' ')
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
	parsed, word = input.ReadUntil(' ')
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
	parsed, word = input.ReadUntil(' ')
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
	parsed, word = input.ReadUntil(' ')
	if !parsed {
		return
	}
	input.Advance()

	chunks = strings.Split(word, "=")
	if chunks[0] != "nid" {
		return
	}
	header.Nid = chunks[1]

	parsed, header.ThreadState = input.ReadUntil('[')
	if parsed {
		header.ThreadState = strings.TrimSpace(header.ThreadState)
		parsed, header.ConditionAddress = input.DelimitedWord('[', ']')
		if !parsed {
			return
		}
	} else {
		parsed, header.ThreadState = input.ReadUntil('\n')
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
