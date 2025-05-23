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

func TestCreateBoard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockBoardRepository(ctrl)
	service := NewBoardService(mockRepository)

	testCases := []struct {
		name          string
		boardName     string
		priority      int
		roomId        int
		mockSetup     func(board *entities.Board)
		expectedError error
	}{
		{
			name:      "Success to create board",
			boardName: "test board",
			priority:  0,
			roomId:    1,
			mockSetup: func(board *entities.Board) {
				mockRepository.EXPECT().Create(board).
					Return(nil)
			},
			expectedError: nil,
		},
		{
			name:          "Failed to create board - Due to number of characters in the name is more than 50",
			boardName:     strings.Repeat("a", 51),
			priority:      0,
			roomId:        1,
			mockSetup:     func(board *entities.Board) {},
			expectedError: errors.New("Invalid name"),
		},
		{
			name:          "Failed to create board - Due to the empty name",
			boardName:     "",
			priority:      0,
			roomId:        1,
			mockSetup:     func(board *entities.Board) {},
			expectedError: errors.New("Invalid name"),
		},
		{
			name:          "Failed to create board - Due to the priority is negative number",
			boardName:     "test board",
			priority:      -1,
			roomId:        1,
			mockSetup:     func(board *entities.Board) {},
			expectedError: errors.New("Invalid priority size"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			board := &entities.Board{
				Name:     tc.boardName,
				Priority: tc.priority,
				RoomId:   tc.roomId,
			}
			tc.mockSetup(board)

			err := service.Create(tc.boardName, tc.priority, tc.roomId)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestUpdateBoard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockBoardRepository(ctrl)
	service := NewBoardService(mockRepository)

	testCases := []struct {
		name          string
		id            int
		boardName     string
		priority      int
		mockSetup     func(board *entities.Board)
		expectedError error
	}{
		{
			name:      "Success to update board",
			id:        1,
			boardName: "test board",
			priority:  0,
			mockSetup: func(board *entities.Board) {
				mockRepository.EXPECT().GetById(board.Id).
					Return(&entities.Board{Id: board.Id}, nil)
				mockRepository.EXPECT().Update(board).
					Return(nil)
			},
			expectedError: nil,
		},
		{
			name:      "Failed to update board - Due to the board not found",
			id:        999,
			boardName: "test board",
			priority:  0,
			mockSetup: func(board *entities.Board) {
				mockRepository.EXPECT().GetById(board.Id).
					Return(nil, sql.ErrNoRows)
			},
			expectedError: sql.ErrNoRows,
		},
		{
			name:      "Failed to update board - Due to number of characters in the name is more than 50",
			id:        1,
			boardName: strings.Repeat("a", 51),
			priority:  0,
			mockSetup: func(board *entities.Board) {
				mockRepository.EXPECT().GetById(board.Id).
					Return(&entities.Board{Id: board.Id}, nil)
			},
			expectedError: errors.New("Invalid name"),
		},
		{
			name:      "Failed to update board - Due to the empty name",
			id:        1,
			boardName: "",
			priority:  0,
			mockSetup: func(board *entities.Board) {
				mockRepository.EXPECT().GetById(board.Id).
					Return(&entities.Board{Id: board.Id}, nil)
			},
			expectedError: errors.New("Invalid name"),
		},
		{
			name:      "Failed to update board - Due to the priority is negative number",
			id:        1,
			boardName: "test board",
			priority:  -1,
			mockSetup: func(board *entities.Board) {
				mockRepository.EXPECT().GetById(board.Id).
					Return(&entities.Board{Id: board.Id}, nil)
			},
			expectedError: errors.New("Invalid priority size"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			updatedBoard := &entities.Board{
				Id:       tc.id,
				Name:     tc.boardName,
				Priority: tc.priority,
			}
			tc.mockSetup(updatedBoard)

			err := service.Update(tc.id, tc.boardName, tc.priority)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestDeleteBoard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockBoardRepository(ctrl)
	service := NewBoardService(mockRepository)

	testCases := []struct {
		name          string
		id            int
		mockSetup     func()
		expectedError error
	}{
		{
			name: "Success to delete board",
			id:   1,
			mockSetup: func() {
				mockRepository.EXPECT().GetById(1).
					Return(&entities.Board{}, nil)
				mockRepository.EXPECT().Delete(1).
					Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Failed to delete board - Due to the board not found",
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
