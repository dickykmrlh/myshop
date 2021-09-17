package checkout

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPercentageDiscountCalculator_Calculate(t *testing.T) {
	type fields struct {
		DiscountPercentage float64
		MinimumQuantity    int
	}
	type args struct {
		price    float64
		quantity int
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
				price:    109.50,
				quantity: 3,
			},
			expected: 32.85,
		},
		{
			name: "should return 0 discount, when minimum quantity condition didnt pass",
			fields: fields{
				DiscountPercentage: 10,
				MinimumQuantity:    3,
			},
			args: args{
				price:    109.50,
				quantity: 2,
			},
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PercentageDiscount{
				DiscountPercentage: tt.fields.DiscountPercentage,
				MinimumQuantity:    tt.fields.MinimumQuantity,
			}
			assert.Equal(t, tt.expected, p.Calculate(tt.args.price, tt.args.quantity), tt.name)
		})
	}
}

func TestFreeProductDiscount_Calculate(t *testing.T) {
	type fields struct {
		MinimumQuantity int
	}
	type args struct {
		price    float64
		quantity int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected float64
	}{
		{
			name: "should return discount amount equal to product price, when minimum quantity pass",
			fields: fields{
				MinimumQuantity: 2,
			},
			args: args{
				price:    109.50,
				quantity: 2,
			},
			expected: 109.50,
		},
		{
			name: "should return discount amount 0, when minimum quantity didnt pass",
			fields: fields{
				MinimumQuantity: 3,
			},
			args: args{
				price:    109.50,
				quantity: 2,
			},
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FreeProductDiscount{
				MinimumQuantity: tt.fields.MinimumQuantity,
			}
			assert.Equal(t, tt.expected, f.Calculate(tt.args.price, tt.args.quantity), tt.name)
		})
	}
}
