package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	name   string = "SomeName"
	course string = "basic golang"
	module string = "4"
	clip   int    = 2
)

func main() {
	fmt.Println("Name & Course are set to ", name, "and", course, ".")
	fmt.Println("Module & Clip are set to ", module, "and", clip, ".")
	fmt.Println("Name is of Type: ", reflect.TypeOf(name))
	fmt.Println("Module is of Type: ", reflect.TypeOf(module))

	iModule, err := strconv.Atoi(module)

	if err == nil {
		total := iModule + clip
		fmt.Println("Module + Clip: ", total)
	}

}
