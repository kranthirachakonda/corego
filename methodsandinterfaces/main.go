package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

func main() {
	products := []*Product{
		{"Kayak", "WaterSports", 275.00},
		{"Lifejacket", "WaterSports", 48.99},
		{"Soccer Ball", "Soccer", 12.99},
	}
	for _, p := range products {
		fmt.Println("Name:", p.name, "Category:", p.category, "Price:", p.price)
	}
}
