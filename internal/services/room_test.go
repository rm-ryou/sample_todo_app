package services

import (
	"database/sql"
	"errors"
	"strings"
	"testing"

	"github.com/rm-ryou/sample_todo_app/internal/entities"
	mock_repository "github.com/rm-ryou/sample_todo_app/internal/interfaces/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateRoom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockRoomRepository(ctrl)
	service := NewRoomService(mockRepository)

	testCases := []struct {
		name          string
		roomName      string
		mockSetup     func(room *entities.Room)
		expectedError error
	}{
		{
			name:     "Success to create room",
			roomName: "test room",
			mockSetup: func(room *entities.Room) {
				mockRepository.EXPECT().Create(room).
					Return(nil)
			},
			expectedError: nil,
		},
		{
			name:          "Failed to create room - Due to number of characters in the room is more than 50",
			roomName:      strings.Repeat("a", 51),
			mockSetup:     func(room *entities.Room) {},
			expectedError: errors.New("Invalid name"),
		},
		{
			name:          "Failed to create room - Due to the empty name",
			roomName:      "",
			mockSetup:     func(room *entities.Room) {},
			expectedError: errors.New("Invalid name"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			room := &entities.Room{
				Name: tc.roomName,
			}
			tc.mockSetup(room)

			err := service.Create(tc.roomName)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestUpdateRoom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockRoomRepository(ctrl)
	service := NewRoomService(mockRepository)

	testCases := []struct {
		name          string
		id            int
		roomName      string
		mockSetup     func(room *entities.Room)
		expectedError error
	}{
		{
			name:     "Success to update room",
			id:       1,
			roomName: "test room",
			mockSetup: func(room *entities.Room) {
				mockRepository.EXPECT().GetById(room.Id).
					Return(&entities.Room{Id: room.Id}, nil)
				mockRepository.EXPECT().Update(room).
					Return(nil)
			},
			expectedError: nil,
		},
		{
			name:     "Failed to update room - Due to the todo not found",
			id:       999,
			roomName: "test room",
			mockSetup: func(room *entities.Room) {
				mockRepository.EXPECT().GetById(room.Id).
					Return(nil, sql.ErrNoRows)
			},
			expectedError: sql.ErrNoRows,
		},
		{
			name:     "Failed to update room - Due to number of characters in the name is more than 50",
			id:       1,
			roomName: strings.Repeat("a", 51),
			mockSetup: func(room *entities.Room) {
				mockRepository.EXPECT().GetById(room.Id).
					Return(&entities.Room{Id: room.Id}, nil)
			},
			expectedError: errors.New("Invalid name"),
		},
		{
			name:     "Failed to update room - Due to the empty room",
			id:       1,
			roomName: "",
			mockSetup: func(room *entities.Room) {
				mockRepository.EXPECT().GetById(room.Id).
					Return(&entities.Room{Id: room.Id}, nil)
			},
			expectedError: errors.New("Invalid name"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			updatedRoom := &entities.Room{
				Id:   tc.id,
				Name: tc.roomName,
			}
			tc.mockSetup(updatedRoom)

			err := service.Update(tc.id, tc.roomName)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestDeleteRoom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockRoomRepository(ctrl)
	service := NewRoomService(mockRepository)

	testCases := []struct {
		name          string
		id            int
		mockSetup     func()
		expectedError error
	}{
		{
			name: "Success to delete room",
			id:   1,
			mockSetup: func() {
				mockRepository.EXPECT().GetById(1).
					Return(&entities.Room{}, nil)
				mockRepository.EXPECT().Delete(1).
					Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Failed to delete room - Due to the room not found",
			id:   999,
			mockSetup: func() {
				mockRepository.EXPECT().GetById(999).
					Return(nil, sql.ErrNoRows)
			},
			expectedError: sql.ErrNoRows,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			err := service.Delete(tc.id)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}
