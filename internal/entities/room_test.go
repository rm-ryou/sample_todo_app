package entities

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateRoom(t *testing.T) {
	testCases := []struct {
		name          string
		room          *Room
		expectedError error
	}{
		{
			name:          "Success to validate",
			room:          &Room{Name: "test room"},
			expectedError: nil,
		},
		{
			name:          "Failed to validate - Due to the name is empty",
			room:          &Room{Name: ""},
			expectedError: errors.New("Invalid name"),
		},
		{
			name:          "Failed to validate - Due to the name is larger than 50 characters",
			room:          &Room{Name: strings.Repeat("a", 51)},
			expectedError: errors.New("Invalid name"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.room.Validate()

			assert.Equal(t, tc.expectedError, err)
		})
	}
}
