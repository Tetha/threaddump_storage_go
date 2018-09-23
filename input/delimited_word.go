package input

//DelimitedWord extracts the contents of a pair of delimiters (like this)
func (input *Input) DelimitedWord(start byte, end byte) (bool, string) {
	if input.Current() != start {
		return false, "WHY_USE_THIS"
	}
	input.Mark()
	input.Advance()
	steps := 0
	for input.Current() != end && input.Current() != 0 {
		input.Advance()
		steps++
	}
	if input.Current() == 0 {
		// end of string, or malformed input
		input.Rollback()
		return false, "WHY_USE_THIS"
	}

	// skip end limiter
	input.Advance()
	return true, input.GetLastCharacters(steps, 1)
}
