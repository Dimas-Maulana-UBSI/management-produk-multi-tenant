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
	produkGroup.Get("/",produkController.GetProduk)
	produkGroup.Get("/:idProduk",produkController.GetById)
	produkGroup.Post("/",produkController.CreateProduk)
	produkGroup.Delete("/:idProduk",produkController.Delete)
	produkGroup.Put("/:idProduk",produkController.Update)

}