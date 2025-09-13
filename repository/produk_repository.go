package repository

import (
	"database/sql"
	"management-produk/model/domain"
	"context"
)

type ProdukRepository interface {
	GetAllProduk(ctx context.Context,tx *sql.Tx)([]domain.Produk,error)
	CreateProduk(ctx context.Context,tx *sql.Tx,produk domain.Produk)(domain.Produk,error)
	GetById(ctx context.Context,tx *sql.Tx,IdProduk int)(domain.Produk,error)
	Delete(ctx context.Context,tx *sql.Tx,idProduk int)error
	Update(ctx context.Context,tx *sql.Tx,idProduk int,produk domain.Produk)(domain.Produk,error)
}