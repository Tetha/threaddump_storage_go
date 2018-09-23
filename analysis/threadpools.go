package analysis

import (
	"fmt"
	"regexp"

	"github.com/tetha/threaddumpstorage-go/model"
)

//ThreadPoolDetection contains the threadpools found in a set of headers
type ThreadPoolDetection struct {
	ThreadPools    map[string][]model.JavaThreadHeader
	UnknownThreads []model.JavaThreadHeader
}

//FigureOutThreadpools uses heuristics to find thread pools in the thread set
func FigureOutThreadpools(threads []model.JavaThreadHeader) (pools ThreadPoolDetection) {
	pools.ThreadPools = make(map[string][]model.JavaThreadHeader)
	pools.UnknownThreads = threads

	findElasticsearchThreadpools(&pools)
	findNumberedThreads(&pools)
	return
}

func findNumberedThreads(pools *ThreadPoolDetection) {
	var newUnknownThreads []model.JavaThreadHeader
	nameRegex := regexp.MustCompile(`(.+?)[-#]?\d+$`)
	for _, thread := range pools.UnknownThreads {
		if match := nameRegex.FindStringSubmatch(thread.Name); match != nil {
			readableName := match[1]
			pools.ThreadPools[readableName] = append(pools.ThreadPools[readableName], thread)
		} else {
			newUnknownThreads = append(newUnknownThreads, thread)
		}
	}
	pools.UnknownThreads = newUnknownThreads
}

func findElasticsearchThreadpools(pools *ThreadPoolDetection) {
	var newUnknownThreads []model.JavaThreadHeader
	nameRegex := regexp.MustCompile(`^elasticsearch\[([^]]+)\]\[([^]]+)\]`)
	for _, thread := range pools.UnknownThreads {
		if match := nameRegex.FindStringSubmatch(thread.Name); match != nil {
			readableName := fmt.Sprintf("elasticsearch (instance=%s, pool=%s)", match[1], match[2])
			pools.ThreadPools[readableName] = append(pools.ThreadPools[readableName], thread)
		} else {
			newUnknownThreads = append(newUnknownThreads, thread)
		}
	}
	pools.UnknownThreads = newUnknownThreads
}
