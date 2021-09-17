package discount

import "myshop/service/order"

type PercentageDiscount struct {
	DiscountPercentage float64
	MinimumQuantity    int
	ProductSkuID       string
}

func (p PercentageDiscount) Calculate(orders []order.Order) float64 {
	var totalPrice float64
	var totalQuantity int

	skuIDpass := false
	for _, order := range orders {
		totalPrice += order.GetPrice()
		totalQuantity += order.Quantity

		if order.Product.SkuID == p.ProductSkuID {
			skuIDpass = true
		}

		if totalQuantity < p.MinimumQuantity {
			return 0
		}
	}

	if !skuIDpass {
		return 0
	}

	return totalPrice * (p.DiscountPercentage / 100)
}

type Calculator interface {
	Calculate([]order.Order) float64
}
