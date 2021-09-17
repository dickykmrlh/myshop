package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInventoryRepository(t *testing.T) {
	t.Run("should not return error, and inventories is initialize", func(t *testing.T) {
		inventoryRepo, err := NewInventoryRepository("data/sample_inventory.json")
		assert.Nil(t, err)
		assert.NotNil(t, inventoryRepo)
	})
}

func Test_inventoryManager_GetByName(t *testing.T) {
	inventoryManager := inventoryManager{
		inventories: []Inventory{
			{
				SkuID: "12345",
				Name:  "iPhone",
				Price: 100.0,
				Qty:   10,
			},
			{
				SkuID: "67890",
				Name:  "bose700",
				Price: 30.0,
				Qty:   3,
			},
		},
	}

	t.Run("should return correct inventory, when product exist", func(t *testing.T) {
		actual := inventoryManager.GetByName("bose700")
		assert.Equal(t, Inventory{SkuID: "67890", Name: "bose700", Price: 30.0, Qty: 3}, actual)
		assert.Equal(t, 2, inventoryManager.inventories[1].Qty)
	})

	t.Run("should return empty inventory, when product doesnt exist", func(t *testing.T) {
		actual := inventoryManager.GetByName("samsung")
		assert.Empty(t, actual)
	})
}
