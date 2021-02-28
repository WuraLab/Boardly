package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wuraLab/boardly/src/backend/internal/models"
)

func TestMigrate(t *testing.T) {
	assert.NoError(t, Migrate(models.DB), "Migration was not successful")
}
