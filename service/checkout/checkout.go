package checkout

import "myshop/repository"

type Service struct {
	inventory repository.InventoryRepository
	promotion repository.PromotionRepository
}

func (s Service) Run(productNames string) string {
	return ""
}
