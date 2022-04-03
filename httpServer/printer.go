package main

import "fmt"

func MyPrintlin(tmpl string, values ...interface{}) {
	fmt.Printf(tmpl+"\n", values...)
}
