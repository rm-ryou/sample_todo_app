package entities

import (
	"errors"
	"time"
)

type Board struct {
	Id        int
	Name      string
	Priority  int
	RoomId    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBoard(name string, priority, roomId int) *Board {
	return &Board{
		Name:     name,
		Priority: priority,
		RoomId:   roomId,
	}
}

func (b *Board) Validate() error {
	if b.Name == "" || len(b.Name) > 50 {
		return errors.New("Invalid name")
	}

	if b.Priority < 0 {
		return errors.New("Invalid priority size")
	}

	return nil
}

func (b *Board) UpdateAttributes(name string, priority int) {
	b.Name = name
	b.Priority = priority
}
