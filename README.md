# API Belajar Go

Proyek ini adalah API sederhana untuk manajemen user yang dibangun menggunakan **Golang** dengan framework **Huma v2**. Proyek ini dirancang dengan arsitektur yang modular, aman, dan terdokumentasi dengan baik secara otomatis.

## 🚀 Fitur Unggulan

- **Huma v2 Framework**: Implementasi REST API modern dengan dukungan OpenAPI 3.1 otomatis.
- **Security Middleware**:
  - **API Key Authentication**: Melindungi setiap endpoint menggunakan header `X-API-Key`.
  - **Rate Limiting**: Mencegah abuse dengan pembatasan jumlah request per detik (Token Bucket).
- **Standardized Response**: Semua API mengembalikan struktur JSON yang seragam (`data`, `status`, `code`, `message`).
- **Partial Updates**: Mendukung update data secara parsial (hanya field yang dikirim di JSON yang akan diupdate).
- **Database GORM**: Integrasi PostgreSQL yang tangguh dengan pemetaan model otomatis.
- **Auto-Generated Documentation**: Dokumentasi interaktif Swagger UI tersedia langsung tanpa konfigurasi manual.

## 🛠️ Tech Stack

- **Language**: Go (Golang)
- **API Framework**: [Huma v2](https://huma.rocks/)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: PostgreSQL
- **Security**: `golang.org/x/time/rate` (Rate Limiter)

## 📁 Struktur Proyek

```text
.
├── internal/
│   ├── config/       # Konfigurasi environment (.env)
│   ├── database/     # Inisialisasi koneksi database
│   ├── models/       # GORM Models & Data Access Layer
│   ├── routes/       # Handlers & Router orchestration
│   │   ├── common/   # Middleware & Response templates (Global)
│   │   └── users/    # Modul khusus User
│   └── services/     # Shared services (Logger, dll)
├── main.go           # Entry point aplikasi
└── .env              # Konfigurasi rahasia (API Key, DB, dll)
```

## ⚙️ Persiapan & Instalasi

1. **Clone Repositori**
   ```bash
   git clone https://github.com/misbakhul29/learn-golang-api.git
   cd belajar
   ```

2. **Konfigurasi Environment**
   Salin file `example.env` menjadi `.env` dan sesuaikan nilainya:
   ```bash
   cp example.env .env
   ```

3. **Instal Dependensi**
   ```bash
   go mod tidy
   ```

4. **Jalankan Aplikasi**
   ```bash
   go run main.go
   ```
   Aplikasi akan berjalan di `http://localhost:8080`.

## 📖 Dokumentasi API

Setelah aplikasi berjalan, Anda dapat mengakses dokumentasi interaktif di:
- **Swagger UI**: `http://localhost:8080/docs`
- **OpenAPI Spec**: `http://localhost:8080/openapi.json`

## 🛡️ Contoh Penggunaan (cURL)

**Mendapatkan Semua User:**
```bash
curl --request GET \
  --url http://localhost:8080/api/users \
  --header 'X-API-Key: mysecretkey'
```

**Update User (Partial):**
```bash
curl --request PUT \
  --url http://localhost:8080/api/users/{id} \
  --header 'Content-Type: application/json' \
  --header 'X-API-Key: mysecretkey' \
  --data '{
    "email": "new_email@example.com"
  }'
```

---
Dibuat dengan ❤️ oleh [Misbakhul Munir](https://misbakhul.my.id)
