package square

import (
	"encoding/json"
	"strings"

	"github.com/amardegan/cleanarch/entity"
	"github.com/amardegan/cleanarch/entity/shapes"
	"github.com/amardegan/cleanarch/usecase/data"
)

type Service struct {
	repo data.Repository
}

func NewService(r data.Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetSquare(id int) (*shapes.Square, error) {
	data, err := s.repo.Get(id)
	if data == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	square := shapes.Square{}
	if data.TypeID != square.TypeID() {
		return nil, entity.ErrTypeMismatch
	}

	err = json.NewDecoder(strings.NewReader(data.Value)).Decode(&square)
	if err != nil {
		return nil, err
	}

	return &square, nil
}

func (s *Service) CreateSquare(Axis int, Area float64) (int, error) {
	sq, err := shapes.NewSquare(Axis, Area)
	if err != nil {
		return 0, err
	}

	jsonSquare, err := json.Marshal(sq)
	if err != nil {
		return 0, err
	}

	data := entity.NewData(sq.TypeID(), string(jsonSquare))

	_, err = s.GetSquare(data.ID)
	if err == entity.ErrNotFound {
		return s.repo.Create(data)
	}
	return data.ID, entity.ErrDuplicatedEntity
}

func (s *Service) ListSquares() ([]*shapes.Square, error) {
	data, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, entity.ErrNotFound
	}

	squares := []*shapes.Square{}
	for _, value := range data {
		square, err := s.GetSquare(value.ID)
		if err == nil {
			squares = append(squares, square)
		}
	}
	return squares, nil
}
