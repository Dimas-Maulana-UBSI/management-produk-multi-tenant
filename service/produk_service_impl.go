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

func (service *ProdukServcieImpl)GetAllProduk(ctx context.Context)([]web.ProdukResponse,error){
	tx,err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	response,err := service.ProdukRepository.GetAllProduk(ctx,tx)
	helper.PanicIfError(err)
	produks := []web.ProdukResponse{}
	for _,produk := range response {
		produks = append(produks, helper.ToProdukResponse(produk))
	}
	return produks,nil
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