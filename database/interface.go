package database

import (
	"github.com/tetha/threaddumpstorage-go/input"
	"github.com/tetha/threaddumpstorage-go/model"
)

/*DataStorage is a simple interface for databases or other storages */
type DataStorage interface {
	ListAllThreadHeadersInDump(threaddumpID string) ([]model.JavaThreadHeader, error)
	ListPagedThreadHeaders(threaddumpID string, limit int, offset int) ([]model.JavaThreadHeader, error)

	ListAllThreaddumps() ([]model.Threaddump, error)

	StoreDump(application, host string, dump input.Threaddump) (string, error)
}
