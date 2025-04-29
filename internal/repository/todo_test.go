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

func TestGet(t *testing.T) {
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
			todo, err := repo.Get(tc.id)
			assert.Equal(t, err, tc.expectedError)
			assert.Equal(t, todo, tc.expectedData)
		})
	}
}
