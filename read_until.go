package input

import "bytes"

func (input *Input) ReadUntil(stop rune) (bool, string) {
	var word bytes.Buffer

	for input.Current() != stop {
		word.WriteRune(input.Current())
		input.Advance()
	}
	return true, word.String()
}
