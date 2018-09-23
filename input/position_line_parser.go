package input

import "strings"
import "strconv"

func (input *Input) parseThreadPosition() (success bool, result StacktraceLine) {
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

	result.Type = PositionLine
	if !input.MatchWord("\tat ") {
		return
	}

	var call string
	parsed, call = input.readUntil('(')
	if !parsed {
		return
	}
	chunks := strings.Split(call, ".")
	result.Class = strings.Join(chunks[:len(chunks)-1], ".")
	result.Method = chunks[len(chunks)-1]

	parsed, sourceLocation := input.DelimitedWord('(', ')')
	if !parsed {
		return
	}
	if strings.Contains(sourceLocation, ":") {
		locationChunks := strings.Split(sourceLocation, ":")
		result.SourceFile = locationChunks[0]

		line, err := strconv.Atoi(locationChunks[1])
		if err != nil {
			return
		}
		result.SourceLine = line
	} else {
		result.SourceFile = sourceLocation
		result.SourceLine = -1
	}

	if !input.MatchWord("\n") {
		return
	}

	success = true
	return
}
