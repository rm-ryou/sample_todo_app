package entities

import "time"

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
