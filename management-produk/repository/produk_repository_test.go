package repository

// import (
// 	"context"
// 	"fmt"
// 	"management-produk/app"
// 	"management-produk/helper"
// 	"management-produk/model/domain"
// 	"testing"
// 	"time"
// )
// var dbTenant = app.NewTenantDb()
// var produkRepository = NewProdukRepository()

// func TestProdukGetall(t *testing.T){
// 	tx, err := dbTenant.Begin()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer tx.Rollback()
// 	ctx := context.WithValue(context.Background(), "db_name", "dimas")
// 	produks, err := produkRepository.GetAllProduk(ctx, tx)
// 	helper.PanicIfError(err)
// 	fmt.Println(produks)
// }
// func TestCreateProduk(t *testing.T){
// 	tx, err := dbTenant.Begin()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	ctx := context.WithValue(context.Background(), "db_name", "dimas")
// 	defer tx.Rollback()
// 	produk := domain.Produk{
// 		Nama: "kucing",
// 		Harga: 6000000,
// 		Created_at: time.Now(),
// 	}
// 	response, err := produkRepository.CreateProduk(ctx, tx,produk)
// 	tx.Commit()
// 	helper.PanicIfError(err)
// 	fmt.Println(response)
// }