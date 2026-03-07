# Management Produk Multi-Tenant

Layanan backend Go untuk manajemen produk dengan arsitektur multi-tenant. Setiap tenant memiliki database terpisah yang diisolasi, diakses melalui API key authentication.

## Daftar Isi
- [Fitur](#fitur)
- [Arsitektur & Struktur Direktori](#arsitektur--struktur-direktori)
- [Persyaratan](#persyaratan)
- [Instalasi & Jalankan](#instalasi--jalankan)
- [Konfigurasi](#konfigurasi)
- [Persiapan Database](#persiapan-database)
- [Penggunaan API](#penggunaan-api)
- [Testing](#testing)
- [Kontak](#kontak)

## Fitur
- **Multi-tenant** — setiap tenant memiliki database terpisah yang terisolasi
- **API Key Authentication** — setiap tenant diidentifikasi melalui API key unik
- **Password Hashing** — password tenant dienkripsi menggunakan bcrypt
- **Pagination** — endpoint list produk mendukung pagination
- **Clean Architecture** — pemisahan layer controller, service, dan repository
- **Error Handling** — penanganan error yang konsisten di semua layer

## Arsitektur & Struktur Direktori
Proyek mengikuti pola layered architecture:
```
management-produk-multi-tenant/
├── app/          — router dan setup aplikasi
├── controller/   — handler HTTP
├── service/      — logika bisnis
├── repository/   — akses data (DB)
├── middleware/   — API key & tenant middleware
├── model/        — definisi domain dan web response
├── helper/       — error handling, validator, utilities
└── db/migrations — skrip migrasi database
```

## Persyaratan
- Go 1.20+
- MySQL 8.0+

## Instalasi & Jalankan

1. Clone repository
```bash
git clone <repo-url>
cd management-produk-multi-tenant
```

2. Install dependency
```bash
go mod download
```

3. Buat file `.env` dari template
```bash
cp .env.example .env
```

4. Isi `.env` sesuai konfigurasi MySQL kamu
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=tenants
APP_PORT=3000
```

5. Jalankan persiapan database — lihat bagian [Persiapan Database](#persiapan-database)

6. Jalankan aplikasi
```bash
go run main.go
```

Aplikasi berjalan di `http://localhost:3000`

## Konfigurasi

| Variable | Keterangan | Default |
|---|---|---|
| `DB_HOST` | Host MySQL | localhost |
| `DB_PORT` | Port MySQL | 3306 |
| `DB_USER` | Username MySQL | root |
| `DB_PASSWORD` | Password MySQL | - |
| `DB_NAME` | Nama database master | tenants |
| `APP_PORT` | Port aplikasi | 3000 |

## Persiapan Database

1. Buat database MySQL dengan nama `tenants`
```sql
CREATE DATABASE tenants;
```

2. Install migrate tool (jika belum ada)
```bash
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

3. Jalankan migrasi
```bash
migrate -path db/migrations -database "mysql://root:password@tcp(localhost:3306)/tenants" up
```

4. Rollback jika diperlukan
```bash
migrate -path db/migrations -database "mysql://root:password@tcp(localhost:3306)/tenants" down
```

> Sesuaikan `root` dan `password` dengan kredensial MySQL kamu

## Penggunaan API

### Auth
| Method | Endpoint | Keterangan |
|---|---|---|
| POST | `/auth/register` | Registrasi tenant baru |
| POST | `/auth/login` | Login tenant |

### Produk
> Semua endpoint produk memerlukan header `X-API-Key`

| Method | Endpoint | Keterangan |
|---|---|---|
| GET | `/api/produk` | List produk (support pagination) |
| POST | `/api/produk` | Tambah produk |
| GET | `/api/produk/:idProduk` | Detail produk |
| PUT | `/api/produk/:idProduk` | Update produk |
| DELETE | `/api/produk/:idProduk` | Hapus produk |

**Contoh request list produk:**
```bash
curl -X GET "http://localhost:3000/api/produk?page=1&limit=10" \
  -H "X-API-Key: your_api_key"
```

**Contoh request tambah produk:**
```bash
curl -X POST "http://localhost:3000/api/produk" \
  -H "X-API-Key: your_api_key" \
  -H "Content-Type: application/json" \
  -d '{"name": "Produk A", "harga": 10000}'
```

## Testing
```bash
go test ./...
```