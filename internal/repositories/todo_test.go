package repositories

import (
	"database/sql"
	"testing"
	"time"

	"github.com/rm-ryou/sample_todo_app/internal/config"
	"github.com/rm-ryou/sample_todo_app/internal/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setup(t *testing.T) *TodoRepository {
	cfg := config.DB{
		Database: MYSQL_DATABASE,
		User:     MYSQL_USER,
		Password: MYSQL_PASSWORD,
		Host:     MYSQL_HOST,
		Port:     MYSQL_PORT,
	}

	db, err := SetupConnection(cfg)
	require.NoError(t, err)

	return NewTodoRepository(db)
}

func getTodoCount(t *testing.T, repo *TodoRepository) int {
	var count int

	query := "SELECT COUNT(*) FROM todos"
	err := repo.db.QueryRow(query).Scan(&count)
	require.NoError(t, err)

	return count
}

func getTodoById(t *testing.T, repo *TodoRepository, id int) *entities.Todo {
	var todo entities.Todo
	query := "SELECT * FROM todos WHERE id = ?"

	err := repo.db.QueryRow(query, id).Scan(
		&todo.Id,
		&todo.Title,
		&todo.Done,
		&todo.Priority,
		&todo.DueDate,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	require.NoError(t, err)

	return &todo
}

func insertDummyTodo(t *testing.T, repo *TodoRepository, todo *entities.Todo) {
	query := `INSERT INTO todos
		(id, title, done, priority, due_date, created_at, updated_at)
	VALUES
		(?, ?, ?, ?, ?, ?, ?)
	`

	res, err := repo.db.Exec(query, todo.Id, todo.Title, todo.Done, todo.Priority, todo.DueDate, todo.CreatedAt, todo.UpdatedAt)
	require.NoError(t, err)
	_, err = res.LastInsertId()
	require.NoError(t, err)
}

func deleteAllTodos(t *testing.T, repo *TodoRepository) {
	query := "DELETE FROM todos"
	_, err := repo.db.Exec(query)
	require.NoError(t, err)
}

func TestGetAll(t *testing.T) {
	repo := setup(t)

	testCases := []struct {
		name          string
		savedTodos    []*entities.Todo
		setup         func(t *testing.T, todos []*entities.Todo)
		expectedError error
		expectedData  []*entities.Todo
	}{
		{
			name: "Success to Get todo",
			savedTodos: []*entities.Todo{
				{
					Id:        1,
					Title:     "done task",
					Done:      true,
					Priority:  0,
					CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
				},
				{
					Id:        2,
					Title:     "Test Task",
					Done:      false,
					Priority:  3,
					CreatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
				},
			},
			setup: func(t *testing.T, todos []*entities.Todo) {
				for _, todo := range todos {
					insertDummyTodo(t, repo, todo)
				}
			},
			expectedError: nil,
			expectedData: []*entities.Todo{
				{
					Id:        1,
					Title:     "done task",
					Done:      true,
					Priority:  0,
					CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
				},
				{
					Id:        2,
					Title:     "Test Task",
					Done:      false,
					Priority:  3,
					CreatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name:          "Returns nil if there is no record",
			savedTodos:    nil,
			setup:         func(t *testing.T, todos []*entities.Todo) {},
			expectedError: nil,
			expectedData:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t, tc.savedTodos)
			defer deleteAllTodos(t, repo)

			todo, err := repo.GetAll()

			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedData, todo)
		})
	}
}

func TestGetById(t *testing.T) {
	repo := setup(t)

	testCases := []struct {
		name          string
		id            int
		savedTodo     *entities.Todo
		setup         func(t *testing.T, todo *entities.Todo)
		expectedError error
		expectedData  *entities.Todo
	}{
		{
			name: "Success to Get todo",
			id:   1,
			savedTodo: &entities.Todo{
				Id:        1,
				Title:     "done task",
				Done:      true,
				Priority:  0,
				CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
			},
			setup: func(t *testing.T, todo *entities.Todo) {
				insertDummyTodo(t, repo, todo)
			},
			expectedError: nil,
			expectedData: &entities.Todo{
				Id:        1,
				Title:     "done task",
				Done:      true,
				Priority:  0,
				CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
			},
		},
		{
			name:          "Failed to Get todo with item not exists",
			id:            999,
			savedTodo:     nil,
			setup:         func(t *testing.T, todo *entities.Todo) {},
			expectedError: sql.ErrNoRows,
			expectedData:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t, tc.savedTodo)
			defer deleteAllTodos(t, repo)

			todo, err := repo.GetById(tc.id)

			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedData, todo)
		})
	}
}

