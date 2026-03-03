package middleware

import (
	"context"
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

		tenantInfo, err := tenantService.GetInfoTenant(c.UserContext(), apiKey)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid API key",
			})
		}
		ctx := context.WithValue(
			c.UserContext(),
			"db_name",
			tenantInfo.DBName,
		)

		c.SetUserContext(ctx)

		return c.Next()
	}
}
