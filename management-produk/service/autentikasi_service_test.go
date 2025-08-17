package service_test

import (
	"context"
	"fmt"
	"management-produk/app"
	"management-produk/model/web"
	"management-produk/repository"
	"management-produk/service"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var db = app.NewMasterDB()
var authRepo = repository.NewAutentikasiRepository()
var authService = service.NewAutentikasiService(db, authRepo)
func TestRegistrasiIntegration(t *testing.T) {
    request := web.RegistrasiRequest{
        Name:   "tenant_test",
        Password: "cobacoba",
       
    }

    ctx := context.Background()
    response, err := authService.Registrasi(ctx, request)

    assert.Nil(t, err)
    assert.Equal(t, request.Name, response.Name)

    fmt.Println(response)
}

func TestLogin(t *testing.T){
    request := web.LoginRequest{
        Name: "Tenant A",
        Password: "pass_a",
    }
    ctx := context.Background()
    response,err := authService.Login(ctx,request)
    assert.Nil(t,err)
    assert.Equal(t,"apikey123",response.ApiKey)

}