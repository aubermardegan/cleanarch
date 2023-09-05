package shapes

import "github.com/amardegan/cleanarch/entity"

type Circle struct {
	CanRoll bool `json:"canroll"`
	Radius  int  `json:"radius"`
}

func (c *Circle) TypeID() int {
	return 2
}

func NewCircle(CanRoll bool, Radius int) (*Circle, error) {
	c := &Circle{
		CanRoll: CanRoll,
		Radius:  Radius,
	}
	err := c.Validate()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Circle) Validate() error {
	if c.Radius < 0 {
		return entity.ErrInvalidEntity
	}
	return nil
}
