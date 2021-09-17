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

type DiscountCalculator interface {
	Calculate(int, []Order) float64
}
