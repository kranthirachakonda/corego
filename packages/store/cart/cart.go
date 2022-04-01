package cart

import "package/store"

type Cart struct {
	CustomerName string
	Products     []store.Product
}

func (cart *Cart) GetTotal() float64 {
	for _, p := range cart.Products {
		total += p.Price()
	}
	return
}
