package main

/*
cd {TO DIRECTORY THAT HAS go.mod file}
go build ./cmd/gows <- this creates gows binary in current directory
./gows <- runs the program
*/

import (
	"fmt"
	"net/http"

	"github.com/imayavgi/gows/internal/pkg/controllers"
)

func main() {
	fmt.Println("Go Web Services command ")

	controllers.RegisterControllers()
	http.ListenAndServe(":9999", nil)
}
