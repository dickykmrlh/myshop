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

	type args struct {
		productNames []string
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "should checkout with correct total price return",
			args: args{
				productNames: []string{"Alexa Speaker", "Alexa Speaker", "Alexa Speaker"},
			},
			expected: "$295.65",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := s.Run(tt.args.productNames)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
