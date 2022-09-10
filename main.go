package main

import (
	"fita-test-02/routes"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Print(err)
	}
	router := routes.InitializeRoute()
	router.Run()
}
