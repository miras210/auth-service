package main

import (
	"auth-service/internal/app"
	"os"
)

func main() {
	err := app.Run()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
