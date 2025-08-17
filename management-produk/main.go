package main

import (
	"management-produk/app"
	"management-produk/controller"
	"management-produk/middleware"
	"management-produk/repository"
	"management-produk/service"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
)

func main() {
	dbMaster := app.NewMasterDB()
	dbTenant := app.NewTenantDb()
	engine := mustache.New("./views",".mustache")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	autentikasiRepository := repository.NewAutentikasiRepository()
	autentikasiService := service.NewAutentikasiService(dbMaster,autentikasiRepository)
	autentikasiController := controller.NewAutentikasiController(autentikasiService)

	app.Get("/login",autentikasiController.LoginView)
	app.Post("/login",autentikasiController.Login)
	app.Get("/docs/*", swagger.New(swagger.Config{
        URL: "/apispec.json", // arahkan ke file apispec.json
    }))
	app.Get("/register",autentikasiController.RegisterView)
	app.Post("/register",autentikasiController.Register)
	produkRepository := repository.NewProdukRepository()
	produkService := service.NewProdukService(produkRepository,dbTenant)
	produkController := controller.NewProdukController(produkService)

	tenantRepository := repository.NewTenantInfoRepository()
	tenantService := service.NewTenantInfoService(tenantRepository,dbMaster)

	produkGroup := app.Group("/produk",middleware.TenantMiddleware(tenantService))
	produkGroup.Get("/getAllProduk",produkController.GetProduk)
	produkGroup.Post("/CreateProduk",produkController.CreateProduk)


    
    app.Static("/", "./") // supaya apispec.j

	app.Listen("localhost:3000")
}