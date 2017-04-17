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

	if !input.MatchWord(" #") {
		return
	}
	parsed, header.Id = input.ReadUntil(' ')
	if !parsed {
		return
	}
	input.Advance() // skip space

	if input.MatchWord("daemon ") {
		header.IsDaemon = true
	}

	// Prio
	parsed, word = input.ReadUntil(' ')
	if !parsed {
		return
	}
	input.Advance()

	chunks := strings.Split(word, "=")
	if chunks[0] != "prio" {
		return
	}
	header.Prio = chunks[1]

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
