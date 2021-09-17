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

func (p PercentageDiscountCalculator) Calculate([]Order) float64 {
	return 0
}

type DiscountCalculator interface {
	Calculate([]Order) float64
}
