package main

import (
	"io"
	"strings"
)

func readData(r io.Reader) {
	bSlice := make([]byte, 2)
	for {
		count, err := r.Read(bSlice)
		if count > 0 {
			MyPrint("Read %v bytes: %v", count, string(bSlice[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

func main() {
	MyPrint("Product: %v, Price: %v", Kayak.Name, Kayak.Price)

	r := strings.NewReader("krachakonda")
	readData(r)
}
