package entities

import (
	"deliva/pkg/helpers"
	"github.com/google/uuid"
)

type ID = helpers.ID

func NewID() ID {
	n := helpers.NewID()
	return n
}

//StringToID convert a string to an entity ID
func StringToID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}