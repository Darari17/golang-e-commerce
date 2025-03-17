# Golang E-Commerce

🚀 **Golang E-Commerce** adalah proyek backend untuk sistem e-commerce yang dibangun menggunakan Golang. Proyek ini dibuat sebagai bagian dari challenge dan mencakup fitur-fitur utama seperti autentikasi pengguna, manajemen produk, dan transaksi.

## 📌 Fitur

- ✅ Autentikasi menggunakan JWT
- ✅ Manajemen produk (CRUD)
- ✅ Manajemen kategori produk
- ✅ Manajemen pengguna (registrasi, login, update profil)
- ✅ Manajemen pesanan (checkout, pembayaran, status pesanan)
- ✅ Middleware otorisasi

## 🏗️ Teknologi yang Digunakan

- **Golang** sebagai backend utama
- **Gin** sebagai web framework
- **GORM** sebagai ORM untuk database
- **PostgreSQL/MySQL** sebagai database
- **JWT** untuk autentikasi

## 🚀 Instalasi dan Menjalankan Proyek

### 1️⃣ Clone Repository

```sh
git clone https://github.com/Darari17/golang-e-commerce.git
cd golang-e-commerce
```

### 2️⃣ Konfigurasi Database

Buat database baru di PostgreSQL atau MySQL, lalu sesuaikan file `.env`:

```env
DB_HOST=localhost
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_db
DB_PORT=your_port
DB_SSLMODE=disable
TIMEZONE=UTC
```

### 3️⃣ Install Dependencies

```sh
go mod tidy
```

### 4️⃣ Jalankan Migrasi Database

```sh
go run main.go migrate
```

### 5️⃣ Jalankan Server

```sh
go run main.go
```

Server akan berjalan di `http://localhost:8080`.

## 🔥 API Documentation

Untuk dokumentasi API, gunakan Postman atau Swagger:

- [Postman Collection](#) (Tambahkan link jika tersedia)
- [Swagger Docs](#) (Tambahkan link jika tersedia)

## 📌 Struktur Proyek

```bash
📂 golang-e-commerce
├── 📂 config       # Konfigurasi database dan environment
├── 📂 controllers  # Handler untuk HTTP request
├── 📂 middlewares  # Middleware untuk autentikasi dan otorisasi
├── 📂 models       # Struktur data dan ORM model
├── 📂 repositories # Layer akses database
├── 📂 routes       # Routing untuk setiap fitur
├── 📂 services     # Logika bisnis utama
├── 📂 utils        # Helper functions
└── main.go        # Entry point aplikasi
```

## 🔗 Roadmap & Pengembangan Selanjutnya

- [ ] Implementasi pembayaran
- [ ] Pencarian produk dengan filter
- [ ] Sistem review dan rating
- [ ] Integrasi dengan frontend

## 🤝 Kontribusi

Jika ingin berkontribusi, silakan fork repository ini dan ajukan pull request.

## 📝 Lisensi

Proyek ini dirilis di bawah lisensi **MIT**.

---

💡 _Jika ada pertanyaan atau saran, jangan ragu untuk menghubungi saya melalui [GitHub Issues](https://github.com/Darari17/golang-e-commerce/issues)._
