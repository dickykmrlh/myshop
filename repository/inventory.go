package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var inventoryRepo *inventoryManager

type inventoryManager struct {
	inventories []Inventory
}

type Inventory struct {
	SkuID string  `json:"sku"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Qty   int     `json:"inventory_qty"`
}

func NewProductRepository() (inventoryRepository, error) {
	if inventoryRepo != nil {
		return inventoryRepo, nil
	}

	fileContents, err := ioutil.ReadFile("data/sample_inventory.json")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	inventoryRepo := &inventoryManager{}
	err = json.Unmarshal(fileContents, &inventoryRepo.inventories)
	if err != nil {
		return nil, err
	}

	return inventoryRepo, nil
}

func (p *inventoryManager) GetByName(productName string) Inventory {
	return Inventory{}
}

type inventoryRepository interface {
	GetByName(string) Inventory
}
