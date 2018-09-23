package input

//MatchWord advances the input iff it starts with the given string
func (input *Input) MatchWord(word string) bool {
	// There is no mark-handling here, because MatchWord
	// has no speculative input consumption. It does a
	// n-word lookahead without moving the cursor and
	// changes the cursor only if the lookahead is the
	// expected word.
	inputWord := input.CurrentMany(len(word))
	if inputWord != word {
		return false
	}
	input.AdvanceMany(len(word))
	return true
}
