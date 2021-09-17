package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrder_getPrice(t *testing.T) {
	type fields struct {
		Product  Product
		Quantity int
	}
	tests := []struct {
		name     string
		fields   fields
		expected float64
	}{
		{
			name: "Should return correct price based on quantity bough",
			fields: fields{
				Product: Product{
					SkuID: "SK1234",
					Price: 30.00,
				},
				Quantity: 3,
			},
			expected: 90.00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Order{
				Product:  tt.fields.Product,
				Quantity: tt.fields.Quantity,
			}
			assert.Equal(t, tt.expected, o.getPrice(), tt.name)
		})
	}
}

func TestPercentageDiscountCalculator_Calculate(t *testing.T) {
	type fields struct {
		DiscountPercentage float64
		MinimumQuantity    int
	}
	type args struct {
		orders []Order
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected float64
	}{
		{
			name: "should return correct discount, when validation rule pass",
			fields: fields{
				DiscountPercentage: 10,
				MinimumQuantity:    3,
			},
			args: args{
				orders: []Order{
					{
						Product: Product{
							SkuID: "SKU1234",
							Price: 109.50,
						},
						Quantity: 3,
					},
				},
			},
			expected: 32.85,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PercentageDiscountCalculator{
				DiscountPercentage: tt.fields.DiscountPercentage,
			}
			assert.Equal(t, tt.expected, p.Calculate(tt.args.orders), tt.name)
		})
	}
}
