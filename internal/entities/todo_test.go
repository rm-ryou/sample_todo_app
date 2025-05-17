package entities

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name          string
		todo          *Todo
		expectedError error
	}{
		{
			name: "Success to validate",
			todo: &Todo{
				Title:    "valid",
				Done:     false,
				Priority: 1,
				DueDate:  &now,
			},
			expectedError: nil,
		},
		{
			name: "Failed to validate - Due to the title is empty",
			todo: &Todo{
				Title:    "",
				Done:     false,
				Priority: 1,
				DueDate:  &now,
			},
			expectedError: errors.New("Title is required"),
		},
		{
			name: "Failed to validate - Due to the title is larger than 50 characters",
			todo: &Todo{
				Title:    strings.Repeat("a", 51),
				Done:     false,
				Priority: 1,
				DueDate:  &now,
			},
			expectedError: errors.New("Title must be 50 characters or less"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.todo.Validate()

			assert.Equal(t, tc.expectedError, err)
		})
	}
}
