package checkout

type PercentageDiscount struct {
	DiscountPercentage float64
	MinimumQuantity    int
}

func (p PercentageDiscount) Calculate(productPrice float64, quantity int) float64 {
	if quantity < p.MinimumQuantity {
		return 0
	}

	return (productPrice * float64(quantity)) * (p.DiscountPercentage / 100)
}

type FreeProductDiscount struct {
	MinimumQuantity int
}

func (f FreeProductDiscount) Calculate(productPrice float64, quantity int) float64 {
	if quantity < f.MinimumQuantity {
		return 0
	}
	return productPrice
}

type Calculator interface {
	Calculate(float64, int) float64
}
