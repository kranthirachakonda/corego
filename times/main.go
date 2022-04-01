package main

import (
	"fmt"
	"time"
)

func PrintTime(label string, t *time.Time) {
	layout := "Day: 01 Month: Apr Year: 2006"
	fmt.Println(label, t.Format(layout))
}

func main() {
	curr := time.Now()
	spec := time.Date(2022, time.April, 1, 17, 42, 10, 11, time.Local)
	unx := time.Unix(1648849456, 0)

	PrintTime("Current:", &curr)
	PrintTime("Defined:", &spec)
	PrintTime("Unix time:", &unx)
}
