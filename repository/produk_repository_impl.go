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

func (repository *ProdukRepositoryImpl) GetAllProduk(ctx context.Context, tx *sql.Tx,limit int,offset int) ([]domain.Produk, error) {
	dbName := ctx.Value("db_name").(string)

	if err := helper.ValidateDBName(dbName); err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT id, nama, harga, created_at FROM %s.produk LIMIT ? OFFSET ?", dbName)
	rows, err := tx.QueryContext(ctx, query,limit,offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	produks := []domain.Produk{}
	for rows.Next() {
		produk := domain.Produk{}
		err := rows.Scan(&produk.Id, &produk.Nama, &produk.Harga, &produk.Created_at)
		if err != nil {
			return nil, err
		}
		produks = append(produks, produk)
	}

	if err := rows.Err(); err != nil {
		return nil, err
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
		return domain.Produk{}, err
	}

	idProduk, err := row.LastInsertId()
	if err != nil {
		return domain.Produk{}, err
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
			return domain.Produk{}, errors.New("produk tidak ditemukan")
		}
		return domain.Produk{}, err
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
		return err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("barang tidak ditemukan/sudah dihapus")
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
		return domain.Produk{}, err
	}

	rowAffected, err := row.RowsAffected()
	if err != nil {
		return domain.Produk{}, err
	}
	if rowAffected == 0 {
		return domain.Produk{}, errors.New("produk tidak ditemukan")
	}

	produk.Id = idProduk
	return produk, nil
}

func (repository *ProdukRepositoryImpl) CountProduk(ctx context.Context,tx *sql.Tx) (int, error) {

    dbNameValue := ctx.Value("db_name")
    dbName, ok := dbNameValue.(string)
    if !ok || dbName == "" {
        return 0, errors.New("invalid db name")
    }

    if err := helper.ValidateDBName(dbName); err != nil {
        return 0, err
    }

    query := fmt.Sprintf("SELECT COUNT(*) FROM %s.produk", dbName)

    row := tx.QueryRowContext(ctx, query)

    var total int
    if err := row.Scan(&total); err != nil {
        return 0, err
    }

    return total, nil
}
