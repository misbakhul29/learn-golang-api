# API Belajar Go

RESTful API modern, aman, dan berkinerja tinggi yang dibangun menggunakan **Golang** dengan framework **Huma v2**. Proyek ini dirancang dengan prinsip **Clean Architecture**, terintegrasi dengan **Redis** untuk *high-performance caching*, dan dilengkapi dengan *client-IP-based rate limiting*.

## 🚀 Fitur Unggulan

- **Huma v2 Framework**: Implementasi REST API modern dengan validasi cepat, pemetaan DTO otomatis, dan dukungan OpenAPI 3.1 & Swagger UI out-of-the-box.
- **Clean Architecture Principles**: Pemisahan modul yang terstruktur antara Handler/Controller, Service (Business Logic), Repository (Data Access), dan DTO.
- **Client IP-Based Rate Limiting**: Pembatasan request yang adil per-IP menggunakan algoritma Token Bucket yang thread-safe, lengkap dengan goroutine background untuk pembersihan otomatis (*garbage collection*) memori tak terpakai.
- **High-Performance Redis Caching**:
  - **Cache-Aside Pattern (Lazy Loading)**: Caching data on-demand untuk list user, detail user, list post, dan detail post demi performa respon sub-milidetik.
  - **Active Cache Invalidation**: Otomatis membersihkan (*evict*) cache terkait secara real-time pada operasi pembuatan, pengubahan, atau penghapusan data (mencegah data basi).
- **Flexible Preloading**: Mendukung parameter query opsional `include_users=true` pada GET posts untuk menampilkan relasi user secara instan dan efisien.
- **OpenAPI & Swagger Documentation**: Halaman dokumentasi API interaktif yang otomatis dibuat dan disinkronisasikan di endpoint `/docs`.

## 🛠️ Tech Stack

- **Language**: Go (Golang)
- **API Framework**: [Huma v2](https://huma.rocks/)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: PostgreSQL (Relasional)
- **In-Memory Store**: Redis (Cache)
- **Rate Limiting**: `golang.org/x/time/rate`
- **Hot Reload**: [Air](https://github.com/air-verse/air)

## 📁 Struktur Proyek

```text
.
├── internal/
│   ├── config/         # Konfigurasi environment (.env)
│   ├── database/       # Koneksi database & database seeder
│   ├── dto/            # Data Transfer Object (Request/Response contracts)
│   ├── handlers/       # Handler API & routing endpoint (Controller)
│   ├── middlewares/    # Middleware (API Key, Rate Limiter per-IP)
│   ├── models/         # Entity / Struct model database GORM
│   ├── repositories/   # Lapisan akses data GORM (Repository)
│   └── services/       # Lapisan bisnis & utilitas
│       ├── api_services/ # Implementasi business logic (User, Post)
│       └── redis/        # Koneksi, key registry, & store wrapper Redis
├── cmd/                # Entrypoint alternatif
├── main.go             # Entrypoint utama aplikasi
├── .env                # Konfigurasi rahasia (.gitignore)
└── example.env         # Contoh file konfigurasi environment
```

## ⚙️ Persiapan & Instalasi

### 1. Clone Repositori
```bash
git clone https://github.com/misbakhul29/learn-golang-api.git
cd belajar
```

### 2. Konfigurasi Environment
Salin file `example.env` menjadi `.env` dan sesuaikan nilainya (termasuk kredensial DB dan Redis):
```bash
cp example.env .env
```

### 3. Jalankan Docker (Opsional)
Pastikan PostgreSQL dan Redis Anda sudah berjalan di port yang sesuai dengan konfigurasi `.env`.

### 4. Instal Dependensi
```bash
go mod tidy
```

### 5. Jalankan Aplikasi
Untuk pengembangan lokal dengan hot reload:
```bash
air
```
Atau jalankan langsung menggunakan perintah bawaan Go:
```bash
go run main.go
```
Aplikasi akan berjalan secara default di `http://localhost:8080`.

## 📖 Dokumentasi API

Setelah aplikasi berjalan, dokumentasi interaktif dapat diakses langsung pada browser:
- **Swagger UI**: `http://localhost:8080/docs`
- **Schemas**: `http://localhost:8080/schemas`

## 🛡️ Contoh Penggunaan (cURL)

Setiap request wajib menyertakan header `X-API-Key` sesuai konfigurasi di `.env`.

### A. Endpoint Users

**Mendapatkan Semua User (Cached):**
```bash
curl --request GET \
  --url http://localhost:8080/api/users \
  --header 'X-API-Key: mysecretkey'
```

**Mendapatkan Detail User (Cached):**
```bash
curl --request GET \
  --url http://localhost:8080/api/users/USER_UUID \
  --header 'X-API-Key: mysecretkey'
```

---

### B. Endpoint Posts

**Mendapatkan Semua Post (Tanpa Data User):**
```bash
curl --request GET \
  --url http://localhost:8080/api/posts \
  --header 'X-API-Key: mysecretkey'
```

**Mendapatkan Semua Post beserta Detail Penulis (Preloaded):**
```bash
curl --request GET \
  --url http://localhost:8080/api/posts?include_users=true \
  --header 'X-API-Key: mysecretkey'
```

---
Dibuat dengan ❤️ oleh [Misbakhul Munir](https://misbakhul.my.id)
