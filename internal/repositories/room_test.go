package repositories

import (
	"database/sql"
	"testing"
	"time"

	"github.com/rm-ryou/sample_todo_app/internal/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getRoomCount(t *testing.T) int {
	var count int

	query := "SELECT COUNT(*) FROM rooms"
	err := RoomRepo.db.QueryRow(query).Scan(&count)
	require.NoError(t, err)

	return count
}

func getRoomById(t *testing.T, id int) *entities.Room {
	var room entities.Room
	query := "SELECT * FROM rooms WHERE id = ?"

	err := RoomRepo.db.QueryRow(query, id).Scan(
		&room.Id,
		&room.Name,
		&room.CreatedAt,
		&room.UpdatedAt,
	)
	require.NoError(t, err)

	return &room
}

func insertDummyRoom(t *testing.T, room *entities.Room) {
	query := `INSERT INTO rooms
		(id, name, created_at, updated_at)
	VALUES
		(?, ?, ?, ?)
	`

	res, err := RoomRepo.db.Exec(query, room.Id, room.Name, room.CreatedAt, room.UpdatedAt)
	require.NoError(t, err)
	_, err = res.LastInsertId()
	require.NoError(t, err)
}

func deleteAllRooms(t *testing.T) {
	query := "DELETE FROM rooms"
	_, err := RoomRepo.db.Exec(query)
	require.NoError(t, err)
}

func TestGetAllRooms(t *testing.T) {
	testCases := []struct {
		name          string
		savedRooms    []*entities.Room
		setup         func(t *testing.T, todos []*entities.Room)
		expectedError error
		expectedData  []*entities.Room
	}{
		{
			name: "Success to Get room",
			savedRooms: []*entities.Room{
				{
					Id:        1,
					Name:      "test room",
					CreatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
				},
				{
					Id:        2,
					Name:      "sample room",
					CreatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
				},
			},
			setup: func(t *testing.T, rooms []*entities.Room) {
				for _, room := range rooms {
					insertDummyRoom(t, room)
				}
			},
			expectedError: nil,
			expectedData: []*entities.Room{
				{
					Id:        1,
					Name:      "test room",
					CreatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
				},
				{
					Id:        2,
					Name:      "sample room",
					CreatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name:          "Returns nil if there is no record",
			savedRooms:    nil,
			setup:         func(t *testing.T, rooms []*entities.Room) {},
			expectedError: nil,
			expectedData:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t, tc.savedRooms)
			defer deleteAllRooms(t)

			rooms, err := RoomRepo.GetAll()

			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedData, rooms)
		})
	}
}

func TestGetByIdRoom(t *testing.T) {
	testCases := []struct {
		name          string
		id            int
		savedRoom     *entities.Room
		setup         func(t *testing.T, room *entities.Room)
		expectedError error
		expectedData  *entities.Room
	}{
		{
			name: "Success to Get room",
			id:   1,
			savedRoom: &entities.Room{
				Id:        1,
				Name:      "test room",
				CreatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
			},
			setup: func(t *testing.T, room *entities.Room) {
				insertDummyRoom(t, room)
			},
			expectedError: nil,
			expectedData: &entities.Room{
				Id:        1,
				Name:      "test room",
				CreatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
			},
		},
		{
			name:          "Failed to Get room with item not exists",
			id:            999,
			savedRoom:     nil,
			setup:         func(t *testing.T, room *entities.Room) {},
			expectedError: sql.ErrNoRows,
			expectedData:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t, tc.savedRoom)
			defer deleteAllRooms(t)

			room, err := RoomRepo.GetById(tc.id)

			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedData, room)
		})
	}
}

func TestCreateRoom(t *testing.T) {
	beforeCount := getRoomCount(t)

	testCases := []struct {
		name                string
		room                *entities.Room
		expectedError       error
		expectedRecordCount int
	}{
		{
			name:                "Success to Create room",
			room:                &entities.Room{Name: "Test!!"},
			expectedError:       nil,
			expectedRecordCount: beforeCount + 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer deleteAllRooms(t)

			err := RoomRepo.Create(tc.room)

			afterCount := getRoomCount(t)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedRecordCount, afterCount)
		})
	}
}

func assertRoomHelper(t *testing.T, expected, actual *entities.Room) {
	t.Helper()

	assert.Equal(t, expected.Id, actual.Id)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.CreatedAt, actual.CreatedAt)
}

func TestUpdateRoom(t *testing.T) {
	testCases := []struct {
		name          string
		savedData     *entities.Room
		updateData    *entities.Room
		setup         func(t *testing.T, room *entities.Room)
		expectedError error
		expectedData  *entities.Room
	}{
		{
			name: "Success to update room",
			savedData: &entities.Room{
				Id:        1,
				Name:      "test room",
				CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
			},
			updateData: &entities.Room{
				Id:        1,
				Name:      "updated!!",
				CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
			},
			setup: func(t *testing.T, room *entities.Room) {
				insertDummyRoom(t, room)
			},
			expectedError: nil,
			expectedData: &entities.Room{
				Id:        1,
				Name:      "updated!!",
				CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
			},
		},
		{
			name:      "No changed when not exist id",
			savedData: &entities.Room{},
			updateData: &entities.Room{
				Id:        999,
				Name:      "updated!!",
				CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
			},
			setup:         func(t *testing.T, todo *entities.Room) {},
			expectedError: nil,
			expectedData:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t, tc.savedData)
			defer deleteAllRooms(t)

			err := RoomRepo.Update(tc.updateData)

			assert.Equal(t, tc.expectedError, err)
			if tc.expectedData != nil {
				actual := getRoomById(t, tc.savedData.Id)
				assertRoomHelper(t, tc.expectedData, actual)
			}
		})
	}
}

func TestDeleteRoom(t *testing.T) {
	savedRoom := &entities.Room{
		Id:        1,
		Name:      "test room",
		CreatedAt: time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
	}
	relatedBoard := &entities.Board{
		Id:        1,
		Name:      "test board",
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
			name:     "Success to Delete room. the board associated with it is also deleted",
			deleteId: savedRoom.Id,
			setup: func(t *testing.T) {
				insertDummyRoom(t, savedRoom)
				insertDummyBoard(t, relatedBoard)
			},
			expectedError:       nil,
			expectedRecordCount: 0,
		},
		{
			name:     "No Error when not exist Id",
			deleteId: 999,
			setup: func(t *testing.T) {
				insertDummyRoom(t, savedRoom)
				insertDummyBoard(t, relatedBoard)
			},
			expectedError:       nil,
			expectedRecordCount: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t)
			defer deleteAllRooms(t)
			defer deleteAllBoards(t)

			err := RoomRepo.Delete(tc.deleteId)

			afterRoomCount := getRoomCount(t)
			afterBoardCount := getBoardCount(t)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedRecordCount, afterRoomCount)
			assert.Equal(t, tc.expectedRecordCount, afterBoardCount)
		})
	}
}
