package main

import (
	"context"
	"fmt"

	"github.com/rishabhthakur11/orders-api.git/application"
)

func main() {
	app := application.New()
	err := app.Start((context.TODO()))
	if err != nil {
		fmt.Println("failed to start the application: %w", err)
	} else {
		fmt.Println("Application started successfully")
	}
}
