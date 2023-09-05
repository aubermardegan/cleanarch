package circle

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

func (s *Service) GetCircle(id int) (*shapes.Circle, error) {
	data, err := s.repo.Get(id)
	if data == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	circle := shapes.Circle{}
	if data.TypeID != circle.TypeID() {
		return nil, entity.ErrTypeMismatch
	}

	err = json.NewDecoder(strings.NewReader(data.Value)).Decode(&circle)
	if err != nil {
		return nil, err
	}

	return &circle, nil
}

func (s *Service) CreateCircle(CanRoll bool, Radius int) (int, error) {
	c, err := shapes.NewCircle(CanRoll, Radius)
	if err != nil {
		return 0, err
	}

	jsonCircle, err := json.Marshal(c)
	if err != nil {
		return 0, err
	}

	data := entity.NewData(c.TypeID(), string(jsonCircle))

	_, err = s.GetCircle(data.ID)
	if err == entity.ErrNotFound {
		return s.repo.Create(data)
	}
	return data.ID, entity.ErrDuplicatedEntity
}

func (s *Service) ListCircles() ([]*shapes.Circle, error) {
	data, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, entity.ErrNotFound
	}

	circles := []*shapes.Circle{}
	for _, value := range data {
		circle, err := s.GetCircle(value.ID)
		if err == nil {
			circles = append(circles, circle)
		}
	}
	return circles, nil
}
