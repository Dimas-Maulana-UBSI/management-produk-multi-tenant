package controller

import "github.com/gofiber/fiber/v2"

type AutentikasiController interface {
	Login(ctx *fiber.Ctx) error
	LoginView(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	RegisterView(ctx *fiber.Ctx) error
}