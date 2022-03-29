package main

import (
	"fmt"
	"packages/store"
	"packages/store/cart"
)

func main() {
	product := store.NewProduct("Kayak", "WaterSports", 275.99)
	cart := cart.Cart{
		CustomerName: "Kraz",
		Products:     []store.Product{*product},
	}
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", product.Price())
}
