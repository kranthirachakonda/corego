package main

import (
	"fmt"
	"os"
)

func loadConf() (err error) {
	data, err := os.ReadFile("config.json")
	if err == nil {
		fmt.Println(string(data))
	}
	return

}

func main() {
	err := loadConf()
	if err != nil {
		fmt.Println("Err loading config", err.Error())
	}
}
