package repository

import (
	"testing"

	"github.com/rm-ryou/sample_todo_app/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestSetupConnection(t *testing.T) {
	testCases := []struct {
		name        string
		config      config.DB
		expectError bool
	}{
		{
			name: "Success to connection",
			config: config.DB{
				Database: MYSQL_DATABASE,
				User:     MYSQL_USER,
				Password: MYSQL_PASSWORD,
				Host:     MYSQL_HOST,
				Port:     MYSQL_PORT,
			},
			expectError: false,
		},
		{
			name: "Failed to connection with wrong port",
			config: config.DB{
				Database: MYSQL_DATABASE,
				User:     MYSQL_USER,
				Password: MYSQL_PASSWORD,
				Host:     MYSQL_HOST,
				Port:     "invalidPort",
			},
			expectError: true,
		},
		{
			name: "Failed to connection with wrong usre info",
			config: config.DB{
				Database: MYSQL_DATABASE,
				User:     "invalidUser",
				Password: "invalidPassword",
				Host:     MYSQL_HOST,
				Port:     MYSQL_PORT,
			},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, err := SetupConnection(tc.config)

			if !tc.expectError {
				assert.NoError(t, err)
				assert.NotNil(t, db)
				db.Close()
			} else {
				assert.Error(t, err)
				assert.Nil(t, db)
			}
		})
	}
}
