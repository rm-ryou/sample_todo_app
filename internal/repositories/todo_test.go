package repositories

import (
	"database/sql"
	"testing"
	"time"

	"github.com/rm-ryou/sample_todo_app/internal/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var referencedBoardData = entities.Board{
	Id:        1,
	Name:      "referencedRoom",
	Priority:  0,
	RoomId:    1,
	CreatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
	UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
}

func getTodoCount(t *testing.T) int {
	var count int

	query := "SELECT COUNT(*) FROM todos"
	err := TodoRepo.db.QueryRow(query).Scan(&count)
	require.NoError(t, err)

	return count
}

func getTodoById(t *testing.T, id int) *entities.Todo {
	var todo entities.Todo
	query := `SELECT
		id,
		title,
		done,
		priority,
		due_date,
		board_id,
		created_at,
		updated_at
	FROM
		todos
	WHERE id = ?`

	err := TodoRepo.db.QueryRow(query, id).Scan(
		&todo.Id,
		&todo.Title,
		&todo.Done,
		&todo.Priority,
		&todo.DueDate,
		&todo.BoardId,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	require.NoError(t, err)

	return &todo
}

func insertDummyTodo(t *testing.T, todo *entities.Todo) {
	query := `INSERT INTO todos
		(id, title, done, priority, due_date, board_id, created_at, updated_at)
	VALUES
		(?, ?, ?, ?, ?, ?, ?, ?)
	`

	res, err := TodoRepo.db.Exec(
		query,
		todo.Id,
		todo.Title,
		todo.Done,
		todo.Priority,
		todo.DueDate,
		todo.BoardId,
		todo.CreatedAt,
		todo.UpdatedAt,
	)
	require.NoError(t, err)
	_, err = res.LastInsertId()
	require.NoError(t, err)
}

func deleteAllTodos(t *testing.T) {
	query := "DELETE FROM todos"
	_, err := TodoRepo.db.Exec(query)
	require.NoError(t, err)
}

func TestGetByIdTodo(t *testing.T) {
	insertDummyRoom(t, &referencedRoomData)
	defer deleteAllRooms(t)
	insertDummyBoard(t, &referencedBoardData)
	defer deleteAllBoards(t)

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
				BoardId:   1,
				Title:     "done task",
				Done:      true,
				Priority:  0,
				CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
			},
			setup: func(t *testing.T, todo *entities.Todo) {
				insertDummyTodo(t, todo)
			},
			expectedError: nil,
			expectedData: &entities.Todo{
				Id:        1,
				BoardId:   1,
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
			defer deleteAllTodos(t)

			todo, err := TodoRepo.GetById(tc.id)

			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedData, todo)
		})
	}
}

func TestCreateTodo(t *testing.T) {
	insertDummyRoom(t, &referencedRoomData)
	defer deleteAllRooms(t)
	insertDummyBoard(t, &referencedBoardData)
	defer deleteAllBoards(t)

	beforeCount := getTodoCount(t)

	testCases := []struct {
		name                string
		todo                *entities.Todo
		expectedError       error
		expectedRecordCount int
	}{
		{
			name: "Success to Create todo",
			todo: &entities.Todo{
				Title:    "test todo",
				Done:     false,
				Priority: 1,
				BoardId:  1,
			},
			expectedError:       nil,
			expectedRecordCount: beforeCount + 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer deleteAllTodos(t)

			err := TodoRepo.Create(tc.todo)

			afterCount := getTodoCount(t)
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
	assert.Equal(t, expected.BoardId, actual.BoardId)
	assert.Equal(t, expected.CreatedAt, actual.CreatedAt)
}

func TestUpdateTodo(t *testing.T) {
	insertDummyRoom(t, &referencedRoomData)
	defer deleteAllRooms(t)
	insertDummyBoard(t, &referencedBoardData)
	defer deleteAllBoards(t)

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
				BoardId:   1,
				CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
			},
			updateData: &entities.Todo{
				Id:        1,
				Title:     "updated!!",
				Done:      true,
				Priority:  0,
				BoardId:   1,
				CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
			},
			setup: func(t *testing.T, todo *entities.Todo) {
				insertDummyTodo(t, todo)
			},
			expectedError: nil,
			expectedData: &entities.Todo{
				Id:        1,
				Title:     "updated!!",
				Done:      true,
				Priority:  0,
				BoardId:   1,
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
				BoardId:   1,
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
			defer deleteAllTodos(t)

			err := TodoRepo.Update(tc.updateData)

			assert.Equal(t, tc.expectedError, err)
			if tc.expectedData != nil {
				actual := getTodoById(t, tc.savedData.Id)
				assertTodoHelper(t, tc.expectedData, actual)
			}
		})
	}
}

func TestDeleteTodo(t *testing.T) {
	insertDummyRoom(t, &referencedRoomData)
	defer deleteAllRooms(t)
	insertDummyBoard(t, &referencedBoardData)
	defer deleteAllBoards(t)

	savedTodo := &entities.Todo{
		Id:        1,
		Title:     "done task",
		Done:      true,
		Priority:  0,
		BoardId:   1,
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
				insertDummyTodo(t, savedTodo)
			},
			expectedError:       nil,
			expectedRecordCount: 0,
		},
		{
			name:     "No Error when not exist Id",
			deleteId: 999,
			setup: func(t *testing.T) {
				insertDummyTodo(t, savedTodo)
			},
			expectedError:       nil,
			expectedRecordCount: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t)
			defer deleteAllTodos(t)

			err := TodoRepo.Delete(tc.deleteId)

			afterCount := getTodoCount(t)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedRecordCount, afterCount)
		})
	}
}
