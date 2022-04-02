package main

import (
	"io"
	"strings"
)

func processData(r io.Reader, w io.Writer) {
	bSlice := make([]byte, 2)
	for {
		count, err := r.Read(bSlice)
		if count > 0 {
			w.Write(bSlice[0:count])
			MyPrint("Read %v bytes: %v", count, string(bSlice[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

func copyData(r io.Reader, w io.Writer) {
	count, err := io.Copy(w, r)
	if err == nil {
		MyPrint("Copied bytes: %v", count)
	} else {
		MyPrint("Err copying %v", err.Error())
	}
}

func main() {
	MyPrint("Product: %v, Price: %v", Kayak.Name, Kayak.Price)

	r := strings.NewReader("krachakonda")
	var builder strings.Builder
	processData(r, &builder)
	MyPrint("strings builder content: %s", builder.String())
	cr := strings.NewReader("copystring")
	copyData(cr, &builder)
	MyPrint("copied data: %s", builder.String())
}
