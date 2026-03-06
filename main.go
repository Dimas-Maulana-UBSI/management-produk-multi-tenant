package main

import (
	"management-produk/app"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	app := app.NewRouter()
	app.Static("/", "./")
	app.Listen(":" + port)
}