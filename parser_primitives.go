
package input

func (input *Input) MatchWord(word string) bool {
	for _, currentExpected := range word {
		current := input.Current()
		if current == currentExpected {
			input.Advance()
		} else {
			return false
		}
	}
	return true
}
