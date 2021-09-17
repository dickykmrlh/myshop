package checkout

type Product struct {
	SkuID string
	Price float64
}

type Order struct {
	Product  Product
	Quantity int
}

func (o Order) getPrice() float64 {
	return o.Product.Price * float64(o.Quantity)
}

type Cart struct {
	Orders             []Order
	DiscountCalculator DiscountCalculator
}

type PercentageDiscountCalculator struct {
	DiscountPercentage float64
	MinimumQuantity    int
	ProductSkuID       string
}

func (p PercentageDiscountCalculator) Calculate(orders []Order) float64 {
	var totalPrice float64
	var totalQuantity int

	skuIDpass := false
	for _, order := range orders {
		totalPrice += order.getPrice()
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

type DiscountCalculator interface {
	Calculate([]Order) float64
}
