
package input

import "bytes"

func (input *Input) DelimitedWord(start rune, end rune) (bool, string) {
	if input.Current() != start {
		return false, ""
	}
	var word bytes.Buffer
	input.Advance()
	for input.Current() != end {
		word.WriteRune(input.Current())
		input.Advance()
	}
	input.Advance()
	return true, word.String()
}
