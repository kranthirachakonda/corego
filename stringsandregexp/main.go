package main

import (
	"fmt"
	"strings"
)

func main() {
	desc := "This is a sample string for tests"
	splits := strings.Split(desc, " ")

	for _, x := range splits {
		fmt.Println("Split >>" + x + "<<")

	}
	later := strings.SplitAfter(desc, " ")
	for _, x := range later {
		fmt.Println("After ==" + x + "==")
	}
}
