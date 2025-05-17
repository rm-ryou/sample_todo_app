package repositories

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rm-ryou/sample_todo_app/internal/entities"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (tr *TodoRepository) GetById(id int) (*entities.Todo, error) {
	var todo entities.Todo
	query := `
	SELECT
		id,
		title,
		done,
		priority,
		due_date,
		created_at,
		updated_at
	FROM
		todos
	WHERE
		id = ?
	`

	if err := tr.db.QueryRow(query, id).Scan(
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

func (tr *TodoRepository) Create(todo *entities.Todo) error {
	query := "INSERT INTO todos (title, done, priority, due_date) VALUES (?, ?, ?, ?)"

	stmt, err := tr.db.Prepare(query)
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

func (tr *TodoRepository) Update(todo *entities.Todo) error {
	query := "UPDATE todos SET title = ?, done = ?, priority = ?, due_date = ? WHERE id = ?"

	stmt, err := tr.db.Prepare(query)
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

func (tr *TodoRepository) Delete(id int) error {
	query := "DELETE FROM todos WHERE id = ?"

	_, err := tr.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
