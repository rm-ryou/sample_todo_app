package repositories

import (
	"database/sql"
	"testing"
	"time"

	"github.com/rm-ryou/sample_todo_app/internal/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var referencedRoomData = entities.Room{
	Id:        1,
	Name:      "referencedRoom",
	CreatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
	UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
}

func getBoardCount(t *testing.T) int {
	var count int

	query := "SELECT COUNT(*) FROM boards"
	err := BoardRepo.db.QueryRow(query).Scan(&count)
	require.NoError(t, err)

	return count
}

func getBoardById(t *testing.T, id int) *entities.Board {
	var board entities.Board
	query := "SELECT * FROM boards WHERE id = ?"

	err := BoardRepo.db.QueryRow(query, id).Scan(
		&board.Id,
		&board.RoomId,
		&board.Name,
		&board.Priority,
		&board.CreatedAt,
		&board.UpdatedAt,
	)
	require.NoError(t, err)

	return &board
}

func insertDummyBoard(t *testing.T, board *entities.Board) {
	query := `INSERT INTO boards
		(id, room_id, name, priority, created_at, updated_at)
	VALUES
		(?, ?, ?, ?, ?, ?)
	`

	_, err := BoardRepo.db.Exec(
		query,
		board.Id,
		board.RoomId,
		board.Name,
		board.Priority,
		board.CreatedAt,
		board.UpdatedAt,
	)
	require.NoError(t, err)
}

func deleteAllBoards(t *testing.T) {
	query := "DELETE FROM boards"
	_, err := BoardRepo.db.Exec(query)
	require.NoError(t, err)
}

func TestGetAllBoards(t *testing.T) {
	insertDummyRoom(t, &referencedRoomData)
	defer deleteAllRooms(t)

	testCases := []struct {
		name          string
		savedBoard    []*entities.Board
		setup         func(t *testing.T, board []*entities.Board)
		expectedError error
		expectedData  []*entities.Board
	}{
		{
			name: "Success to Get board",
			savedBoard: []*entities.Board{
				{
					Id:        1,
					Name:      "test board",
					Priority:  0,
					RoomId:    1,
					CreatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
				},
				{
					Id:        2,
					Name:      "sample board",
					Priority:  0,
					RoomId:    1,
					CreatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
				},
			},
			setup: func(t *testing.T, boards []*entities.Board) {
				for _, board := range boards {
					insertDummyBoard(t, board)
				}
			},
			expectedError: nil,
			expectedData: []*entities.Board{
				{
					Id:        1,
					Name:      "test board",
					Priority:  0,
					RoomId:    1,
					CreatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
				},
				{
					Id:        2,
					Name:      "sample board",
					Priority:  0,
					RoomId:    1,
					CreatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name:          "Returns nil if there is no record",
			savedBoard:    nil,
			setup:         func(t *testing.T, boards []*entities.Board) {},
			expectedError: nil,
			expectedData:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t, tc.savedBoard)
			defer deleteAllBoards(t)

			rooms, err := BoardRepo.GetAll()

			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedData, rooms)
		})
	}
}

func TestGetByIdBoard(t *testing.T) {
	insertDummyRoom(t, &referencedRoomData)
	defer deleteAllRooms(t)

	testCases := []struct {
		name          string
		id            int
		savedBoard    *entities.Board
		setup         func(t *testing.T, board *entities.Board)
		expectedError error
		expectedData  *entities.Board
	}{
		{
			name: "Success to Get board",
			id:   1,
			savedBoard: &entities.Board{
				Id:        1,
				Name:      "test board",
				Priority:  0,
				RoomId:    1,
				CreatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
			},
			setup: func(t *testing.T, board *entities.Board) {
				insertDummyBoard(t, board)
			},
			expectedError: nil,
			expectedData: &entities.Board{
				Id:        1,
				Name:      "test board",
				Priority:  0,
				RoomId:    1,
				CreatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
			},
		},
		{
			name:          "Failed to Get board with item not exists",
			id:            999,
			savedBoard:    nil,
			setup:         func(t *testing.T, board *entities.Board) {},
			expectedError: sql.ErrNoRows,
			expectedData:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t, tc.savedBoard)
			defer deleteAllBoards(t)

			room, err := BoardRepo.GetById(tc.id)

			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedData, room)
		})
	}
}

