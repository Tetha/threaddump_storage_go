package input

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

func (input *Input) Current() byte {
    return input.content[input.position]
}
