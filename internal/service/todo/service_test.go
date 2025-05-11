package todo

import (
	"database/sql"
	"testing"

	"github.com/rm-ryou/sample_todo_app/internal/entity"
	mock_repository "github.com/rm-ryou/sample_todo_app/internal/service/todo/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockRepository(ctrl)
	service := NewService(mockRepository)

	testCases := []struct {
		name          string
		todo          *entity.Todo
		mockSetup     func(todo *entity.Todo)
		expectedError error
	}{
		{
			name: "Success to update todo",
			todo: &entity.Todo{
				Id:       1,
				Title:    "Test title",
				Done:     false,
				Priority: 0,
				DueDate:  nil,
			},
			mockSetup: func(todo *entity.Todo) {
				mockRepository.EXPECT().GetById(todo.Id).
					Return(&entity.Todo{}, nil)
				mockRepository.EXPECT().Update(todo).
					Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Failed to update todo - Due to the todo not found",
			todo: &entity.Todo{
				Id:       999,
				Title:    "Test title",
				Done:     false,
				Priority: 0,
				DueDate:  nil,
			},
			mockSetup: func(todo *entity.Todo) {
				mockRepository.EXPECT().GetById(todo.Id).
					Return(nil, sql.ErrNoRows)
			},
			expectedError: sql.ErrNoRows,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup(tc.todo)

			err := service.UpdateTodo(tc.todo)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestDeleteTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockRepository(ctrl)
	service := NewService(mockRepository)

	testCases := []struct {
		name          string
		id            int
		mockSetup     func()
		expectedError error
	}{
		{
			name: "Success to update todo",
			id:   1,
			mockSetup: func() {
				mockRepository.EXPECT().GetById(1).
					Return(&entity.Todo{}, nil)
				mockRepository.EXPECT().Delete(1).
					Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Failed to delete todo - Due to the todo not found",
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

			err := service.DeleteTodo(tc.id)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}
