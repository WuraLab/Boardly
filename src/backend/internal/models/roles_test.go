package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wuraLab/boardly/src/backend/internal/models"
)

func TestRole(t *testing.T) {
	r := &models.Role{Name: "Employee"}
	t.Run("Create", func(t *testing.T) {
		rows, err := r.Create(DB)
		assert.NoError(t, err, "Error creating role in DB")
		assert.Equal(t, rows, 1, "Rows return is not 1")
	})
	t.Run("Save", func(t *testing.T) {
		rows, err := r.Save(DB)
		assert.NoError(t, err, "Error saving role in DB")
		assert.Equal(t, rows, 1, "Rows return is not 1")
	})
	t.Run("Delete", func(t *testing.T) {
		assert.NoError(t, r.Delete(DB), "Error deleting user from DB")
	})
}
