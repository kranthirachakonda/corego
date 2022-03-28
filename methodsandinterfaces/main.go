package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

type Supplier struct {
	name, city string
}

// newProduct definition
func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

// printDetails method
// (product *Product) receiver
func (product *Product) printDetails() {
	fmt.Println("Name:", product.name, ", Category:", product.category, ", Price:", product.calcTax(0.5, 10))
}

func (supplier *Supplier) printDetails() {
	fmt.Println("Name:", supplier.name, ", City:", supplier.city)
}

// calcTax method
// (product *Product) receiver; (rate, threshold float64) Parameters of method with return Type
func (product *Product) calcTax(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + (product.price * rate)
	}
	return product.price
}

func main() {
	products := []*Product{
		{"Kayak", "WaterSports", 275.00},
		{"Lifejacket", "WaterSports", 48.99},
		{"Soccer Ball", "Soccer", 12.99},
	}
	for _, p := range products {
		p.printDetails() //Invoke method
	}

	suppliers := []*Supplier{
		{"KR", "Charlotte"},
		{"RR", "SanAntonio"},
		{"RVG", "Plano"},
	}
	for _, s := range suppliers {
		s.printDetails() //Invoke method
	}

}
