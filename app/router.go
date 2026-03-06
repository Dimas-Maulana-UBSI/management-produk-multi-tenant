package app

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/mustache/v2"
)

func NewRouter() *fiber.App {
	dbMaster := NewMasterDB()
	engine := mustache.New("./views", ".mustache")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(recover.New())

	NewAuthRouter(app, dbMaster)
	NewProdukRouter(app, dbMaster)

	app.Get("/docs/*", swagger.HandlerDefault)

	return app
}