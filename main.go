package main

import (
	"fmt"

	"github.com/ipkalid/order-api/app"
)

func main() {
	app := app.NewApp()

	fmt.Println("app start at http://localhost:3000")
	err := app.Start()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
