package checkout

import "myshop/repository"

type Product struct {
	SkuID string
	Price float64
}

func NewProduct(item repository.Inventory) Product {
	return Product{
		SkuID: item.SkuID,
		Price: item.Price,
	}
}
