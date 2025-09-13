package controller

import "github.com/gofiber/fiber/v2"

type ProdukController interface {
	GetProduk(ctx *fiber.Ctx)error
	CreateProduk(ctx *fiber.Ctx)error
	GetById(ctx *fiber.Ctx)error
	Delete(ctx *fiber.Ctx)error
}