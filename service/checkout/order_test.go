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
			assert.Equal(t, tt.expected, o.GetPrice(nil), tt.name)
		})
	}
}

func TestCart_AddOrder(t *testing.T) {
	cart := NewCart()
	t.Run("should add new order, when no order created for that product", func(t *testing.T) {
		cart.AddOrder(OrderLine{
			product:  Product{SkuID: "12345", Price: 100},
			quantity: 1,
		})
		assert.Equal(t, OrderLine{product: Product{SkuID: "12345", Price: 100}, quantity: 1}, cart["12345"])
	})

	t.Run("should added quantity when order of the same product", func(t *testing.T) {
		cart.AddOrder(OrderLine{
			product:  Product{SkuID: "12345", Price: 100},
			quantity: 1,
		})
		assert.Equal(t, OrderLine{product: Product{SkuID: "12345", Price: 100}, quantity: 2}, cart["12345"])
	})

	t.Run("should added new order for another order of different product", func(t *testing.T) {
		cart.AddOrder(OrderLine{
			product:  Product{SkuID: "78901", Price: 50.0},
			quantity: 1,
		})
		assert.Equal(t, OrderLine{product: Product{SkuID: "78901", Price: 50.0}, quantity: 1}, cart["78901"])
	})
}
