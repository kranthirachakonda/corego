package main

import "math/rand"

func Intrange(min, max int) int {
	return rand.Intn(max-min) + min
}
