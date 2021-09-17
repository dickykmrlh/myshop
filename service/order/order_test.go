package order

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrder_getPrice(t *testing.T) {
	type fields struct {
		product  Product
		quantity int
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Order{
				product:  tt.fields.product,
				quantity: tt.fields.quantity,
			}
			assert.Equal(t, tt.expected, o.GetPrice(), tt.name)
		})
	}
}
