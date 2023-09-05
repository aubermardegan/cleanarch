package circle

import (
	"github.com/amardegan/cleanarch/entity/shapes"
)

type UseCase interface {
	GetCircle(id string) (*shapes.Circle, error)
	CreateCircle(CanRoll bool, Radius int) (int, error)
	ListCircles() (*shapes.Circle, error)
}
