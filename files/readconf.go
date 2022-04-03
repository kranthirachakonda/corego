package main

import (
	"fmt"
	"os"
	"path/filepath"
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

	homePath, err := os.UserHomeDir()
	if err == nil {
		homePath = filepath.Join(homePath, "workspace", "code", "tempfile.json")
	}
	fmt.Println("Path:", homePath)
	fmt.Println("dir:", filepath.Dir(homePath))
	fmt.Println(filepath.Ext(homePath))

}
