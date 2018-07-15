package input

import "bytes"
import "unicode/utf8"

func (input *Input) DelimitedWord(start rune, end rune) (bool, string) {
	if input.Current() != start {
		return false, "WHY_USE_THIS"
	}
	input.Mark()
	var word bytes.Buffer
	input.Advance()
	for input.Current() != end && input.Current() != utf8.RuneError {
		word.WriteRune(input.Current())
		input.Advance()
	}
	if input.Current() == utf8.RuneError {
		// end of string, or malformed input
		input.Rollback()
		return false, "WHY_USE_THIS"
	} else {
		// skip end limiter
		input.Advance()
		return true, word.String()
	}
}
