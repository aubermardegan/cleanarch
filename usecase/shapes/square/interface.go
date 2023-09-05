package square

import (
	"github.com/amardegan/cleanarch/entity/shapes"
)

type UseCase interface {
	GetSquare(id string) (*shapes.Circle, error)
	CreateSquare(Axis int, Area int) (int, error)
	ListSquares() (*shapes.Square, error)
}
