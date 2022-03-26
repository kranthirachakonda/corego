package main

import "fmt"

func main() {
	type Product struct {
		name     string
		category string
		price    float64
	}
	kayak := Product{
		name:     "Kayak",
		category: "WaterSports",
		price:    275.99,
	}
	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 299
	fmt.Println("Updated price:", kayak.price)
}
