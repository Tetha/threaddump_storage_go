package database

import (
	"github.com/tetha/threaddumpstorage-go/model"
)

/*DataStorage is a simple interface for databases or other storages */
type DataStorage interface {
	ListAllThreadHeadersInDump(threaddumpID string) ([]model.JavaThreadHeader, error)
}
