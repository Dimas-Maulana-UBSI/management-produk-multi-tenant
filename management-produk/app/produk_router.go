package app

import (
	"database/sql"
	"management-produk/controller"
	"management-produk/middleware"
	"management-produk/repository"
	"management-produk/service"

	"github.com/gofiber/fiber/v2"
)

func NewProdukRouter(app *fiber.App,dbMaster *sql.DB){
	dbTenant := NewTenantDb()
	produkRepository := repository.NewProdukRepository()
	produkService := service.NewProdukService(produkRepository,dbTenant)
	produkController := controller.NewProdukController(produkService)

	tenantRepository := repository.NewTenantInfoRepository()
	tenantService := service.NewTenantInfoService(tenantRepository,dbMaster)

	produkGroup := app.Group("/produk",middleware.TenantMiddleware(tenantService))
	produkGroup.Get("/getAllProduk",produkController.GetProduk)
	produkGroup.Get("/getById/:idProduk",produkController.GetById)
	produkGroup.Post("/CreateProduk",produkController.CreateProduk)
	produkGroup.Delete("/delete/:idProduk",produkController.Delete)

}