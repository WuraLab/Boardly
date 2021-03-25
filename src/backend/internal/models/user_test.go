package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wuraLab/boardly/src/backend/internal/models"
)

func TestUser(t *testing.T) {
	u := &models.User{FirstName: "Test", LastName: "User", Email: "testuser@example.com", Password: "test"}
	t.Run("Create", func(t *testing.T) {
		rows, err := u.Create(DB)
		assert.NoError(t, err, "Error create user in DB")
		assert.Equal(t, rows, 1, "Rows return is not 1")
	})
}
