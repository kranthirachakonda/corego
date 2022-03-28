package main

import (
	"fmt"
	"packages/store"
)

func main() {
	product := store.NewProduct("Kayak", "WaterSports", 275.99)
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", product.Price())
}
