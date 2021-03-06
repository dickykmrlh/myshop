package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var promotionRepo *promotionManager

type Promotion struct {
	Sku                string  `json:"sku"`
	Type               string  `json:"type"`
	DiscountPercentage float64 `json:"discount_percentage"`
	Rule               Rule    `json:"rule"`
}

type Rule struct {
	MinimumQty int    `json:"minimum_qty"`
	MustBuy    string `json:"must_buy"`
}

type promotionManager struct {
	promotions []Promotion
}

func NewPromotionRepository(file string) (PromotionRepository, error) {
	if promotionRepo != nil {
		return promotionRepo, nil
	}

	fileContents, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	promotionRepo := &promotionManager{}
	err = json.Unmarshal(fileContents, &promotionRepo.promotions)
	if err != nil {
		return nil, err
	}

	return promotionRepo, nil
}

func (a *promotionManager) GetPromotion(skuID string) Promotion {
	for _, promotion := range a.promotions {
		if promotion.Sku == skuID {
			return promotion
		}
	}
	return Promotion{}
}

type PromotionRepository interface {
	GetPromotion(string) Promotion
}
