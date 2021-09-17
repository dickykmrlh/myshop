package main

import (
	"myshop/repository"
	"myshop/service/checkout"
)

func main() {
	//checkoutService := setupServer()

}

func setupServer() checkout.Server {
	inventory, err := repository.NewInventoryRepository("../repository/data/sample_inventory.json")
	if err != nil {
		panic(err)
	}

	promotion, err := repository.NewPromotionRepository("../repository/data/sample_promotion.json")
	if err != nil {
		panic(err)
	}

	return checkout.NewCheckoutService(inventory, promotion)
}
