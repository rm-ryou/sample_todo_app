package repository

import (
	"database/sql"
	"testing"
	"time"

	"github.com/rm-ryou/sample_todo_app/internal/config"
	"github.com/rm-ryou/sample_todo_app/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setup(t *testing.T) *Todo {
	cfg := config.DB{
		Database: MYSQL_DATABASE,
		User:     MYSQL_USER,
		Password: MYSQL_PASSWORD,
		Host:     MYSQL_HOST,
		Port:     MYSQL_PORT,
	}

	db, err := SetupConnection(cfg)
	require.NoError(t, err)

	return NewTodo(db)
}

func fetchTodoRows(t *testing.T, repo *Todo) int {
	var todoRows int

	query := "SELECT COUNT(*) FROM todos"
	err := repo.db.QueryRow(query).Scan(&todoRows)
	require.NoError(t, err)

	return todoRows
}

func TestCreate(t *testing.T) {
	repo := setup(t)
	beforeTodoRows := fetchTodoRows(t, repo)

	testCases := []struct {
		name             string
		todo             *entity.Todo
		expectedTodoRows int
		expectedError    error
	}{
		{
			name: "Success to Create todo",
			todo: &entity.Todo{
				Title:    "Test!!",
				Done:     false,
				Priority: 1,
			},
			expectedTodoRows: beforeTodoRows + 1,
			expectedError:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.Create(tc.todo)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.Nil(t, err)

				todoRows := fetchTodoRows(t, repo)
				assert.Equal(t, tc.expectedTodoRows, todoRows)

				t.Cleanup(func() {
					query := `DELETE FROM todos WHERE title = ?`
					_, err := repo.db.Exec(query, tc.todo.Title)
					require.NoError(t, err)
				})
			}
		})
	}
}

func TestGetById(t *testing.T) {
	repo := setup(t)

	testCases := []struct {
		name          string
		id            int
		expectedError error
		expectedData  *entity.Todo
	}{
		{
			name:          "Success to Get todo",
			id:            1,
			expectedError: nil,
			expectedData: &entity.Todo{
				Id:        1,
				Title:     "Test Task",
				Done:      false,
				Priority:  3,
				CreatedAt: time.Date(2025, 5, 1, 10, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2025, 5, 1, 10, 0, 0, 0, time.UTC),
			},
		},
		{
			name:          "Failed to Get todo with item not exists",
			id:            999,
			expectedError: sql.ErrNoRows,
			expectedData:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			todo, err := repo.GetById(tc.id)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedData, todo)
		})
	}
}
