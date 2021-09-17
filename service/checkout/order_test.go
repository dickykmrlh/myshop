package checkout

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrder_getPrice(t *testing.T) {
	type fields struct {
		product            Product
		quantity           int
		discountCalculator Calculator
	}
	tests := []struct {
		name     string
		fields   fields
		expected float64
	}{
		{
			name: "Should return correct price based on quantity bough",
			fields: fields{
				product: Product{
					SkuID: "SK1234",
					Price: 30.00,
				},
				quantity: 3,
			},
			expected: 90.00,
		},
		{
			name: "Should return correct price with discount, when product had discount",
			fields: fields{
				product: Product{
					SkuID: "SK1234",
					Price: 30.00,
				},
				quantity:           3,
				discountCalculator: FreeProductDiscount{MinimumQuantity: 2},
			},
			expected: 60.00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := OrderLine{
				product:            tt.fields.product,
				quantity:           tt.fields.quantity,
				discountCalculator: tt.fields.discountCalculator,
			}
			assert.Equal(t, tt.expected, o.GetPrice(), tt.name)
		})
	}
}
