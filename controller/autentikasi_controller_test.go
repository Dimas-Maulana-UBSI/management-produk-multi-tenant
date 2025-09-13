package controller

// import (
// 	"bytes"
// 	"management-produk/app"
// 	"management-produk/repository"
// 	"management-produk/service"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/stretchr/testify/assert"
// )

// var dbMaster = app.NewMasterDB()

// var autentikasiRepository = repository.NewAutentikasiRepository()
// var autentikasiService = service.NewAutentikasiService(dbMaster,autentikasiRepository)
// var autentikasiController = NewAutentikasiController(autentikasiService)
// func TestLoginController(t *testing.T){
// 	app := fiber.New()

// 	// register route
// 	app.Post("/login", autentikasiController.Login)

// 	// isi data request (username & password harus sesuai yang ada di DB kamu)
// 	body := bytes.NewBufferString("username=dimas&password=rahasia")
// 	req := httptest.NewRequest(http.MethodPost, "/login", body)
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// jalankan request
// 	resp, err := app.Test(req)
// 	assert.NoError(t, err)

// 	// cek status code
// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// }

// func TestRegisterController(t *testing.T){
// 	app := fiber.New()

// 	// register route
// 	app.Post("/register", autentikasiController.Register)

// 	// isi data request (username & password harus sesuai yang ada di DB kamu)
// 	body := bytes.NewBufferString("username=pt_sejahtera&password=kamudandia")
// 	req := httptest.NewRequest(http.MethodPost, "/register", body)
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// jalankan request
// 	resp, err := app.Test(req)
// 	assert.NoError(t, err)

// 	// cek status code
// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// }