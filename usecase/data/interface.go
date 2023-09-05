package data

import (
	"github.com/amardegan/cleanarch/entity"
)

type Data interface {
	TypeID() int
}

type Reader interface {
	Get(id int) (*entity.Data, error)
	List() ([]*entity.Data, error)
}

type Writer interface {
	Create(c *entity.Data) (int, error)
}

type Repository interface {
	Reader
	Writer
}
