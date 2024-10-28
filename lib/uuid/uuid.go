package uuid

import (
	"github.com/google/uuid"
)

func NewString() string {
	u, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	return u.String()
}
