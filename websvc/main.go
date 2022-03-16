package main

import (
	"fmt"

	"github.com/kzdevops/corego/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Kraz",
		LastName:  "Rach",
	}
	fmt.Println(u)
}
