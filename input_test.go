package input

import "testing"

func TestCreateInput(t *testing.T) {
    result := CreateInput("someString")
    if(result.content != "someString") {
        t.Errorf("Expected content to be <someString>, but was %s", result.content)
    }

    if(result.position != 0) {
        t.Errorf("Expected position to be 0, but was %d", result.position)
    }
}
