package handlers

import (
	"github.com/tetha/threaddumpstorage-go/database"
)

func NewEnvironment(db database.DataStorage) *RuntimeEnvironment {
	return &RuntimeEnvironment{db}
}

type RuntimeEnvironment struct {
	db database.DataStorage
}