func TestCreateBoard(t *testing.T) {
	insertDummyRoom(t, &referencedRoomData)
	defer deleteAllRooms(t)

	beforeCount := getBoardCount(t)

	testCases := []struct {
		name                string
		board               *entities.Board
		expectedError       error
		expectedRecordCount int
	}{
		{
			name: "Success to Create room",
			board: &entities.Board{
				Name:     "test board",
				Priority: 0,
				RoomId:   1,
			},
			expectedError:       nil,
			expectedRecordCount: beforeCount + 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer deleteAllBoards(t)

			err := BoardRepo.Create(tc.board)

			afterCount := getBoardCount(t)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedRecordCount, afterCount)
		})
	}
}

func assertBoardHelper(t *testing.T, expected, actual *entities.Board) {
	t.Helper()

	assert.Equal(t, expected.Id, actual.Id)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Priority, actual.Priority)
	assert.Equal(t, expected.RoomId, actual.RoomId)
	assert.Equal(t, expected.CreatedAt, actual.CreatedAt)
}

func TestUpdateBoard(t *testing.T) {
	insertDummyRoom(t, &referencedRoomData)
	defer deleteAllRooms(t)

	testCases := []struct {
		name          string
		savedData     *entities.Board
		updateData    *entities.Board
		setup         func(t *testing.T, board *entities.Board)
		expectedError error
		expectedData  *entities.Board
	}{
		{
			name: "Success to update board",
			savedData: &entities.Board{
				Id:        1,
				Name:      "test board",
				Priority:  0,
				RoomId:    1,
				CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
			},
			updateData: &entities.Board{
				Id:        1,
				Name:      "update board",
				Priority:  2,
				RoomId:    1,
				CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
			},
			setup: func(t *testing.T, board *entities.Board) {
				insertDummyBoard(t, board)
			},
			expectedError: nil,
			expectedData: &entities.Board{
				Id:        1,
				Name:      "update board",
				Priority:  2,
				RoomId:    1,
				CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
			},
		},
		{
			name:      "No changed when not exist id",
			savedData: &entities.Board{},
			updateData: &entities.Board{
				Id:        999,
				Name:      "update board",
				Priority:  2,
				RoomId:    1,
				CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
			},
			setup:         func(t *testing.T, board *entities.Board) {},
			expectedError: nil,
			expectedData:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t, tc.savedData)
			defer deleteAllBoards(t)

			err := BoardRepo.Update(tc.updateData)

			assert.Equal(t, tc.expectedError, err)
			if tc.expectedData != nil {
				actual := getBoardById(t, tc.savedData.Id)
				assertBoardHelper(t, tc.expectedData, actual)
			}
		})
	}
}

func TestDeleteBoard(t *testing.T) {
	insertDummyRoom(t, &referencedRoomData)
	defer deleteAllRooms(t)

	savedBoard := &entities.Board{
		Id:        1,
		Name:      "test Board",
		Priority:  0,
		RoomId:    1,
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
			name:     "Success to Delete board",
			deleteId: savedBoard.Id,
			setup: func(t *testing.T) {
				insertDummyBoard(t, savedBoard)
			},
			expectedError:       nil,
			expectedRecordCount: 0,
		},
		{
			name:     "No Error when not exist Id",
			deleteId: 999,
			setup: func(t *testing.T) {
				insertDummyBoard(t, savedBoard)
			},
			expectedError:       nil,
			expectedRecordCount: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t)
			defer deleteAllBoards(t)

			err := BoardRepo.Delete(tc.deleteId)

			afterCount := getBoardCount(t)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedRecordCount, afterCount)
		})
	}
}
