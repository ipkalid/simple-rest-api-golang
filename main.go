package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/ipkalid/order-api/app"
)

func main() {
	app := app.NewApp()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		panic(err)
		// fmt.Printf("Error: %v\n", err)
	}
}
