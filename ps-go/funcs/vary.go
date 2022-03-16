package main

import "fmt"

func main() {
	bestFinish := champFinishes(12, 3, 5, 6, 1, 7, 2, 2, 11)
	fmt.Println("Best rank: ", bestFinish)
}

// ...int: take any number of integer inputs
func champFinishes(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best

}
