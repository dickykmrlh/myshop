package checkout

// Order
type OrderLine struct {
	product            Product
	quantity           int
	discountCalculator Calculator
}

func NewOrder(product Product, discountCalculator Calculator) OrderLine {
	return OrderLine{
		product:            product,
		quantity:           1,
		discountCalculator: discountCalculator,
	}
}

func (o OrderLine) GetPrice() float64 {
	var discount float64
	if o.discountCalculator != nil {
		discount = o.discountCalculator.Calculate(o.product.Price, o.quantity)
	}

	return (o.product.Price * float64(o.quantity)) - discount
}

// Cart
type Cart map[string]OrderLine

func NewCart() Cart {
	cart := make(map[string]OrderLine)
	return cart
}

func (cart Cart) AddOrder(order OrderLine) {

}
