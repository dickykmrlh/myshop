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

func (o OrderLine) Empty() bool {
	return o.product.SkuID == ""
}

// Cart
type Cart map[string]OrderLine

func NewCart() Cart {
	cart := make(map[string]OrderLine)
	return cart
}

func (cart Cart) AddOrder(newOrder OrderLine) {
	currentOrder := cart[newOrder.product.SkuID]
	if currentOrder.Empty() {
		cart[newOrder.product.SkuID] = newOrder
	}

	currentOrder.quantity += 1
	cart[currentOrder.product.SkuID] = currentOrder
}

func (cart Cart) GetTotalPrice() float64 {
	var totalPrice float64
	for _, orderLine := range cart {
		totalPrice += orderLine.GetPrice()
	}

	return totalPrice
}
