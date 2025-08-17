package repository

import (
	"database/sql"
	"management-produk/model/domain"
	"context"
)

type ProdukRepository interface {
	GetAllProduk(ctx context.Context,tx *sql.Tx)([]domain.Produk,error)
	CreateProduk(ctx context.Context,tx *sql.Tx,produk domain.Produk)(domain.Produk,error)
}