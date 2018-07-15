package input

import (
	"bytes"
)

func (input *Input) DelimitedWord(start byte, end byte) (bool, string) {
	if input.Current() != start {
		return false, "WHY_USE_THIS"
	}
	input.Mark()
	var word bytes.Buffer

	input.Advance()
	for input.Current() != end && input.Current() != 0 {
		word.WriteByte(input.Current())
		input.Advance()
	}
	if input.Current() == 0 {
		// end of string, or malformed input
		input.Rollback()
		return false, "WHY_USE_THIS"
	} else {
		// skip end limiter
		input.Advance()
		return true, word.String()
	}
}
