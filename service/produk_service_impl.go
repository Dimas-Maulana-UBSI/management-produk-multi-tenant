package service

import (
	"context"
	"database/sql"
	"errors"
	"management-produk/helper"
	"management-produk/model/domain"
	"management-produk/model/web"
	"management-produk/repository"
)
const (
	DefaultLimit = 10
	MaxLimit = 30
)
type ProdukServcieImpl struct {
	ProdukRepository repository.ProdukRepository
	DB *sql.DB
}

func NewProdukService(repository repository.ProdukRepository,db *sql.DB)ProdukServcie{
	return &ProdukServcieImpl{
		ProdukRepository: repository,
		DB: db,
	}
}

func (service *ProdukServcieImpl)GetAllProduk(ctx context.Context,limit int,page int)(web.ProdukPaginationResponse,error){
	tx, err := service.DB.Begin()
    if err != nil {
        return web.ProdukPaginationResponse{}, err
    }
    defer tx.Rollback()

    if limit <= 0 {
        limit = DefaultLimit
    }

    if limit > MaxLimit {
        limit = MaxLimit
    }

    if page <= 0 {
        page = 1
    }

    offset := (page - 1) * limit

    result, err := service.ProdukRepository.
        GetAllProduk(ctx, tx, limit, offset)
    if err != nil {
        return web.ProdukPaginationResponse{}, err
    }

    totalData, err := service.ProdukRepository.
        CountProduk(ctx, tx)
    if err != nil {
        return web.ProdukPaginationResponse{}, err
    }

    totalPage := 0
    if totalData > 0 {
        totalPage = (totalData + limit - 1) / limit
    }

    var produks []web.ProdukResponse
    for _, produk := range result {
        produks = append(produks, helper.ToProdukResponse(produk))
    }

    if err := tx.Commit(); err != nil {
        return web.ProdukPaginationResponse{}, err
    }

    return web.ProdukPaginationResponse{
        Data:      produks,
        Page:      page,
        Limit:     limit,
        TotalData: totalData,
        TotalPage: totalPage,
    }, nil
}

func(service *ProdukServcieImpl)CreateProduk(ctx context.Context,produk web.ProdukRequest)(web.ProdukResponse,error){
	tx,err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	request := domain.Produk{
		Nama: produk.Name,
		Harga: produk.Harga,
	}
	
	response,err := service.ProdukRepository.CreateProduk(ctx,tx,request)
	err = tx.Commit()
    if err != nil {
        return web.ProdukResponse{}, err
    }
	helper.PanicIfError(err)
	return helper.ToProdukResponse(response),nil
}

func(service *ProdukServcieImpl)GetById(ctx context.Context,IdProduk int)(web.ProdukResponse,error){
	tx,err := service.DB.Begin()
	if err != nil {
		return web.ProdukResponse{},errors.New(err.Error())
	}
	response,err := service.ProdukRepository.GetById(ctx,tx,IdProduk)
	if err != nil {
		return web.ProdukResponse{},errors.New(err.Error())
	}
	return helper.ToProdukResponse(response),nil
}

func(service *ProdukServcieImpl)Delete(ctx context.Context,idProduk int)error{
	tx,err := service.DB.Begin()
	if err != nil {
		return err
	}
	err = service.ProdukRepository.Delete(ctx, tx, idProduk)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func(service *ProdukServcieImpl)Update(ctx context.Context,idProduk int,produk web.ProdukRequest)(web.ProdukResponse,error){
	tx,err := service.DB.Begin()
	if err != nil {
		return web.ProdukResponse{},err
	}

	request := domain.Produk{
		Nama: produk.Name,
		Harga: produk.Harga,
	}
	response,err := service.ProdukRepository.Update(ctx,tx,idProduk,request)
	if err != nil {
		return web.ProdukResponse{},err
	}
	err = tx.Commit()
    if err != nil {
        return web.ProdukResponse{}, err
    }
	return helper.ToProdukResponse(response),nil
}