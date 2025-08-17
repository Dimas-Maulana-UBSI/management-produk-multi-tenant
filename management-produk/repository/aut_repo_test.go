package repository

import (
	"context"
	"fmt"
	"management-produk/app"
	"management-produk/helper"
	"management-produk/model/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)
var db = app.NewMasterDB() 
var authRepository = NewAutentikasiRepository()

func TestAuhtRepo(t *testing.T){
	tx,err := db.Begin()
	helper.PanicIfError(err)
	ctx := context.Background()
	tenantRequest := domain.Tenant{
		Name: "tes",
		DBHost: "localhost",
		DBPort: 3306,
		DBName: "tes_test",
		
	}
	tenant,err:= authRepository.CreateDb(ctx,tx,tenantRequest)
	tx.Commit()
	assert.Nil(t,err)
	fmt.Println(tenant)
}