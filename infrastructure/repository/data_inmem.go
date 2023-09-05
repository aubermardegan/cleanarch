package repository

import "github.com/amardegan/cleanarch/entity"

type InMem struct {
	m map[int]*entity.Data
}

func NewInMem() *InMem {
	var m = map[int]*entity.Data{}
	return &InMem{
		m: m,
	}
}

func (r *InMem) Create(e *entity.Data) (int, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

func (r *InMem) Get(id int) (*entity.Data, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

func (r *InMem) List() ([]*entity.Data, error) {
	var d []*entity.Data
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}
