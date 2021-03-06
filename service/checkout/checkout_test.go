package checkout

import (
	"myshop/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_Run(t *testing.T) {
	inventory, _ := repository.NewInventoryRepository("../../repository/data/sample_inventory.json")
	promotion, _ := repository.NewPromotionRepository("../../repository/data/sample_promotion.json")

	s := NewCheckoutService(inventory, promotion)

	t.Run("should checkout with correct total with price and percentage", func(t *testing.T) {
		actual := s.Run([]string{"Alexa Speaker", "Alexa Speaker", "Alexa Speaker"})
		assert.Equal(t, "Total: $295.65", actual)
		assert.Equal(t, 7, inventory.GetByName("Alexa Speaker").Qty, "should update the inventory")
	})

	t.Run("should checkout with correct total with price and free product discount", func(t *testing.T) {
		actual := s.Run([]string{"Google Home", "Google Home", "Google Home"})
		assert.Equal(t, "Total: $99.98", actual)
		assert.Equal(t, 7, inventory.GetByName("Google Home").Qty, "should update the inventory")
	})

	t.Run("should checkout with correct total with price and partially discount applied", func(t *testing.T) {
		actual := s.Run([]string{"Google Home", "Google Home", "Alexa Speaker", "Alexa Speaker", "Alexa Speaker"})
		assert.Equal(t, "Total: $395.63", actual)
	})

	t.Run("should checkout with correct price with no discount applied", func(t *testing.T) {
		actual := s.Run([]string{"Google Home", "Alexa Speaker", "Alexa Speaker"})
		assert.Equal(t, "Total: $268.99", actual)
	})

	t.Run("should checkout with correct price and free item", func(t *testing.T) {
		actual := s.Run([]string{"MacBook Pro", "Raspberry Pi B"})
		assert.Equal(t, "Total: $5,399.99", actual)
	})

	t.Run("should checkout with correct price and free item for once item only", func(t *testing.T) {
		actual := s.Run([]string{"Raspberry Pi B", "MacBook Pro", "Raspberry Pi B"})
		assert.Equal(t, "Total: $5,429.99", actual)
	})

	t.Run("should checkout empty price if given 0 item", func(t *testing.T) {
		actual := s.Run([]string{})
		assert.Equal(t, "Total: $0.00", actual)
	})

}
