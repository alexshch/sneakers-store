package main

import (
	"context"
	"fmt"
	"log"
	"sneakers/internal/app"
)

func main() {
	ctx := context.Background()
	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app %s", err.Error())
	}
	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app %s", err.Error())
	}
	fmt.Println("run sneakers api...")
}
