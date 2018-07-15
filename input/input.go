package input

import "errors"

type Input struct {
	content  string
	position int
	marks    []int
}

func CreateInput(content string) (r Input) {
	r = Input{content: content, position: 0}
	r.marks = make([]int, 0, 10)
	return
}

func (input *Input) Slice(start int, end int) string {
	return input.content[start:end]
}

func (input *Input) Position() int {
	return input.position
}

func (input *Input) Length() int {
	return len(input.content)
}

func (input *Input) Current() byte {
	if input.position < input.Length() {
		return input.content[input.position]
	} else {
		return 0
	}
}

func (input *Input) CurrentMany(num int) string {
	if input.position+num < input.Length() {
		return input.content[input.position : input.position+num]
	} else {
		return input.content[input.position:]
	}
}

func (input *Input) GetLastCharacters(n int, skip int) string {
	return input.content[input.position-n-skip : input.position-skip]
}

// these methods are incorreect as soon we encounter a string with
// >1 byte per code point. But it's quite a bit faster this way.

func (input *Input) Advance() {
	//_, width := utf8.DecodeRuneInString(input.content[input.position:])
	input.position++
}

func (input *Input) AdvanceMany(num int) {
	//_, width := utf8.DecodeRuneInString(input.content[input.position:])
	input.position += num
}

func (input *Input) Mark() {
	input.marks = append(input.marks, input.position)
}

func (input *Input) Rollback() error {
	if len(input.marks) == 0 {
		return errors.New("no previous mark")
	}
	lastPosition := -1
	lastPosition, input.marks = input.marks[len(input.marks)-1], input.marks[:len(input.marks)-1]
	input.position = lastPosition
	return nil
}

func (input *Input) Commit() error {
	if len(input.marks) == 0 {
		return errors.New("no previous mark")
	}
	input.marks = input.marks[:len(input.marks)-1]
	return nil
}
