package repository

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"management-produk/model/domain"

	"github.com/DATA-DOG/go-sqlmock"
)

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}
	return db, mock
}

func TestGetAllProduk_Success(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("begin tx: %v", err)
	}

	ctx := context.WithValue(context.Background(), "db_name", "db_test")

	limit := 10
	offset := 0
	query := "SELECT id, nama, harga, created_at FROM db_test.produk LIMIT ? OFFSET ?"
	rows := sqlmock.NewRows([]string{"id", "nama", "harga", "created_at"}).
		AddRow(1, "produk A", 10000, time.Now()).
		AddRow(2, "produk B", 20000, time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(limit, offset).WillReturnRows(rows)

	repo := NewProdukRepository()
	produks, err := repo.GetAllProduk(ctx, tx, limit, offset)
	if err != nil {
		t.Fatalf("GetAllProduk returned error: %v", err)
	}
	if len(produks) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(produks))
	}

	mock.ExpectRollback()
	if err := tx.Rollback(); err != nil {
		t.Fatalf("rollback tx: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestCreateProduk_Success(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("begin tx: %v", err)
	}

	ctx := context.WithValue(context.Background(), "db_name", "db_test")

	produk := domain.Produk{Nama: "kucing", Harga: 6000000}
	exec := "INSERT INTO db_test.produk(nama, harga, created_at) VALUES (?, ?, ?)"

	mock.ExpectExec(regexp.QuoteMeta(exec)).WithArgs(produk.Nama, produk.Harga, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(5, 1))

	repo := NewProdukRepository()
	res, err := repo.CreateProduk(ctx, tx, produk)
	if err != nil {
		t.Fatalf("CreateProduk returned error: %v", err)
	}
	if res.Id != 5 {
		t.Fatalf("expected id 5, got %d", res.Id)
	}

	mock.ExpectRollback()
	if err := tx.Rollback(); err != nil {
		t.Fatalf("rollback tx: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestGetById_Update_Delete_Count_Success(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("begin tx: %v", err)
	}

	ctx := context.WithValue(context.Background(), "db_name", "db_test")

	// GetById
	getQuery := "SELECT id, nama, harga, created_at FROM db_test.produk WHERE id = ?"
	row := sqlmock.NewRows([]string{"id", "nama", "harga", "created_at"}).AddRow(10, "produk X", 5000, time.Now())
	mock.ExpectQuery(regexp.QuoteMeta(getQuery)).WithArgs(10).WillReturnRows(row)

	// Update
	updateExec := "UPDATE db_test.produk SET nama = ?, harga = ? WHERE id = ?"
	mock.ExpectExec(regexp.QuoteMeta(updateExec)).WithArgs("produk X updated", 7000, 10).WillReturnResult(sqlmock.NewResult(0, 1))

	// Delete
	deleteExec := "DELETE FROM db_test.produk WHERE id = ?"
	mock.ExpectExec(regexp.QuoteMeta(deleteExec)).WithArgs(10).WillReturnResult(sqlmock.NewResult(0, 1))

	// Count
	countQuery := "SELECT COUNT(*) FROM db_test.produk"
	countRow := sqlmock.NewRows([]string{"count"}).AddRow(3)
	mock.ExpectQuery(regexp.QuoteMeta(countQuery)).WillReturnRows(countRow)

	repo := NewProdukRepository()

	// GetById
	p, err := repo.GetById(ctx, tx, 10)
	if err != nil {
		t.Fatalf("GetById error: %v", err)
	}
	if p.Id != 10 {
		t.Fatalf("expected id 10, got %d", p.Id)
	}

	// Update
	updatedProd := domain.Produk{Nama: "produk X updated", Harga: 7000}
	up, err := repo.Update(ctx, tx, 10, updatedProd)
	if err != nil {
		t.Fatalf("Update error: %v", err)
	}
	if up.Id != 10 {
		t.Fatalf("expected updated id 10, got %d", up.Id)
	}

	// Delete
	if err := repo.Delete(ctx, tx, 10); err != nil {
		t.Fatalf("Delete error: %v", err)
	}

	// Count
	total, err := repo.CountProduk(ctx, tx)
	if err != nil {
		t.Fatalf("CountProduk error: %v", err)
	}
	if total != 3 {
		t.Fatalf("expected count 3, got %d", total)
	}

	mock.ExpectRollback()
	if err := tx.Rollback(); err != nil {
		t.Fatalf("rollback tx: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}
