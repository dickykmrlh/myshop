package checkout

import (
	"myshop/service/discount"
	"myshop/service/order"
)

type Cart struct {
	Orders             []order.Order
	DiscountCalculator discount.Calculator
}