func TestCreate(t *testing.T) {
	repo := setup(t)
	beforeCount := getTodoCount(t, repo)

	testCases := []struct {
		name                string
		todo                *entities.Todo
		expectedError       error
		expectedRecordCount int
	}{
		{
			name: "Success to Create todo",
			todo: &entities.Todo{
				Title:    "Test!!",
				Done:     false,
				Priority: 1,
			},
			expectedError:       nil,
			expectedRecordCount: beforeCount + 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer deleteAllTodos(t, repo)

			err := repo.Create(tc.todo)

			afterCount := getTodoCount(t, repo)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedRecordCount, afterCount)
		})
	}
}

func assertTodoHelper(t *testing.T, expected, actual *entities.Todo) {
	t.Helper()

	assert.Equal(t, expected.Id, actual.Id)
	assert.Equal(t, expected.Title, actual.Title)
	assert.Equal(t, expected.Done, actual.Done)
	assert.Equal(t, expected.Priority, actual.Priority)
	assert.Equal(t, expected.DueDate, actual.DueDate)
	assert.Equal(t, expected.CreatedAt, actual.CreatedAt)
}

func TestUpdate(t *testing.T) {
	repo := setup(t)

	testCases := []struct {
		name          string
		savedData     *entities.Todo
		updateData    *entities.Todo
		setup         func(t *testing.T, todo *entities.Todo)
		expectedError error
		expectedData  *entities.Todo
	}{
		{
			name: "Success to update todo",
			savedData: &entities.Todo{
				Id:        1,
				Title:     "done task",
				Done:      true,
				Priority:  0,
				CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
			},
			updateData: &entities.Todo{
				Id:        1,
				Title:     "updated!!",
				Done:      true,
				Priority:  0,
				CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
			},
			setup: func(t *testing.T, todo *entities.Todo) {
				insertDummyTodo(t, repo, todo)
			},
			expectedError: nil,
			expectedData: &entities.Todo{
				Id:        1,
				Title:     "updated!!",
				Done:      true,
				Priority:  0,
				CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
			},
		},
		{
			name:      "No changed when not exist id",
			savedData: &entities.Todo{},
			updateData: &entities.Todo{
				Id:        999,
				Title:     "updated!!",
				Done:      true,
				Priority:  0,
				CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
			},
			setup:         func(t *testing.T, todo *entities.Todo) {},
			expectedError: nil,
			expectedData:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t, tc.savedData)
			defer deleteAllTodos(t, repo)

			err := repo.Update(tc.updateData)

			assert.Equal(t, tc.expectedError, err)
			if tc.expectedData != nil {
				actual := getTodoById(t, repo, tc.savedData.Id)
				assertTodoHelper(t, tc.expectedData, actual)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	repo := setup(t)
	savedTodo := &entities.Todo{
		Id:        1,
		Title:     "done task",
		Done:      true,
		Priority:  0,
		CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
	}

	testCases := []struct {
		name                string
		deleteId            int
		setup               func(t *testing.T)
		expectedError       error
		expectedRecordCount int
	}{
		{
			name:     "Success to Delete todo",
			deleteId: savedTodo.Id,
			setup: func(t *testing.T) {
				insertDummyTodo(t, repo, savedTodo)
			},
			expectedError:       nil,
			expectedRecordCount: 0,
		},
		{
			name:     "No Error when not exist Id",
			deleteId: 999,
			setup: func(t *testing.T) {
				insertDummyTodo(t, repo, savedTodo)
			},
			expectedError:       nil,
			expectedRecordCount: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t)
			defer deleteAllTodos(t, repo)

			err := repo.Delete(tc.deleteId)

			afterCount := getTodoCount(t, repo)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedRecordCount, afterCount)
		})
	}
}
