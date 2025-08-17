package app

import (
	"database/sql"
	"management-produk/helper"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func NewMasterDB() *sql.DB {
    dsn := "root@tcp(localhost:3306)/management_produk_test?parseTime=true"

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(err)
    }

    db.SetMaxIdleConns(5)
    db.SetMaxOpenConns(20)
    db.SetConnMaxLifetime(60 * time.Minute)
    db.SetConnMaxIdleTime(10 * time.Minute)

    if err := db.Ping(); err != nil {
        panic(err)
    }

    return db
}

// func GetOrCreatePool(tenantKey,user,host string,port int,dbname string)(*sql.DB ,error){
//     mu.RLock()
//     if db,ok := tenantPools[tenantKey];ok{
//         mu.RUnlock()
//         return db,nil
//     }
//     mu.RUnlock()
//     dsn := fmt.Sprintf("%s@tcp(%s:%d)/%s?parseTime=true", user, host, port, dbname)
// 	db, err := sql.Open("mysql", dsn)
//     helper.PanicIfError(err)
//     db.SetMaxIdleConns(5)
// 	db.SetMaxOpenConns(20)
// 	db.SetConnMaxLifetime(60 * time.Minute)
// 	db.SetConnMaxIdleTime(10 * time.Minute)

// 	// simpan ke cache
// 	mu.Lock()
// 	tenantPools[tenantKey] = db
// 	mu.Unlock()

// 	return db, nil
// }

func NewTenantDb() *sql.DB {
    dsn := "root@tcp(localhost:3306)/?parseTime=true"
    db,err := sql.Open("mysql",dsn)
    helper.PanicIfError(err)

    db.SetMaxIdleConns(20)
    db.SetMaxOpenConns(30)
    db.SetConnMaxLifetime(60 * time.Minute)
    db.SetConnMaxIdleTime(10 * time.Minute)

    if err := db.Ping(); err != nil {
        panic(err)
    }

    return db
}
