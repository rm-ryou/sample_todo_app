package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}
