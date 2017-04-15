package input

func (input *Input) MatchWord(word string) bool {
	input.Mark()
	for _, currentExpected := range word {
		current := input.Current()
		if current == currentExpected {
			input.Advance()
		} else {
			input.Rollback()
			return false
		}
	}
	input.Commit()
	return true
}
