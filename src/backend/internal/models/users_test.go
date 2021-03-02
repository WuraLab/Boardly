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
	t.Run("UpdateName", func(t *testing.T) {
		u.FirstName = "Peter"
		assert.NoError(t, u.UpdateName(DB), "Error updating name in DB")
	})
	t.Run("UpdatePassword", func(t *testing.T) {
		u.Password = "newpassword"
		assert.NoError(t, u.UpdatePassword(DB), "Error updating password in DB")
	})
	t.Run("UpdateEmail", func(t *testing.T) {
		u.Email = "newemail@gmail.com"
		assert.NoError(t, u.UpdateEmail(DB), "Error updating email in DB")
	})
	t.Run("Save", func(t *testing.T) {
		assert.NoError(t, u.Save(DB), "Error saving User in DB")
	})
	t.Run("Delete", func(t *testing.T) {
		assert.NoError(t, u.Delete(DB), "Error deleting user from DB")
	})
}
