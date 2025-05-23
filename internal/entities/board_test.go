package entities

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateBoard(t *testing.T) {
	testCases := []struct {
		name          string
		board         *Board
		expectedError error
	}{
		{
			name: "Success to validate",
			board: &Board{
				Name:     "test board",
				Priority: 0,
			},
			expectedError: nil,
		},
		{
			name: "Failed to validate - Due to the name is empty",
			board: &Board{
				Name:     "",
				Priority: 0,
			},
			expectedError: errors.New("Invalid name"),
		},
		{
			name: "Failed to validate - Due to the name is larger than 50 characters",
			board: &Board{
				Name:     strings.Repeat("a", 51),
				Priority: 0,
			},
			expectedError: errors.New("Invalid name"),
		},
		{
			name: "Failed to validate - Due to the priority is negative number",
			board: &Board{
				Name:     "test board",
				Priority: -1,
			},
			expectedError: errors.New("Invalid priority size"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.board.Validate()

			assert.Equal(t, tc.expectedError, err)
		})
	}
}
