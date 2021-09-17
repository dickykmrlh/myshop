package checkout

type Product struct {
	SkuID string
	Price float64
}

type Order struct {
	product            Product
	quantity           int
	discountCalculator Calculator
}

func (o Order) GetPrice() float64 {
	var discount float64
	if o.discountCalculator != nil {
		discount = o.discountCalculator.Calculate(o.product.Price, o.quantity)
	}

	return (o.product.Price * float64(o.quantity)) - discount
}
