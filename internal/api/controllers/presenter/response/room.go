package response

import (
	"time"

	"github.com/rm-ryou/sample_todo_app/internal/entities"
)

type ListRoom struct {
	Rooms []*Room `json:"rooms"`
}

type Room struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func convertRoomResponse(room *entities.Room) *Room {
	return &Room{
		Id:        room.Id,
		Name:      room.Name,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}
}

func ConvertRoomsResponse(rooms []*entities.Room) *ListRoom {
	listRoom := []*Room{}

	for _, room := range rooms {
		listRoom = append(listRoom, convertRoomResponse(room))
	}
	return &ListRoom{Rooms: listRoom}
}
