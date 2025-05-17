package services

import (
	"database/sql"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/rm-ryou/sample_todo_app/internal/entities"
	mock_repository "github.com/rm-ryou/sample_todo_app/internal/interfaces/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockTodoRepository(ctrl)
	service := NewTodoService(mockRepository)

	testCases := []struct {
		name          string
		title         string
		done          bool
		priority      int
		dueDate       *time.Time
		mockSetup     func(todo *entities.Todo)
		expectedError error
	}{
		{
			name:     "Success to create todo",
			title:    "Test title",
			done:     false,
			priority: 0,
			dueDate:  nil,
			mockSetup: func(todo *entities.Todo) {
				mockRepository.EXPECT().Create(todo).
					Return(nil)
			},
			expectedError: nil,
		},
		{
			name:          "Failed to create todo - Due to number of characters in the title is more than 50",
			title:         strings.Repeat("a", 51),
			done:          false,
			priority:      0,
			dueDate:       nil,
			mockSetup:     func(todo *entities.Todo) {},
			expectedError: errors.New("Invalid title"),
		},
		{
			name:          "Failed to create todo - Due to the empty title",
			title:         "",
			done:          false,
			priority:      0,
			dueDate:       nil,
			mockSetup:     func(todo *entities.Todo) {},
			expectedError: errors.New("Invalid title"),
		},
		{
			name:          "Failed to create todo - Due to the priority is negative number",
			title:         "Test title",
			done:          false,
			priority:      -1,
			dueDate:       nil,
			mockSetup:     func(todo *entities.Todo) {},
			expectedError: errors.New("Invalid priority size"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			todo := &entities.Todo{
				Title:    tc.title,
				Done:     tc.done,
				Priority: tc.priority,
				DueDate:  tc.dueDate,
			}
			tc.mockSetup(todo)

			err := service.Create(tc.title, tc.done, tc.priority, tc.dueDate)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestUpdateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockTodoRepository(ctrl)
	service := NewTodoService(mockRepository)

	testCases := []struct {
		name          string
		id            int
		title         string
		done          bool
		priority      int
		dueDate       *time.Time
		mockSetup     func(todo *entities.Todo)
		expectedError error
	}{
		{
			name:     "Success to update todo",
			id:       1,
			title:    "Test title",
			done:     false,
			priority: 0,
			dueDate:  nil,
			mockSetup: func(todo *entities.Todo) {
				mockRepository.EXPECT().GetById(todo.Id).
					Return(&entities.Todo{Id: todo.Id}, nil)
				mockRepository.EXPECT().Update(todo).
					Return(nil)
			},
			expectedError: nil,
		},
		{
			name:     "Failed to update todo - Due to the todo not found",
			id:       999,
			title:    "Test title",
			done:     false,
			priority: 0,
			dueDate:  nil,
			mockSetup: func(todo *entities.Todo) {
				mockRepository.EXPECT().GetById(todo.Id).
					Return(nil, sql.ErrNoRows)
			},
			expectedError: sql.ErrNoRows,
		},
		{
			name:     "Failed to update todo - Due to number of characters in the title is more than 50",
			id:       1,
			title:    strings.Repeat("a", 51),
			done:     false,
			priority: 0,
			dueDate:  nil,
			mockSetup: func(todo *entities.Todo) {
				mockRepository.EXPECT().GetById(todo.Id).
					Return(&entities.Todo{Id: todo.Id}, nil)
			},
			expectedError: errors.New("Invalid title"),
		},
		{
			name:     "Failed to update todo - Due to the empty title",
			id:       1,
			title:    "",
			done:     false,
			priority: 0,
			dueDate:  nil,
			mockSetup: func(todo *entities.Todo) {
				mockRepository.EXPECT().GetById(todo.Id).
					Return(&entities.Todo{Id: todo.Id}, nil)
			},
			expectedError: errors.New("Invalid title"),
		},
		{
			name:     "Failed to update todo - Due to the priority is negative number",
			id:       1,
			title:    "Test title",
			done:     false,
			priority: -1,
			dueDate:  nil,
			mockSetup: func(todo *entities.Todo) {
				mockRepository.EXPECT().GetById(todo.Id).
					Return(&entities.Todo{Id: todo.Id}, nil)
			},
			expectedError: errors.New("Invalid priority size"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			updatedTodo := &entities.Todo{
				Id:       tc.id,
				Title:    tc.title,
				Done:     tc.done,
				Priority: tc.priority,
				DueDate:  tc.dueDate,
			}
			tc.mockSetup(updatedTodo)

			err := service.Update(tc.id, tc.title, tc.done, tc.priority, tc.dueDate)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestDeleteTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockTodoRepository(ctrl)
	service := NewTodoService(mockRepository)

	testCases := []struct {
		name          string
		id            int
		mockSetup     func()
		expectedError error
	}{
		{
			name: "Success to delete todo",
			id:   1,
			mockSetup: func() {
				mockRepository.EXPECT().GetById(1).
					Return(&entities.Todo{}, nil)
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

			err := service.Delete(tc.id)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}
