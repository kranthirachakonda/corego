package main

import "fmt"

func MyPrint(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}
