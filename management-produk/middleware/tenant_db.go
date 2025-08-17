package middleware

import (
	"management-produk/helper"
	"management-produk/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)



func TenantMiddleware(tenantService service.TenantService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		apiKey := c.Get("X-API-Key")
		if apiKey == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "API key required",
			})
		}

		tenantInfo, err := tenantService.GetInfoTenant(c.Context(), apiKey)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid API key",
			})
		}

		// buat DSN
		helper.PanicIfError(err)
		c.Locals("db_name",tenantInfo.DBName)
		return c.Next()
	}
}
