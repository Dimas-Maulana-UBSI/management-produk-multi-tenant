project untuk management produk multi tenant

## persiapan database
sebelum run server buat database mysal dengan nama tenants terlebih dahulu

lalu jalankan perintah
migrate -path db/migrations -database "mysql://root:@tcp(localhost:3306)/tenants" up

lalu untuk rollback jika migrations gagal
migrate -path db/migrations -database "mysql://root:@tcp(localhost:3306)/tenants" down


## endpoint api
untuk endpointnya
/register digunakan untuk membuat akun
/login digunan untuk login dan mengambil api keynya

tambahkan di header json x-api-key dengan api key yang sudah di ambil

GET
/produk untuk mengambil semua produk
/produk/1 untuk mengambil produk berdasarkan id misalnya 1

POST
/produk untuk membuat produk baru

DELETE
/produk/1 untuk menghapus produk berdasarkan id misalnya 1

struktur jsonnya
{
    "name":"nama_produk",
    "harga":200000,
}




