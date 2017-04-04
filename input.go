package input

type input struct {
    content string;
    position int;
    marks []int;
}

func CreateInput(content string) (r input) {
    r = input{content: content, position: 0}
    r.marks = make([]int, 10)
    return
}
