package entities

import (
	"errors"
	"time"
)

type Room struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewRoom(name string) *Room {
	return &Room{
		Name: name,
	}
}

func (r *Room) Validate() error {
	if r.Name == "" || len(r.Name) > 50 {
		return errors.New("Invalid name")
	}

	return nil
}

func (r *Room) UpdateAttributes(name string) {
	r.Name = name
}
