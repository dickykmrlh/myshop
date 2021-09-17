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
}

func (p PercentageDiscountCalculator) Calculate(orders []Order) float64 {
	var totalPrice float64
	for _, order := range orders {
		totalPrice += order.getPrice()
	}

	return totalPrice * (p.DiscountPercentage / 100)
}

type DiscountCalculator interface {
	Calculate([]Order) float64
}
