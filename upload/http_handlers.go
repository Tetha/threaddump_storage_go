package upload

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tetha/threaddumpstorage-go/input"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	unparsedDump, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//  TODO
		http.Error(w, "500: Error reading body", 500)
		log.Printf("Error  reading input body: %s", err)
		return
	}

	input := input.CreateInput(string(unparsedDump))
	parseFailure, dump := input.ParseThreaddump()
	if parseFailure != "" {
		log.Printf("Could not parse threaddumps")
		http.Error(w, "400: Could not parse dump: "+parseFailure, 400)
		return
	}
	log.Printf("Parsed input with %d threads", len(dump.Threads))
}
