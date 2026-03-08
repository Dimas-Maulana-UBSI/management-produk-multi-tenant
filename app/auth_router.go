package app

import (
	"database/sql"
	"management-produk/controller"
	"management-produk/repository"
	"management-produk/service"
	"time"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func NewAuthRouter(app *fiber.App, dbMaster *sql.DB) {
	autentikasiRepository := repository.NewAutentikasiRepository()
	autentikasiService := service.NewAutentikasiService(dbMaster, autentikasiRepository)
	autentikasiController := controller.NewAutentikasiController(autentikasiService)

	registerLimiter := limiter.New(limiter.Config{
		Max:        1,
		Expiration: 10 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "terlalu banyak percobaan registrasi, coba lagi dalam 10 menit",
			})
		},
	})

	loginLimiter := limiter.New(limiter.Config{
		Max:        5,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "terlalu banyak percobaan login, coba lagi dalam 1 menit",
			})
		},
	})

	app.Get("/login", autentikasiController.LoginView)
	app.Post("/login", loginLimiter, autentikasiController.Login)
	app.Get("/register", autentikasiController.RegisterView)
	app.Post("/register", registerLimiter, autentikasiController.Register)
	app.Get("/home", autentikasiController.HomeView)
	app.Post("/logout", autentikasiController.Logout)
	app.Get("/docs/*", swagger.New(swagger.Config{
		URL: "/apispec.json",
	}))
}