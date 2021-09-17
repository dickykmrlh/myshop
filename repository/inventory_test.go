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
	})

	t.Run("should return empty inventory, when product doesnt exist", func(t *testing.T) {
		actual := inventoryManager.GetByName("samsung")
		assert.Empty(t, actual)
	})
}

func Test_inventoryManager_UpdateQty(t *testing.T) {
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

	t.Run("should update qty of product correctly", func(t *testing.T) {
		inventoryManager.UpdateQty("12345", 2)
		actual := inventoryManager.GetByName("iPhone")
		assert.Equal(t, 8, actual.Qty)
	})

	t.Run("should not update qty if no product found", func(t *testing.T) {
		inventoryManager.UpdateQty("xxxxxx", 2)
		assert.Equal(t, 10, inventoryManager.inventories[0].Qty)
		assert.Equal(t, 3, inventoryManager.inventories[1].Qty)
	})
}
