package entities

import "time"

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
