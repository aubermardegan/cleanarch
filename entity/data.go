package entity

import "math/rand"

type Data struct {
	ID     int
	TypeID int
	Value  string
}

func NewData(typeID int, value string) *Data {
	d := &Data{
		ID:     rand.Int(),
		TypeID: typeID,
		Value:  value,
	}
	return d
}
