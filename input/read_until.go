package input

func (input *Input) ReadUntil(stop byte) (bool, string) {
	// safety measure in case the other conditions never match
	// this look needs to consider each character in the string
	// once, so the loop never needs to iterate more than the
	// number of characters in the input
	var steps = 0
	input.Mark()
	for input.Current() != stop && steps < input.Length() {
		input.Advance()
		steps++
	}
	if input.Current() != stop {
		input.Rollback()
		return false, ""
	} else {
		input.Commit()
		return true, input.GetLastCharacters(steps, 0)
	}
}
