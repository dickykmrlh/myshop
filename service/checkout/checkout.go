package checkout

import (
	repo "myshop/repository"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var checkoutService *Service

type Service struct {
	inventory repo.InventoryRepository
	promotion repo.PromotionRepository
}

func NewCheckoutService(inventory repo.InventoryRepository, promotion repo.PromotionRepository) Server {
	if checkoutService != nil {
		return checkoutService
	}

	checkoutService := &Service{
		inventory: inventory,
		promotion: promotion,
	}

	return checkoutService
}

func (s Service) Run(productNames []string) string {
	cart := NewCart()
	for _, name := range productNames {
		product := NewProduct(s.inventory.GetByName(strings.TrimSpace(name)))
		discountCalculator := NewDiscount(s.promotion.GetPromotion(product.SkuID))
		orderLine := NewOrder(product, discountCalculator)
		cart.AddOrder(orderLine)
	}

	formatter := message.NewPrinter(language.English)
	return formatter.Sprintf("$%.2f", cart.GetTotalPrice())
}

type Server interface {
	Run([]string) string
}
