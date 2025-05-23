package repositories

import (
	"database/sql"

	"github.com/rm-ryou/sample_todo_app/internal/entities"
)

type BoardRepository struct {
	db *sql.DB
}

func NewBoardRepository(db *sql.DB) *BoardRepository {
	return &BoardRepository{
		db: db,
	}
}

func (br *BoardRepository) GetAll() ([]*entities.Board, error) {
	query := "SELECT id, name, priority, room_id, created_at, updated_at FROM boards"

	stmt, err := br.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var boards []*entities.Board
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var b entities.Board
		if err := rows.Scan(
			&b.Id,
			&b.Name,
			&b.Priority,
			&b.RoomId,
			&b.CreatedAt,
			&b.UpdatedAt,
		); err != nil {
			return nil, err
		}
		boards = append(boards, &b)
	}

	return boards, nil
}

func (br *BoardRepository) GetById(id int) (*entities.Board, error) {
	var board entities.Board
	query := "SELECT id, name, priority, room_id, created_at, updated_at FROM boards WHERE id = ?"

	if err := br.db.QueryRow(query, id).Scan(
		&board.Id,
		&board.Name,
		&board.Priority,
		&board.RoomId,
		&board.CreatedAt,
		&board.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &board, nil
}

func (br *BoardRepository) Create(board *entities.Board) error {
	query := "INSERT INTO boards (name, priority, room_id) VALUES (?, ?, ?)"

	stmt, err := br.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(board.Name, board.Priority, board.RoomId)
	if err != nil {
		return err
	}

	return nil
}

func (br *BoardRepository) Update(board *entities.Board) error {
	query := "UPDATE boards SET name = ?, priority = ? WHERE id = ?"

	stmt, err := br.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(board.Name, board.Priority, board.Id)
	if err != nil {
		return err
	}

	return nil
}

func (br *BoardRepository) Delete(id int) error {
	query := "DELETE FROM boards WHERE id = ?"

	_, err := br.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
