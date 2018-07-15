package input

func (input *Input) MatchWord(word string) bool {
	input.Mark()
	for i := 0; i < len(word); i++ {
		current := input.Current()
		if current == word[i] {
			input.Advance()
		} else {
			input.Rollback()
			return false
		}
	}
	input.Commit()
	return true
}
