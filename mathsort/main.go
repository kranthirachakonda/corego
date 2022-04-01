package main

import (
	"math/rand"
	"time"
)

func main() {
	MyPrint("Math & Sorting")
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 5; i++ {
		MyPrint("%v: %v", i, Intrange(5, 12))
	}
}
