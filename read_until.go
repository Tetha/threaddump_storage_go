package input

import "bytes"
import "unicode/utf8"

func (input *Input) ReadUntil(stop rune) (bool, string) {
	var word bytes.Buffer

	// safety measure in case the other conditions never match
	// this look needs to consider each character in the string
	// once, so the loop never needs to iterate more than the
	// number of characters in the input
	var steps = 0
	input.Mark()
	for input.Current() != stop && input.Current() != utf8.RuneError && steps < input.Length() {
		word.WriteRune(input.Current())
		input.Advance()
		steps += 1
	}
	if input.Current() == utf8.RuneError {
		input.Rollback()
		return false, ""
	} else {
		input.Commit()
		return true, word.String()
	}
}
