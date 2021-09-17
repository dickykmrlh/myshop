package order

import "myshop/service/product"

type Order struct {
	Product  product.Product
	Quantity int
}

func (o Order) GetPrice() float64 {
	return o.Product.Price * float64(o.Quantity)
}
