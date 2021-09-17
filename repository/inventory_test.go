package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProductRepository(t *testing.T) {
	t.Run("should not return error, and inventories is initialize", func(t *testing.T) {
		inventoryRepo, err := NewProductRepository()
		assert.Nil(t, err)
		assert.NotNil(t, inventoryRepo)
	})
}
