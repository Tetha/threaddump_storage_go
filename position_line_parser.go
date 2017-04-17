package input;

import "strings"
import "strconv"

func (input *Input) ParseThreadPosition() (success bool, result StacktraceLine) {
	input.Mark()
	var parsed = false
	success = false

	result.Type = PositionLine
	if !input.MatchWord("\tat ") {
		input.Rollback()
		return
	}

	var call string
	parsed, call = input.ReadUntil('(')
	if !parsed {
		input.Rollback()
		return
	}
	chunks := strings.Split(call, ".")
	result.Class = strings.Join(chunks[:len(chunks)-1], ".")
	result.Method = chunks[len(chunks)-1]

	parsed, sourceLocation := input.DelimitedWord('(', ')')
	if !parsed {
		input.Rollback()
		return
	}
	if sourceLocation == "Native Method" {
		result.SourceFile = "Native Method"
		result.SourceLine = -1
	} else {
		locationChunks := strings.Split(sourceLocation, ":")
		result.SourceFile = locationChunks[0]

		line, err := strconv.Atoi(locationChunks[1])
		if err != nil {
			input.Rollback()
			return
		}
		result.SourceLine = line
	}

	if !input.MatchWord("\n") {
		input.Rollback()
		return
	}

	input.Commit()
	success = true
	return
}
