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

type ProdukRepositoryImpl struct {
}

func NewProdukRepository() ProdukRepository {
	return &ProdukRepositoryImpl{}
}

func (repository *ProdukRepositoryImpl) GetAllProduk(ctx context.Context, tx *sql.Tx, limit int, offset int) ([]domain.Produk, error) {
	dbName := ctx.Value("db_name").(string)

	if err := helper.ValidateDBName(dbName); err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT id, nama, harga, created_at FROM %s.produk LIMIT ? OFFSET ?", dbName)
	rows, err := tx.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, helper.Internal("query produk failed", err)
	}
	defer rows.Close()

	produks := []domain.Produk{}
	for rows.Next() {
		produk := domain.Produk{}
		err := rows.Scan(&produk.Id, &produk.Nama, &produk.Harga, &produk.Created_at)
		if err != nil {
			return nil, helper.Internal("scan produk failed", err)
		}
		produks = append(produks, produk)
	}

	if err := rows.Err(); err != nil {
		return nil, helper.Internal("rows iteration error", err)
	}

	return produks, nil
}

func (repository *ProdukRepositoryImpl) CreateProduk(ctx context.Context, tx *sql.Tx, produk domain.Produk) (domain.Produk, error) {
	dbName := ctx.Value("db_name").(string)
	if err := helper.ValidateDBName(dbName); err != nil {
		return domain.Produk{}, err
	}

	exec := fmt.Sprintf("INSERT INTO %s.produk(nama, harga, created_at) VALUES (?, ?, ?)", dbName)
	row, err := tx.ExecContext(ctx, exec, produk.Nama, produk.Harga, time.Now())
	if err != nil {
		return domain.Produk{}, helper.Internal("insert produk failed", err)
	}

	idProduk, err := row.LastInsertId()
	if err != nil {
		return domain.Produk{}, helper.Internal("last insert id failed", err)
	}

	produk.Id = int(idProduk)
	return produk, nil
}

func (repository *ProdukRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, IdProduk int) (domain.Produk, error) {
	dbName := ctx.Value("db_name").(string)
	if err := helper.ValidateDBName(dbName); err != nil {
		return domain.Produk{}, err
	}

	query := fmt.Sprintf("SELECT id, nama, harga, created_at FROM %s.produk WHERE id = ?", dbName)
	row := tx.QueryRowContext(ctx, query, IdProduk)

	produk := domain.Produk{}
	err := row.Scan(
		&produk.Id,
		&produk.Nama,
		&produk.Harga,
		&produk.Created_at,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Produk{}, helper.NotFound("produk tidak ditemukan")
		}
		return domain.Produk{}, helper.Internal("scan produk failed", err)
	}

	return produk, nil
}

func (repository *ProdukRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, idProduk int) error {
	dbName := ctx.Value("db_name").(string)
	if err := helper.ValidateDBName(dbName); err != nil {
		return err
	}

	exec := fmt.Sprintf("DELETE FROM %s.produk WHERE id = ?", dbName)
	row, err := tx.ExecContext(ctx, exec, idProduk)
	if err != nil {
		return helper.Internal("delete produk failed", err)
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return helper.Internal("rows affected failed", err)
	}
	if rowsAffected == 0 {
		return helper.NotFound("barang tidak ditemukan/sudah dihapus")
	}

	return nil
}

func (repository *ProdukRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, idProduk int, produk domain.Produk) (domain.Produk, error) {
	dbName := ctx.Value("db_name").(string)
	if err := helper.ValidateDBName(dbName); err != nil {
		return domain.Produk{}, err
	}

	query := fmt.Sprintf("UPDATE %s.produk SET nama = ?, harga = ? WHERE id = ?", dbName)
	row, err := tx.ExecContext(ctx, query, produk.Nama, produk.Harga, idProduk)
	if err != nil {
		return domain.Produk{}, helper.Internal("update produk failed", err)
	}

	rowAffected, err := row.RowsAffected()
	if err != nil {
		return domain.Produk{}, helper.Internal("rows affected failed", err)
	}
	if rowAffected == 0 {
		return domain.Produk{}, helper.NotFound("produk tidak ditemukan")
	}

	produk.Id = idProduk
	return produk, nil
}

func (repository *ProdukRepositoryImpl) CountProduk(ctx context.Context, tx *sql.Tx) (int, error) {

	dbNameValue := ctx.Value("db_name")
	dbName, ok := dbNameValue.(string)
	if !ok || dbName == "" {
		return 0, errors.New("invalid db name")
	}

	if err := helper.ValidateDBName(dbName); err != nil {
		return 0, helper.BadRequest("invalid db name")
	}

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s.produk", dbName)

	row := tx.QueryRowContext(ctx, query)

	var total int
	if err := row.Scan(&total); err != nil {
		return 0, helper.Internal("count query failed", err)
	}

	return total, nil
}
