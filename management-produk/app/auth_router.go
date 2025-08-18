package app

import (
	"database/sql"
	"management-produk/controller"
	"management-produk/repository"
	"management-produk/service"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func NewAuthRouter(app *fiber.App,dbMaster *sql.DB){
	autentikasiRepository := repository.NewAutentikasiRepository()
	autentikasiService := service.NewAutentikasiService(dbMaster,autentikasiRepository)
	autentikasiController := controller.NewAutentikasiController(autentikasiService)

	app.Get("/login",autentikasiController.LoginView)
	app.Post("/login",autentikasiController.Login)
	app.Get("/docs/*", swagger.New(swagger.Config{
        URL: "/apispec.json", 
    }))
	app.Get("/register",autentikasiController.RegisterView)
	app.Post("/register",autentikasiController.Register)
}