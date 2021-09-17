package checkout

import repo "myshop/repository"

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
	/*
		for _, name := range productNames {
			item := s.inventory.GetByName(name)
			product := NewProduct(item)
			promotion := s.promotion.GetPromotion(product.SkuID)
			orderLine := NewOrder(product)
		}
	*/
	return ""
}

type Server interface {
	Run([]string) string
}
