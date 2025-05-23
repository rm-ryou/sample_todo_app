package repositories

import (
	"database/sql"

	"github.com/rm-ryou/sample_todo_app/internal/entities"
)

type RoomRepository struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) *RoomRepository {
	return &RoomRepository{
		db: db,
	}
}

func (rr *RoomRepository) GetAll() ([]*entities.Room, error) {
	query := "SELECT id, name, created_at, updated_at FROM rooms"

	stmt, err := rr.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var rooms []*entities.Room
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var r entities.Room
		if err := rows.Scan(
			&r.Id,
			&r.Name,
			&r.CreatedAt,
			&r.UpdatedAt,
		); err != nil {
			return nil, err
		}
		rooms = append(rooms, &r)
	}

	return rooms, nil
}

func (rr *RoomRepository) Create(room *entities.Room) error {
	query := "INSERT INTO rooms (name) VALUES (?)"

	stmt, err := rr.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(room.Name)
	if err != nil {
		return err
	}

	return nil
}

func (rr *RoomRepository) Update(room *entities.Room) error {
	query := "UPDATE rooms SET name = ? WHERE id = ?"

	stmt, err := rr.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(room.Name, room.Id)
	if err != nil {
		return err
	}

	return nil
}

func (rr *RoomRepository) Delete(id int) error {
	query := "DELETE FROM rooms WHERE id = ?"

	_, err := rr.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
