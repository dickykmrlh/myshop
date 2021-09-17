package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPromotionRepository(t *testing.T) {
	t.Run("should not return error, and promotion repository is initialize", func(t *testing.T) {
		promotionRepo, err := NewPromotionRepository("data/sample_promotion.json")
		assert.Nil(t, err)
		assert.NotNil(t, promotionRepo)
	})
}

func Test_promotionManager_GetPromotion(t *testing.T) {
	promotionManager := promotionManager{
		promotions: []Promotion{
			{
				Sku:  "12345",
				Type: "free",
				Rule: Rule{
					MinimumQty: 2,
				},
			},
			{
				Sku:  "67890",
				Type: "percentage",
				Rule: Rule{
					MinimumQty: 4,
				},
			},
		},
	}

	t.Run("should return correct inventory, when product exist", func(t *testing.T) {
		actual := promotionManager.GetPromotion("12345")
		assert.Equal(t, Promotion{Sku: "12345", Type: "free", Rule: Rule{MinimumQty: 2}}, actual)
	})

	t.Run("should return empty inventory, when product doesnt exist", func(t *testing.T) {
		actual := promotionManager.GetPromotion("xxxxxx")
		assert.Empty(t, actual)
	})
}
