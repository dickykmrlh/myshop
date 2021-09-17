package order

import (
	"github.com/stretchr/testify/assert"
	"myshop/service/product"
	"testing"
)

func TestOrder_getPrice(t *testing.T) {
	type fields struct {
		Product  product.Product
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
				Product: product.Product{
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
			assert.Equal(t, tt.expected, o.GetPrice(), tt.name)
		})
	}
}
