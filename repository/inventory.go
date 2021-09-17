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

func NewInventoryRepository(file string) (InventoryRepository, error) {
	if inventoryRepo != nil {
		return inventoryRepo, nil
	}

	fileContents, err := ioutil.ReadFile(file)
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
	for _, inventory := range p.inventories {
		if inventory.Name == productName {
			return inventory
		}
	}
	return Inventory{}
}

func (p *inventoryManager) UpdateQty(SkuID string, qty int) {
	for i := 0; i < len(p.inventories); i++ {
		if p.inventories[i].SkuID == SkuID {
			p.inventories[i].Qty -= qty
		}
	}
}

type InventoryRepository interface {
	GetByName(string) Inventory
	UpdateQty(string, int)
}
