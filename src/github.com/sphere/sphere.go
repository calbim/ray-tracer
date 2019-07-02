package sphere

import (
	"errors"

	uuid "github.com/nu7hatch/gouuid"
)

// Sphere represents a unique sphere
type Sphere struct {
	id string
}

// New returns a new sphere
func New() Sphere {
	id, err := uuid.NewV4()
	if err != nil {
		errors.New("failed to generate a unique identifier for sphere")
	}
	return Sphere{
		id: id.String(),
	}
}
