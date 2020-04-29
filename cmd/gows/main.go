package main

/*
cd {TO DIRECTORY THAT HAS go.mod file}
go build ./cmd/gows <- this creates gows binary in current directory
./gows <- runs the program
*/

import (
	"fmt"

	"github.com/imayavgi/gows/internal/pkg/models"
)

func main() {
	fmt.Println("Go Web Services command ")
	u := models.User{
		ID:        2,
		FirstName: "Imaya",
		LastName:  "Kulothungan",
	}

	fmt.Println(u)
}
