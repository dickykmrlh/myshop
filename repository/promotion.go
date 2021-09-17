package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var promotionRepo *promotionManager

type Promotion struct {
	Sku                string `json:"sku"`
	Type               string `json:"type"`
	DiscountPercentage int    `json:"discount_percentage"`
	Rule               Rule   `json:"rule"`
}

type Rule struct {
	MinimumQty int `json:"minimum_qty"`
}

type promotionManager struct {
	promotions []Promotion
}

func NewPromotionRepository() (promotionRepository, error) {
	if promotionRepo != nil {
		return promotionRepo, nil
	}

	fileContents, err := ioutil.ReadFile("data/sample_promotion.json")
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
	return Promotion{}
}

type promotionRepository interface {
	GetPromotion(string) Promotion
}
