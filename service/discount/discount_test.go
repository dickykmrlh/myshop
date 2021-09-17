package discount

import (
	"github.com/stretchr/testify/assert"
	"myshop/service/order"
	"myshop/service/product"
	"testing"
)

func TestPercentageDiscountCalculator_Calculate(t *testing.T) {
	type fields struct {
		DiscountPercentage float64
		MinimumQuantity    int
		ProductSkuID       string
	}
	type args struct {
		orders []order.Order
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
				ProductSkuID:       "SKU1234",
			},
			args: args{
				orders: []order.Order{
					{
						Product: product.Product{
							SkuID: "SKU1234",
							Price: 109.50,
						},
						Quantity: 3,
					},
				},
			},
			expected: 32.85,
		},
		{
			name: "should return 0 discount, when minimum quantity condition didnt pass",
			fields: fields{
				DiscountPercentage: 10,
				MinimumQuantity:    3,
				ProductSkuID:       "SKU1234",
			},
			args: args{
				orders: []order.Order{
					{
						Product: product.Product{
							SkuID: "SKU1234",
							Price: 109.50,
						},
						Quantity: 2,
					},
				},
			},
			expected: 0,
		},
		{
			name: "should return 0 discount, when no discount for the product present",
			fields: fields{
				DiscountPercentage: 10,
				MinimumQuantity:    3,
				ProductSkuID:       "SKU1234",
			},
			args: args{
				orders: []order.Order{
					{
						Product: product.Product{
							SkuID: "SKU1234",
							Price: 109.50,
						},
						Quantity: 2,
					},
					{
						Product: product.Product{
							SkuID: "SKU4567",
							Price: 88.50,
						},
						Quantity: 1,
					},
				},
			},
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PercentageDiscount{
				DiscountPercentage: tt.fields.DiscountPercentage,
				MinimumQuantity:    tt.fields.MinimumQuantity,
				ProductSkuID:       tt.fields.ProductSkuID,
			}
			assert.Equal(t, tt.expected, p.Calculate(tt.args.orders), tt.name)
		})
	}
}
