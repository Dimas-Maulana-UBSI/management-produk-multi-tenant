package main

import (
	"management-produk/app"
)

func main() {
	app := app.NewRouter()
    app.Static("/", "./")
	app.Listen("localhost:3000")
}