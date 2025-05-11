package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rm-ryou/sample_todo_app/internal/entity"
)

type Todo struct {
	db *sql.DB
}

func NewTodo(db *sql.DB) *Todo {
	return &Todo{
		db: db,
	}
}

func (t *Todo) GetById(id int) (*entity.Todo, error) {
	var todo entity.Todo
	query := `SELECT * FROM todos WHERE id = ?`

	if err := t.db.QueryRow(query, id).Scan(
		&todo.Id,
		&todo.Title,
		&todo.Done,
		&todo.Priority,
		&todo.DueDate,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *Todo) Create(todo *entity.Todo) error {
	query := `INSERT INTO todos (title, done, priority, due_date) VALUES (?, ?, ?, ?)`

	stmt, err := t.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Title, todo.Done, todo.Priority, todo.DueDate)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todo) Update(todo *entity.Todo) error {
	query := `UPDATE todos SET title = ?, done = ?, priority = ?, due_date = ? WHERE id = ?`

	stmt, err := t.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Title, todo.Done, todo.Priority, todo.DueDate, todo.Id)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todo) Delete(id int) error {
	query := `DELETE FROM todos WHERE id = ?`

	_, err := t.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
