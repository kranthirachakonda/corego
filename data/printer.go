package main

import "fmt"

func MyPrint(tpl string, values ...interface{}) {
	fmt.Printf(tpl+"\n", values...)
}
