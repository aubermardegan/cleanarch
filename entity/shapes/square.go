package shapes

import "github.com/amardegan/cleanarch/entity"

type Square struct {
	Axis int     `json:"axis"`
	Area float64 `json:"area"`
}

func (s *Square) TypeID() int {
	return 1
}

func NewSquare(Axis int, Area float64) (*Square, error) {
	s := &Square{Axis: Axis,
		Area: Area,
	}
	err := s.Validate()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Square) Validate() error {
	if s.Axis != 4 || s.Area < 0 {
		return entity.ErrInvalidEntity
	}
	return nil
}
