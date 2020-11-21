package courier

import "github.com/fakorede-bolu/deliva/pkg/helpers"

type ID = helpers.ID

func NewID() ID {
	n := helpers.NewID()
	return n
}
