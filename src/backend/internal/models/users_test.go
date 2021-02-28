package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	u := &User{Name: "John", Email: "johndoe@gmail.com"}
	t.Run("Create", func(t *testing.T) {
		rows, err := u.Create(DB)
		assert.NoError(t, err, "Error create user in DB")
		assert.Equal(t, rows, 1, "Rows return is not 1")
	})
	t.Run("UpdateName", func(t *testing.T) {
		newName := "Peter"
		assert.NoError(t, u.UpdateName(DB, newName), "Error create user in DB")
	})
	t.Run("Delete", func(t *testing.T) {
		assert.NoError(t, u.Delete(DB), "Error create user in DB")
	})
}
