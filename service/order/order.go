package order

type Order struct {
	product            Product
	quantity           int
	discountCalculator Calculator
}

func (o Order) GetPrice() float64 {
	return o.product.Price * float64(o.quantity)
}
