package main

import "fmt"

func main() {
	categories := []string{"Watersports", "Chess"}
	for _, ctg := range categories {
		total := Products.TotalPrice(ctg)
		fmt.Println("Category:", ctg, "Total:", ToCurrency(total))
	}
}
