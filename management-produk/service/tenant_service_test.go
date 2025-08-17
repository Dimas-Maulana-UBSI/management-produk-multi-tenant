package service

import (
	"context"
	"fmt"
	"management-produk/app"
	"management-produk/repository"
	"testing"
)
var db = app.NewMasterDB()
var tenantRepository = repository.NewTenantInfoRepository()
var tenantService = NewTenantInfoService(tenantRepository,db)

func TestTenantInfo(t *testing.T) {
    ctx := context.Background()
    APIKey := "apikey123"

    result, err := tenantService.GetInfoTenant(ctx, APIKey)
    if err != nil {
        panic(err)
    }
    fmt.Println(result.APIKey)
    fmt.Println(result.DBName)

}
