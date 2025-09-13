package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"management-produk/helper"
	"management-produk/model/domain"
	"time"
)

type ProdukRepositoryImpl struct{

}

func NewProdukRepository() ProdukRepository{
	return &ProdukRepositoryImpl{}
}

func (repository *ProdukRepositoryImpl) GetAllProduk(ctx context.Context,tx *sql.Tx)([]domain.Produk,error){
	dbName := ctx.Value("db_name").(string)
	query := fmt.Sprintf("select id,nama,harga,created_at from %s.produk",dbName)
	row,err := tx.QueryContext(ctx,query)
	helper.PanicIfError(err)
	defer row.Close()
	produks := []domain.Produk{}
	for row.Next(){
		produk := domain.Produk{}
		err := row.Scan(&produk.Id,&produk.Nama,&produk.Harga,&produk.Created_at)
		helper.PanicIfError(err)
		produks = append(produks, produk)
	}

	return produks,nil
}

func(repositoty *ProdukRepositoryImpl)CreateProduk(ctx context.Context,tx *sql.Tx,produk domain.Produk)(domain.Produk,error){
	dbName := ctx.Value("db_name").(string)
	exec := fmt.Sprintf("insert into %s.produk(nama,harga,created_at) values (?,?,?)",dbName)
	row,err := tx.ExecContext(ctx,exec,produk.Nama,produk.Harga,time.Now())
	helper.PanicIfError(err)
	idProduk,err:= row.LastInsertId()
	helper.PanicIfError(err)
	produk.Id = int(idProduk)
	return produk,nil
}

func (repository *ProdukRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, IdProduk int) (domain.Produk, error) {
    dbName := ctx.Value("db_name").(string)
    query := fmt.Sprintf("SELECT id, nama, harga, created_at FROM %s.produk WHERE id=?", dbName)
    
    row, err := tx.QueryContext(ctx, query, IdProduk)
    if err != nil {
        return domain.Produk{}, err 
    }
    defer row.Close()

    produk := domain.Produk{}
    if row.Next() {
        err := row.Scan(&produk.Id, &produk.Nama, &produk.Harga, &produk.Created_at)
        if err != nil {
            return domain.Produk{}, err
        }
        return produk, nil
    }
    return domain.Produk{}, errors.New("produk tidak ditemukan")
}

func(repository *ProdukRepositoryImpl)Delete(ctx context.Context,tx *sql.Tx,idProduk int)error{
	dbName := ctx.Value("db_name").(string)
	exec := fmt.Sprintf("delete from %s.produk where id = ?",dbName)
	row,err := tx.ExecContext(ctx,exec,idProduk)
	rowsAffected,err := row.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0{
		return errors.New("barang tidak ditemukan/sudah dihapus")
	}

	return nil
	
}