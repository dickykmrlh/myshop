package checkout

import (
	"myshop/repository"
	"myshop/utils"
)

func NewDiscount(promotion repository.Promotion) Calculator {
	switch promotion.Type {
	case "free":
		return FreeProductDiscount{
			MustBuy:         promotion.Rule.MustBuy,
			MinimumQuantity: promotion.Rule.MinimumQty,
		}
	case "percentage":
		return PercentageDiscount{
			DiscountPercentage: promotion.DiscountPercentage,
			MinimumQuantity:    promotion.Rule.MinimumQty,
		}
	}

	return nil
}

type PercentageDiscount struct {
	DiscountPercentage float64
	MinimumQuantity    int
}

func (p PercentageDiscount) Calculate(productPrice float64, quantity int, _ []string) float64 {
	if quantity < p.MinimumQuantity {
		return 0
	}

	return (productPrice * float64(quantity)) * (p.DiscountPercentage / 100)
}

type FreeProductDiscount struct {
	MustBuy         string
	MinimumQuantity int
}

func (f FreeProductDiscount) Calculate(productPrice float64, quantity int, itemsBought []string) float64 {
	if f.MustBuy != "" {
		if !utils.StringArrContains(itemsBought, f.MustBuy) {
			return 0
		}
	}

	if quantity < f.MinimumQuantity {
		return 0
	}
	return productPrice
}

type Calculator interface {
	Calculate(float64, int, []string) float64
}
