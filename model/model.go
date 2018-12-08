package model

import "time"

/*A Threaddump contains the metadata about a single dump
 */
type Threaddump struct {
	ID          int       `json:"id"`
	Application string    `json:"application"`
	Host        string    `json:"host"`
	Uploaded    time.Time `json:"uploaded"`
}

/*A JavaThreadHeader corresponds to one java thread in the threaddump.
 *
 * This javaThread only provides the information available in the
 * thread header. It does not immediately provide the stacktraces
 * as well, since that would make iterating threads a lot more
 * costly.
 */
type JavaThreadHeader struct {
	ID                     int
	Name                   string
	JavaID                 string
	IsDaemon               bool
	Prio                   int
	OsPrio                 int
	Tid                    string
	Nid                    string
	NativeThreadState      string
	ConditionAddress       string
	JavaThreadState        string
	JavaStateClarification string
}
