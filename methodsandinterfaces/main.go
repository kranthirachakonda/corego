package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

// define Method for printing
func printDetails(product *Product) {
	fmt.Println("Name:", product.name, "Category:", product.category, "Price:", product.price)
}

func main() {
	products := []*Product{
		{"Kayak", "WaterSports", 275.00},
		{"Lifejacket", "WaterSports", 48.99},
		{"Soccer Ball", "Soccer", 12.99},
	}
	for _, p := range products {
		printDetails(p)
	}
}
