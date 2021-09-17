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
	return 0
}

type Cart struct {
	Orders             []Order
	DiscountCalculator DiscountCalculator
}

type DiscountCalculator interface {
	Calculate(int, []Order) float64
}
