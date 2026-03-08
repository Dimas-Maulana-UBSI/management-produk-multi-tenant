package controller

import (
	"management-produk/helper"
	"management-produk/model/web"
	"management-produk/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store = session.New()

type AutentikasiControllerImpl struct {
	AutentikasiService service.AutentikasiService
}

func NewAutentikasiController(service service.AutentikasiService) AutentikasiController {
	return &AutentikasiControllerImpl{
		AutentikasiService: service,
	}
}

func (c *AutentikasiControllerImpl) LoginView(ctx *fiber.Ctx) error {
	sess, err := store.Get(ctx)
	if err == nil && sess.Get("name") != nil {
		return ctx.Redirect("/home")
	}
	return ctx.Render("login", fiber.Map{
		"Title": "Login Page",
	})
}

func (controller *AutentikasiControllerImpl) Login(ctx *fiber.Ctx) error {
	name := ctx.FormValue("username")
	password := ctx.FormValue("password")
	request := web.LoginRequest{
		Name:     name,
		Password: password,
	}
	response, err := controller.AutentikasiService.Login(ctx.Context(), request)
	if err != nil {
		return helper.RespondFiberError(ctx, err)
	}

	sess, err := store.Get(ctx)
	if err != nil {
		return helper.RespondFiberError(ctx, helper.Internal("gagal membuat session", err))
	}
	sess.Set("name", response.Name)
	sess.Set("api_key", response.ApiKey)
	if err := sess.Save(); err != nil {
		return helper.RespondFiberError(ctx, helper.Internal("gagal menyimpan session", err))
	}

	return ctx.Redirect("/home")
}

func (controller *AutentikasiControllerImpl) RegisterView(ctx *fiber.Ctx) error {
	sess, err := store.Get(ctx)
	if err == nil && sess.Get("name") != nil {
		return ctx.Redirect("/home")
	}
	return ctx.Render("register", fiber.Map{})
}

func (controller *AutentikasiControllerImpl) Register(ctx *fiber.Ctx) error {
	name := ctx.FormValue("username")
	password := ctx.FormValue("password")
	request := web.RegistrasiRequest{
		Name:     name,
		Password: password,
	}
	_, err := controller.AutentikasiService.Registrasi(ctx.Context(), request)
	if err != nil {
		return helper.RespondFiberError(ctx, err)
	}
	return ctx.Redirect("/login")
}

func (controller *AutentikasiControllerImpl) HomeView(ctx *fiber.Ctx) error {
	sess, err := store.Get(ctx)
	if err != nil || sess.Get("name") == nil {
		return ctx.Redirect("/login")
	}
	return ctx.Render("home", fiber.Map{
		"name":    sess.Get("name"),
		"api_key": sess.Get("api_key"),
	})
}

func (controller *AutentikasiControllerImpl) Logout(ctx *fiber.Ctx) error {
	sess, err := store.Get(ctx)
	if err != nil {
		return ctx.Redirect("/login")
	}
	if err := sess.Destroy(); err != nil {
		return helper.RespondFiberError(ctx, helper.Internal("gagal logout", err))
	}
	return ctx.Redirect("/login")
}