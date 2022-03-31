package main

import "fmt"

func main() {
	categories := []string{"Watersports", "Chess", "Running"}
	channel := make(chan ChanMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for msg := range channel {
		if msg.CategoryError == nil {
			fmt.Println(msg.Category, "total:", ToCurrency(msg.Total))
		} else {
			fmt.Println(msg.Category, "(no such category)")
		}
	}
	/* 	for _, ctg := range categories {
		total, err := Products.TotalPrice(ctg)
		if err == nil {
			fmt.Println("Category:", ctg, "Total:", ToCurrency(total))
		} else {
			fmt.Println("Category:", ctg, "(does'nt exist)")
		}

	} */
}
