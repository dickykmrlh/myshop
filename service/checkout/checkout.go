package checkout

import (
	"myshop/service/order"
)

type Cart struct {
	Orders []order.Order
}
