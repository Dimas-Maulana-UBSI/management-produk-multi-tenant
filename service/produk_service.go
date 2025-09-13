package service

import (
	"context"
	"management-produk/model/web"
)

type ProdukServcie interface {
	GetAllProduk(ctx context.Context)([]web.ProdukResponse,error)
	CreateProduk(ctx context.Context,produk web.ProdukRequest)(web.ProdukResponse,error)
	GetById(ctx context.Context,IdProduk int)(web.ProdukResponse,error)
	Delete(ctx context.Context,idProduk int)error
	Update(ctx context.Context,idProduk int,produk web.ProdukRequest)(web.ProdukResponse,error)
}