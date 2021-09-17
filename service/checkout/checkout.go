package checkout

import (
	"fmt"
	repo "myshop/repository"
)

var checkoutService *CheckoutService

type CheckoutService struct {
	inventory repo.InventoryRepository
	promotion repo.PromotionRepository
}

func NewCheckoutService(inventory repo.InventoryRepository, promotion repo.PromotionRepository) Server {
	if checkoutService != nil {
		return checkoutService
	}

	checkoutService := &CheckoutService{
		inventory: inventory,
		promotion: promotion,
	}

	return checkoutService
}

func (s CheckoutService) Run(productNames []string) string {
	cart := NewCart()
	for _, name := range productNames {
		product := NewProduct(s.inventory.GetByName(name))
		discountCalculator := NewDiscount(s.promotion.GetPromotion(product.SkuID))
		orderLine := NewOrder(product, discountCalculator)
		cart.AddOrder(orderLine)
	}

	return fmt.Sprintf("$%.2f", cart.GetTotalPrice())
}

type Server interface {
	Run([]string) string
}
