package input

import "errors"

type Input struct {
    content string;
    position int;
    marks []int;
}

func CreateInput(content string) (r Input) {
    r = Input{content: content, position: 0}
    r.marks = make([]int, 10)
    return
}

func (input *Input) Current() rune {
    runeValue, width := utf8.DecodeRuneInString(input.content[input.position:])
    return runeValue
}

func (input *Input) Advance() {
    runeValue, width := utf8.DecodeRuneInString(input.content[input.position:])
    input.position += width
}
