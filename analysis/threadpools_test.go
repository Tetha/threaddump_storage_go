package analysis

import (
	"reflect"
	"testing"

	"github.com/tetha/threaddumpstorage-go/model"
)

var threadPoolTests = map[string]struct {
	inputThreads     []model.JavaThreadHeader
	expectedPools    map[string][]model.JavaThreadHeader
	remainingThreads []model.JavaThreadHeader
}{
	"simple ajp threads": {
		[]model.JavaThreadHeader{
			{Name: "ajp-8080-1"},
			{Name: "ajp-8080-2"},
		},
		map[string][]model.JavaThreadHeader{
			"ajp-8080": {
				{Name: "ajp-8080-1"},
				{Name: "ajp-8080-2"},
			},
		},
		nil,
	},
	"elasticsearch threads": {
		[]model.JavaThreadHeader{
			{Name: "elasticsearch[Vindaloo][generic][T#1471]"},
		},
		map[string][]model.JavaThreadHeader{
			"elasticsearch (instance=Vindaloo, pool=generic)": {
				{Name: "elasticsearch[Vindaloo][generic][T#1471]"},
			},
		},
		nil,
	},
}

func TestFigureOutThreadpools(t *testing.T) {
	for name, tt := range threadPoolTests {
		actualPools := FigureOutThreadpools(tt.inputThreads)
		if !reflect.DeepEqual(actualPools.ThreadPools, tt.expectedPools) {
			t.Errorf("%s failed, expected pools %v, got %v", name, tt.expectedPools, actualPools.ThreadPools)
		}
		if !reflect.DeepEqual(actualPools.UnknownThreads, tt.remainingThreads) {
			t.Errorf("%s failed, expected remaining threads %v, got %v", name, tt.remainingThreads, actualPools.UnknownThreads)
		}
	}
}
